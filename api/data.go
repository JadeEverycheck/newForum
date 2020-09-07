package api

import (
	"time"
)

var users = make([]User, 0, 20)
var userCount = 0
var messages = make([]Message, 0, 20)
var messageCount = 0
var messageInit []Message
var discussions = make([]Discussion, 0, 20)
var discussionCount = 0

func appendUser(email string) User {
	userCount++
	user := User{userCount, email}
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

func appendMessage(id int, mail string) {
	messageCount++
	u := User{id, mail}
	messages = append(messages, Message{messageCount, u, time.Now()})
}

func appendDiscussion(sujet string) {
	discussionCount++
	discussions = append(discussions, Discussion{discussionCount, sujet, messageInit})
}

func appendMessToDisc(disc Discussion, mess Message) {
	disc.mess = append(disc.mess, Message{})
}

func InitData() {
	appendUser("test1@example.com")
	appendUser("test2@example.com")
	appendUser("test3@example.com")
	appendMessage(users[0].Id, users[0].Mail)
	appendMessage(users[0].Id, users[0].Mail)
	appendMessage(users[1].Id, users[1].Mail)
	appendDiscussion("Present")
	appendDiscussion("Futur")
	disc1 := Discussion{discussions[0].id, discussions[0].sujet, discussions[0].mess}
	m1 := Message{messages[0].id, messages[0].user, messages[0].date}
	appendMessToDisc(disc1, m1)
}
