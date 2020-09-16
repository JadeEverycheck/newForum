package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"new-forum/apiForum/api"
	"os"
	"path/filepath"
	"strings"
)

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	fmt.Println("Début du forum")
	api.InitData()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "static"))
	FileServer(r, "/static", filesDir)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", api.GetAllUsers)
		r.Get("/{id}", api.GetUser)
		r.Delete("/{id}", api.DeleteUser)
		r.Post("/", api.CreateUser)
		r.Put("/{id}", api.UpdateUser)
	})
	r.Route("/discussions", func(r chi.Router) {
		r.Use(middleware.BasicAuth("real", api.Passwords))

		r.Get("/", api.GetAllDiscussions)
		r.Get("/{id}", api.GetDiscussion)
		r.Delete("/{id}", api.DeleteDiscussion)
		r.Post("/", api.CreateDiscussion)
		r.Get("/{id}/messages", api.GetAllMessages)
		r.Get("/messages/{id}", api.GetMessage)
		r.Post("/{id}/messages", api.CreateMessage)
		r.Delete("/messages/{id}", api.DeleteMessage)
	})
	http.ListenAndServe(":"+port, r)

}
