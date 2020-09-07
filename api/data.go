package api

var Passwords = map[string]string{
	"test1@example.com": "password1",
	"test2@example.com": "password2",
	"test3@example.com": "password3",
}

func InitData() {
	for login, password := range Passwords {
		appendUser(login, password)
	}
	appendDiscussion("Present")
	appendDiscussion("Futur")
	appendMessage(users[0].Id, "jeSuisUnContenuDeMessage", &discussions[0])
	appendMessage(users[0].Id, "jeSuisUnSecondMessage", &discussions[0])
	appendMessage(users[1].Id, "jeSuisUnMessageAutrePart", &discussions[1])
}
