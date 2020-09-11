const newDiscForm = document.getElementById('newDisc-form')
let email = localStorage.getItem("email")
let password = localStorage.getItem("password")
const mail = document.getElementById('user')
mail.appendChild(document.createTextNode(email))

function signOut() {
    localStorage.clear()
    window.location.replace("../../index.html")
}

newDiscForm.onsubmit = function(e) {
    let request = new XMLHttpRequest()
    request.open("POST", host + "/discussions/", true)
    request.setRequestHeader('Authorization', 'Basic '+btoa(email+":"+password))
    request.send(JSON.stringify({ subject: e.target.elements['subject'].value }))
    request.onload = () => {
        console.log(request)
        if (request.status == 201) {
            window.location.replace("../discussion/list.html")
        } else {
            alert("Creation of a new discussion has failed")
        }
    }

    return false
}
