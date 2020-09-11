
let host = localStorage.getItem("host")
	console.log(host)
if (host == null) {
	localStorage.setItem('host', "https://protected-brushlands-86376.herokuapp.com")
	host = localStorage.getItem("host")
	console.log(host)
}

