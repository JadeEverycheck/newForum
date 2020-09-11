if (app.user === null) {
	window.location.replace("./login.html")
}

const deleteButton = document.getElementById('delete')
deleteButton.onclick = () => {
	if (confirm('Are you sure?')) {
		let request = new XMLHttpRequest()
		let id = localStorage.getItem("id")
		request.open("DELETE", host + "/users/" + id, true)
		request.send(null)
		request.onload = () => {
			if (request.status == 204) {
				window.location.replace("./login.html")
			} else {
				alert("Your account has not been deleted")
			}
		}
	}
	return false
}

document.getElementById('login').innerText = app.user.email
document.getElementById('password').innerText = app.user.password