package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"os"
	"strconv"
	"time"
)




var users = make([]User, 0, 20)
var userCount = 0
var messages = make([]Message, 0, 20)
var messageCount = 0
var messageInit []Message
var discussions = make([]Discussion, 0, 20)
var discussionCount = 0

func appendUser(email string) {
	userCount++
	users = append(users, User{userCount, email})
}

func appendMessage(id int, mail string) {
	messageCount++
	u := User{id, mail}
	messages = append(messages, Message{messageCount, u, time.Now()})
}

func appendDiscussion(sujet string) {
	discussionCount++
	discussions = append(discussions, Discussion{discussionCount, sujet, messageInit})
}

func appendMessToDisc(disc Discussion, mess Message) {
	disc.mess = append(disc.mess, Message{)
}

func initData() {
	appendUser("test1@example.com")
	appendUser("test2@example.com")
	appendUser("test3@example.com")
	appendMessage(users[0].id, users[0].mail)
	appendMessage(users[0].id, users[0].mail)
	appendMessage(users[1].id, users[1].mail)
	appendDiscussion("Present")
	appendDiscussion("Futur")
	disc1 := Discussion{discussions[0].id, discussions[0].sujet, discussions[0].mess}
	m1 := Message{messages[0].id, messages[0].user, messages[0].date}
	appendMessToDisc(disc1, m1)
}

func main() {
	initData()
	fmt.Println("Début du forum")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", requestSay)
		r.Get("/", requestSayAll)
		r.Post("/", createUser)
		r.Put("/{id}", updateUser)
		r.Delete("/{id}", deleteUser)
	})
	r.Route("/messages", func(r chi.Router) {
		r.Get("/{id}", requestSayMessage)
		r.Get("/", requestSayAllMessage)
		//r.Post("/", createMessage)
		//r.Put("/{id}", updateMessage)
		r.Delete("/{id}", deleteMessage)

	})
	r.Route("/discussions", func(r chi.Router) {
		r.Get("/{id}", requestSayDiscussion)
		r.Get("/", requestSayAllDiscussion)
		//r.Post("/", createMessage)
		//r.Put("/{id}", updateMessage)
		//r.Delete("/{id}", deleteMessage)

	})
	http.ListenAndServe(":8080", r)
}



