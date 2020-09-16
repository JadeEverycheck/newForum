package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"new-forum/apiForum/response"
	"strconv"
)

type Discussion struct {
	Id      int       `json: "id"`
	Subject string    `json: "subject"`
	Mess    []Message `json: "messages, omitempty"`
}

var discussions = make([]Discussion, 0, 20)
var discussionCount = 0

func appendDiscussion(subject string) Discussion {
	discussionCount++
	var d = Discussion{
		Id:      discussionCount,
		Subject: subject,
		Mess:    []Message{},
	}
	discussions = append(discussions, d)
	return d
}

func removeDiscussion(d Discussion) {
	index := -1
	for i, disc := range discussions {
		if disc.Id == d.Id {
			index = i
			break
		}
	}
	if index == -1 {
		return
	}
	copy(discussions[index:], discussions[index+1:])
	discussions = discussions[:len(discussions)-1]
	return
}

func GetAllDiscussions(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, discussions)
	return
}

func GetDiscussion(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	index, err := strconv.Atoi(data)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	for _, disc := range discussions {
		if disc.Id == index {
			response.Ok(w, disc)
			return
		}
	}
	response.NotFound(w)
}

func DeleteDiscussion(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	index, err := strconv.Atoi(data)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	removeDiscussion(Discussion{Id: index})
	response.Deleted(w)
}

func CreateDiscussion(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ServerError(w, err.Error())
		return
	}
	r.Body.Close()
	var d Discussion
	err = json.Unmarshal(body, &d)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, appendDiscussion(d.Subject))
}
