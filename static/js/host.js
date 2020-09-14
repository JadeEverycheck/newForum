
let host = localStorage.getItem("host")
	console.log(host)
if (host == null) {
	localStorage.setItem('host', "https://protected-brushlands-86376.herokuapp.com")
	// localStorage.setItem('host', "http://localhost:8080")
	host = localStorage.getItem("host")
	console.log(host)
}

