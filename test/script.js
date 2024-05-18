const inputText1 = document.getElementById("mess-1")
const inputText2 = document.getElementById("mess-2")
const ul1 = document.querySelector("send-1")
const ul2 = document.querySelector("send-2")
const socket1 = new WebSocket("ws://localhost:9090/api/message/private/send/user-2?user_id=user-1")
const socket2 = new WebSocket("ws://localhost:9090/api/message/private/send/user-1?user_id=user-2")

socket1.onmessage = function(event) {
    const li = document.createElement("li")
    li.textContent = `sended: ${event.data}`
    ul1.appendChild(li)
}


socket2.onmessage = function(event) {
    const li = document.createElement("li")
    li.textContent = `sended: ${event.data}`
    ul2.appendChild(li)
}

function sendMessage1() {
    const message = inputText1.value
    socket1.send(JSON.stringify({
        content : message
    }))
}

function sendMessage2() {
    const message = inputText2.value
    socket2.send(JSON.stringify({
        content : message
    }))
}

// const b1 = document.querySelector("btn-1")
// b1.addEventListener("click", sendMessage(inputText1, socket1))
// const b2 = document.querySelector("btn-2")
// b2.addEventListener("click", sendMessage(inputText2, socket2))

