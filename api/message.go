package api

import (
	//"encoding/json"
	"forum/response"
	"github.com/go-chi/chi"
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

// 	// id := chi.URLParam(r, "id")
// 	// indice, err := strconv.Atoi(id)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// if (indice - 1) < len(messages) {
// 	// 	e, err := json.MarshalIndent(messages[indice-1].user.Mail, "", "  ")
// 	// 	if err != nil {
// 	// 		w.WriteHeader(http.StatusInternalServerError)
// 	// 		fmt.Println(err)
// 	// 		return
// 	// 	}
// 	// 	w.WriteHeader(http.StatusOK)
// 	// 	fmt.Fprintln(w, string(e))
// 	// 	return
// 	// }
// 	// w.WriteHeader(http.StatusNotFound)
// 	// w.Write([]byte("Ce message n'existe pas"))
// }

// /*func createMessage(w http.ResponseWriter, r *http.Request) {
// 	buf := new(bytes.Buffer)
// 	buf.ReadFrom(r.Body)
// 	messageCount++
// 	appendMessage(buf.String())
// 	w.WriteHeader(http.StatusCreated)
// }*/

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
