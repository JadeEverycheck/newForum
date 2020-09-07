package api

import (
	"encoding/json"
	"forum/response"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Mail string `json:"mail"`
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	for _, user := range users {
		if user.Id == indice {
			response.Ok(w, user)
			return
		}
	}
	response.NotFound(w)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}
	r.Body.Close()
	var u User
	err = json.Unmarshal(body, &u)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	u = appendUser(u.Mail)
	response.Created(w, u)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	u := User{}
	for _, user := range users {
		if user.Id == indice {
			u = user
			break
		}
	}

	if u.Id == 0 {
		response.NotFound(w)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}
	r.Body.Close()
	var updated User
	err = json.Unmarshal(body, &u)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	u.Mail = updated.Mail
	response.Ok(w, u)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.Deleted(w)
		return
	}

	for i, user := range users {
		if user.Id == indice {
			users = append(users[:(i-1)], users[i:]...)
			response.Deleted(w)
			return
		}
	}
	response.Deleted(w)
}
