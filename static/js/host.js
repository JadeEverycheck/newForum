
let host = localStorage.getItem("host")

if (host == null) {
	localStorage.setItem('host', "https://boiling-waters-46327.herokuapp.com")
	// localStorage.setItem('host', "http://localhost:8080")
	host = localStorage.getItem("host")
}

