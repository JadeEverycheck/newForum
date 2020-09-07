package api

import (
	"encoding/json"
	"forum/response"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Id      int       `json:"id"`
	UserId  int       `json:"user_id"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
}

var messageCount = 0

func appendMessage(uId int, content string, disc *Discussion) Message {
	messageCount++
	message := Message{
		Id:      messageCount,
		Date:    time.Now(),
		UserId:  uId,
		Content: content,
	}
	disc.Mess = append(disc.Mess, message)
	return message
}

func removeMessage(m Message) {
	index := -1
	discIndex := -1
	for i, disc := range discussions {
		for j, mess := range disc.Mess {
			if mess.Id == m.Id {
				index = j
				discIndex = i
				break
			}
		}
	}
	if index == -1 || discIndex == -1 {
		return
	}
	copy(discussions[discIndex].Mess[index:], discussions[discIndex].Mess[index+1:])
	discussions[discIndex].Mess = discussions[discIndex].Mess[:len(discussions[discIndex].Mess)-1]
}

func GetAllMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	for _, disc := range discussions {
		if disc.Id == indice {
			response.Ok(w, disc.Mess)
			return
		}
	}
	response.NotFound(w)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	for _, disc := range discussions {
		for _, mess := range disc.Mess {
			if mess.Id == indice {
				response.Ok(w, mess)
				return
			}
		}
	}
	response.NotFound(w)
}

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	d := &Discussion{}
	for i, disc := range discussions {
		if disc.Id == indice {
			d = &discussions[i]
			break
		}
	}
	if d.Id == 0 {
		response.NotFound(w)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}
	r.Body.Close()
	var m Message
	err = json.Unmarshal(body, &m)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	username, _, _ := r.BasicAuth()
	user := User{}
	for _, u := range users {
		if u.Mail == username {
			user = u
		}
	}

	m = appendMessage(user.Id, m.Content, d)
	response.Created(w, m)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.Deleted(w)
		return
	}
	removeMessage(Message{Id: indice})
	response.Deleted(w)
}
