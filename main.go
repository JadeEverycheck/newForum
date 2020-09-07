package main

import (
	"fmt"
	"forum/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func main() {
	api.InitData()
	fmt.Println("Début du forum")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", api.GetUser)
		r.Get("/", api.GetAllUser)
		r.Post("/", api.CreateUser)
		r.Put("/{id}", api.UpdateUser)
		r.Delete("/{id}", api.DeleteUser)
	})

	r.Route("/discussions", func(r chi.Router) {
		r.Use(middleware.BasicAuth("real", api.Passwords))
		r.Get("/{id}", api.GetDiscussion)
		r.Get("/", api.GetAllDiscussion)
		r.Post("/", api.CreateDiscussion)
		r.Delete("/{id}", api.DeleteDiscussion)

		r.Get("/messages/{id}", api.GetMessage)
		r.Get("/{id}/messages", api.GetAllMessage)
		r.Post("/{id}/messages", api.CreateMessage)
		r.Delete("/messages/{id}", api.DeleteMessage)
	})

	http.ListenAndServe(":8080", r)
}
