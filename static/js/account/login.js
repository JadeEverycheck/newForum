const loginForm = document.getElementById('login-form')

loginForm.onsubmit = function(e) {
	localStorage.setItem('user', JSON.stringify({
		email: e.target.elements['email'].value,
		password: e.target.elements['password'].value
	}))
	localStorage.setItem('email', e.target.elements['email'].value)
	localStorage.setItem('password', e.target.elements['password'].value)
	window.location.replace("../discussion/list.html")

	return false
}