package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	id   int
	user User
	date time.Time
}

func requestSayAllMessage(w http.ResponseWriter, r *http.Request) {
	for message := range messages {
		e, err := json.MarshalIndent(messages[message].id, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, string(e))
	}
}

func requestSayMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if (indice - 1) < len(messages) {
		e, err := json.MarshalIndent(messages[indice-1].user.Mail, "", "  ")
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
	w.Write([]byte("Ce message n'existe pas"))
}

/*func createMessage(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	messageCount++
	appendMessage(buf.String())
	w.WriteHeader(http.StatusCreated)
}*/

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	messages = append(messages[:(indice-1)], messages[indice:]...)
	messageCount--
	w.WriteHeader(http.StatusNoContent)
}
