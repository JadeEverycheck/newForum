const registerForm = document.getElementById('register-form')

registerForm.onsubmit = function(e) {
	let request = new XMLHttpRequest()
	request.open("POST", host + "/users/", true)
	request.send(JSON.stringify({ mail: e.target.elements['email'].value }))
	request.onload = () => {
		if (request.status == 201) {
			window.location.replace("../discussion/list.html")
			app.setUser(request.response + JSON.stringify({ password: e.target.elements['password'].value }))
		} else {
			alert("Registration has failed")
		}
	}

	return false
}