const app = {
	user: null,
	setUser: (user) => localStorage.setItem('user', JSON.stringify(user))
}

let user = localStorage.getItem("user")
if (user !== null) {
	app.user = JSON.parse(user)
}

let id = localStorage.getItem("id")
