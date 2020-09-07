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
	Id       int    `json:"id"`
	Mail     string `json:"mail"`
	Password string `json:"-"`
}

type CreateUserType struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

var users = make([]User, 0, 20)
var userCount = 0

func appendUser(email string, password string) User {
	userCount++
	user := User{
		Id:       userCount,
		Mail:     email,
		Password: password,
	}
	users = append(users, user)
	return user
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
	var u CreateUserType
	err = json.Unmarshal(body, &u)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, appendUser(u.Mail, u.Password))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	userIndex := -1
	for i, user := range users {
		if user.Id == indice {
			userIndex = i
			break
		}
	}

	if userIndex < 0 {
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
	err = json.Unmarshal(body, &updated)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	users[userIndex].Mail = updated.Mail
	response.Ok(w, users[userIndex])

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.Deleted(w)
		return
	}
	removeUser(User{Id: indice})
	response.Deleted(w)
}
