import api from "../../index.js"

export default class Message {
    constructor() {
        this.targetElement = null
    }

    setTargetElement(element) {
        this.targetElement = element
    }

    newMessage(content, type) {
        console.log("New message: ", content)
        const message = document.createElement("div")
        message.classList.add("message", type)
        message.textContent = content
        return message
    }

    async onloadDiscussion(contactId, apiMessage) {
        try {
            await apiMessage.getMessages(contactId)
            apiMessage.messages.forEach(message => {
                const target = document.querySelector(".chat-messages")
                target.appendChild(this.newMessage(message.Content, message.SenderId === api.client.Id ? "user" : "other"))
            })
        } catch (error) {
            console.error("Error while loading messages: ", error)
        }
    }

    sendMessage(socket) {
        const input = document.querySelector("#message-text")
        const target = document.querySelector(".chat-messages")
        input.addEventListener("keypress", (event) => {
            if (event.key === "Enter") {
                target.appendChild(this.newMessage(input.value, "user"))
                socket.send(input.value)
                input.value = ""
            }
        })
    }

    receiveMessage(event) {
        const target = document.querySelector(".chat-messages")
        target.appendChild(this.newMessage(event.data, "other"))
    }

    createMessageHTML(contact) {
        const elem = document.createElement("div")
        elem.classList.add("chat")
        elem.innerHTML = `
            <div class="contact bar">
                <div class="profile-photo pic">
                    <img src="./frontend/assets/images/profile-${contact.Id}.jpg" alt="">
                    <div class="active"></div>
                </div>
                <div class="name">
                    ${contact.firstName} ${contact.lastName}
                </div>
                <div class="seen">
                    Today at 12:56
                </div>
            </div>
            <div class="chat-messages">
                <div class="time">
                    Today at 11:41
                </div>
            </div>
            <div class="input">
                <i class="far fa-laugh-beam"></i><input id="message-text" placeholder="Type your message here!" type="text" />
            </div>
        `
        return elem
    }

    render(contactId, apiMessage) {
        const contact = apiMessage.getUserById(contactId)
        for (let i = 0; i < this.targetElement.children.length; i++) {
            if (this.targetElement.children[i].className.includes("chat")) {
                this.targetElement.children[i].remove()
            }
        }
        this.targetElement.appendChild(this.createMessageHTML(contact))
    }
}