package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	//"reflect"
	"os"
	"strconv"
	"time"
)

type User struct {
	id   int
	mail string
}

type Message struct {
	id   int
	user User
	date time.Time
}

type Discussion struct {
	id    int
	sujet string
	mess  []Message
}

var users = make([]User, 0, 20)
var userCount = 0

func appendUser(email string) {
	userCount++
	users = append(users, User{userCount, email})
}

func initData() {
	appendUser("test1@example.com")
	appendUser("test2@example.com")
	appendUser("test3@example.com")
}

func main() {
	initData()
	fmt.Println(users[0])
	fmt.Println("Début du forum")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", requestSay)
		r.Get("/", requestSayAll)
		r.Post("/", createUser)
		r.Put("/{id}", updateData)
		r.Delete("/{id}", deleteUser)
	})
	http.ListenAndServe(":8080", r)
}
func requestSayAll(w http.ResponseWriter, r *http.Request) {
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

func requestSay(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	if (indice - 1) < len(users) {
		//if users[indice-1].mail != "" {
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

func createUser(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	userCount++
	appendUser(buf.String())
	w.WriteHeader(http.StatusCreated)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	users = append(users[:(indice-1)], users[indice:]...)
	userCount--
	w.WriteHeader(http.StatusNoContent)
}
