package api

var Passwords = map[string]string{
	"test1@example.com": "password1",
	"test4@example.com": "password4",
}

func InitData() {
	for key, value := range Passwords {
		appendUser(key, value)
	}
	appendDiscussion("Animals")
	appendDiscussion("Weather")
	appendDiscussion("France")
	appendMessage("Chien", 1, &discussions[0])
	appendMessage("Chat", 1, &discussions[0])
	return
}
