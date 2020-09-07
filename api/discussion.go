package api

import (
	"encoding/json"
	"forum/response"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Discussion struct {
	Id      int       `json:"id"`
	Subject string    `json:"subject"`
	Mess    []Message `json:"message,omitempty"`
}

var discussions = make([]Discussion, 0, 20)
var discussionCount = 0

func appendDiscussion(sujet string) Discussion {
	discussionCount++
	discussion := Discussion{
		Id:      discussionCount,
		Subject: sujet,
		Mess:    []Message{},
	}
	discussions = append(discussions, discussion)
	return discussion
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
}

func GetAllDiscussion(w http.ResponseWriter, r *http.Request) {
	response.Ok(w, discussions)
}

func GetDiscussion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	for _, disc := range discussions {
		if disc.Id == indice {
			response.Ok(w, disc)
			return
		}
	}
	response.NotFound(w)
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
	d = appendDiscussion(d.Subject)
	response.Created(w, d)
}

func DeleteDiscussion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	indice, err := strconv.Atoi(id)
	if err != nil {
		response.Deleted(w)
		return
	}
	removeDiscussion(Discussion{Id: indice})
	response.Deleted(w)
}
