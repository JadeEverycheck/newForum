package api

func InitData() {
	appendUser("test1@example.com")
	appendUser("test2@example.com")
	appendUser("test3@example.com")
	appendDiscussion("Present")
	appendDiscussion("Futur")
	appendMessage(users[0].Id, "jeSuisUnContenuDeMessage", &discussions[0])
	appendMessage(users[0].Id, "jeSuisUnSecondMessage", &discussions[0])
	appendMessage(users[1].Id, "jeSuisUnMessageAutrePart", &discussions[1])
}
