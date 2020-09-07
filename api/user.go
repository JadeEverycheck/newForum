package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
)

type User struct {
	id   int
	mail string
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	for user := range users {
		e, err := json.MarshalIndent(users[user].mail, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(e))
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if (indice - 1) < len(users) {
		e, err := json.MarshalIndent(users[indice-1].mail, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(e))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Cet utilisateur n'existe pas"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	userCount++
	appendUser(buf.String())
	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if (indice - 1) >= len(users) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Cet utilisateur n'existe pas"))
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	users[indice-1].mail = buf.String()
	e, err := json.MarshalIndent(users[indice-1].mail, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(e))
	return
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	users = append(users[:(indice-1)], users[indice:]...)
	userCount--
	w.WriteHeader(http.StatusNoContent)
}
