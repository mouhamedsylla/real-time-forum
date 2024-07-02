import api from "../../index.js"
import { session_expired, alert_token_expire, throttle, alert_loading, prependChild} from "../utils/utils.js"
export default class Message {
    constructor() {
        this.targetElement = null
        this.doLoad = null
        this.messages = []
        this.pointer = 1
        this.amount = 10
    }

    setTargetElement(element) {
        this.targetElement = element
    }

    newMessage(content, type) {
        const message = document.createElement("div")
        message.classList.add("message", type)
        message.textContent = content
        return message
    }

    addMessagesInTop() {
        this.messages.forEach(message => {
            const target = document.querySelector(".chat-messages")
            prependChild(target, this.newMessage(message.Content, message.SenderId === api.client.Id ? "user" : "other"))
        })
    }

    addMessages() {
        this.messages.forEach(message => {
            const target = document.querySelector(".chat-messages")
            target.appendChild(this.newMessage(message.Content, message.SenderId === api.client.Id ? "user" : "other"))
        })
    }

    handleScroll(contactId, apiMessage) {
        const target = document.querySelector(".chat-messages")        
        target.addEventListener('scroll', throttle(async function(e) {    
            if ( target.scrollTop === 0 ) {
                const container = document.querySelector(".chat-messages")
                const firstchild = container.firstChild
                alert_loading(container, this.doLoad)
                await new Promise(resolve => setTimeout(resolve, 2000))
                this.pointer++
                if (!firstchild.className.includes("loading")) {
                    this.messages = await apiMessage.getMessagesPage(contactId, this.pointer, this.amount)
                }
                if (this.messages === null) {
                    this.doLoad = false
                }
                
                await this.addMessagesInTop()
            } 
        }.bind(this), 1000, { leading: true, trailing: true }))
    }

    async onloadDiscussion(contactId, apiMessage) {
        const target = document.querySelector(".chat-messages")        
        try {
            session_expired() ? alert_token_expire() :
            this.pointer = 1
            this.messages = await apiMessage.getMessagesPage(contactId, this.pointer, this.amount)
            this.addMessagesInTop()
            target.scrollTop = target.scrollHeight
            this.doLoad = true
            this.handleScroll(contactId, apiMessage)
        } catch (error) {
            console.error("Error while loading messages: ", error)
        }
    }

    sendMessage(socket) {
        const input = document.querySelector("#message-text")
        const target = document.querySelector(".chat-messages")
        input.addEventListener("keypress", (event) => {
            if (event.key === "Enter") {
                if (session_expired()) { 
                    alert_token_expire()
                } else {
                    target.appendChild(this.newMessage(input.value, "user"))
                    socket.send(input.value)
                    input.value = ""
                }
            }
        })
    }

    receiveMessage(event) {
        const target = document.querySelector(".chat-messages")
        session_expired() ? alert_token_expire() :
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
        const contact = api.getUserById(contactId)
        for (let i = 0; i < this.targetElement.children.length; i++) {
            if (this.targetElement.children[i].className.includes("chat")) {
                this.targetElement.children[i].remove()
            }
        }

        this.targetElement.appendChild(this.createMessageHTML(contact))
    }
}