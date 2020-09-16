package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"new-forum/apiForum/response"
	"strconv"
)

type User struct {
	Id       int    `json: "id"`
	Mail     string `json: "mail"`
	Password string `json: "password"`
}

var users = make([]User, 0, 20)
var userCount = 0

func appendUser(mail string, password string) User {
	userCount++
	var u = User{
		Id:       userCount,
		Mail:     mail,
		Password: password,
	}
	users = append(users, u)
	return u
}

func removeUser(u User) {
	index := -1
	for i, user := range users {
		if user.Id == u.Id {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	copy(users[index:], users[index+1:])
	users = users[:len(users)-1]
	return
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, users)
	return
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	index, err := strconv.Atoi(data)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	for _, user := range users {
		if user.Id == index {
			response.Ok(w, user)
			return
		}
	}
	response.NotFound(w)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	index, err := strconv.Atoi(data)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	removeUser(User{Id: index})
	response.Deleted(w)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	index, err := strconv.Atoi(data)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	indice := -1
	var updated User
	for i, user := range users {
		if user.Id == index {
			indice = i
			break
		}
	}
	if indice == -1 {
		response.NotFound(w)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}
	r.Body.Close()
	err = json.Unmarshal(body, &updated)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	users[indice].Mail = updated.Mail
	response.Ok(w, users[indice])
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
	response.Created(w, appendUser(u.Mail, u.Password))
}
