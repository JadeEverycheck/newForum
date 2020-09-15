let login = document.getElementById('register-form')


login.onsubmit = function(e) {
	localStorage.setItem('mail', e.target.elements['email'].value)
	localStorage.setItem('password', e.target.elements['password'].value)
	return
}