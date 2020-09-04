package main

import (
	//"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	//"reflect"
	"time"
)

type User struct {
	mail string
}

type Message struct {
	user User
	date time.Time
}

type Discussion struct {
	sujet string
	mess  []Message
}

var users = make([]User, 100)

func main() {
	fmt.Println("Début du forum")
	users[0] = User{"jeSuisUnUser@gmail.com"}
	users[1] = User{"jeSuisUnAutreUser@gmail.com"}
	users[2] = User{"jeSuisLeDernierUser@gmail.com"}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/users", func(r chi.Router) {
		//r.Get("/{mail}", requestSay)
		r.Get("/", requestSayAll)
	})
	http.ListenAndServe(":8080", r)
}
func requestSayAll(w http.ResponseWriter, r *http.Request) {
	/*e, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(e))*/
	w.Write([]byte("hello"))
}

/*func requestSay(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "mail")
	value, ok :=
	if ok {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(value))
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Cette data n'existe pas"))
} */
