import api from "../../index.js"
import MessageAPI from "../api/messages.js"
import { alert_token_expire, alert_loading, prependChild, alert_typing } from "../utils/alert.js"
import { session_expired, throttle, formatDate } from "../utils/other.js"
export default class Message {
    constructor() {
        this.apiMessage = new MessageAPI()
        this.targetElement = null
        this.doLoad = null
        this.messages = []
        this.pointer = 1
        this.amount = 10
        this.actualContact = null
        this.typingProgress = {
            id: this.actualContact,
            username: "",
            infos: "is typing..."
        }
    }

    newMessage(content, type, username, date) {
        const message = document.createElement("div")
        message.classList.add("message", type)
        message.innerHTML = `
            <div class="message-body">
                ${content}
            </div>
            <div class="message-time">
                <span>${username}  ${formatDate(date)}</span>
            </div>
        `
        return message
    }

    addMessagesInTop() {
        this.messages.forEach(message => {
            const target = document.querySelector(".chat-messages")
            const user = api.discussionsUsers.filter(user => user.Id == message.SenderId)[0]
            const username = user ? user.nickname : api.client.nickname
            prependChild(target, this.newMessage(message.Content, message.SenderId === api.client.Id ? "user" : "other", username, message.CreatedAt))
        })
    }

    addMessages() {
        this.messages.forEach(message => {
            const target = document.querySelector(".chat-messages")
            target.appendChild(this.newMessage(message.Content, message.SenderId === api.client.Id ? "user" : "other", api.client.nickname))
        })
    }

    handleScroll(contactId) {
        const target = document.querySelector(".chat-messages")        
        target.addEventListener('scroll', throttle(async function(e) {    
            if ( target.scrollTop === 0 ) {
                const container = document.querySelector(".chat-messages")
                const firstchild = container.firstChild
                alert_loading(container, this.doLoad)
                await new Promise(resolve => setTimeout(resolve, 2000))
                this.pointer++
                if (!firstchild.className.includes("loading")) {
                    this.messages = await this.apiMessage.getMessagesPage(contactId, this.pointer, this.amount)
                }
                if (this.messages === null) {
                    this.doLoad = false
                }
                
                await this.addMessagesInTop()
            } 
        }.bind(this), 1000, { leading: true, trailing: true }))
    }

    async onloadDiscussion(contactId) {
        const target = document.querySelector(".chat-messages")        
        try {
            session_expired() ? alert_token_expire() :
            this.pointer = 1
            this.messages = await this.apiMessage.getMessagesPage(contactId, this.pointer, this.amount)
            this.addMessagesInTop()
            target.scrollTop = target.scrollHeight
            this.doLoad = true
            this.handleScroll(contactId)
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
                    this.upadateDiscussionProfile(this.actualContact)
                    input.value = ""
                }
            } else {
                socket.send(JSON.stringify(this.typingProgress)) 
            }
        })
    }

    upadateDiscussionProfile(idUser) {
        const discussionProfile = document.getElementById(`discussion-${idUser}`)
        const container = document.querySelector(".discussions")
        container.insertBefore(discussionProfile, container.firstChild)
    }

    receiveMessage(socket) {
        socket.addEventListener("message", (event) => {
            const target = document.querySelector(".chat-messages")
            console.log("est entrain de taper...")
            let data

            try {
                data = JSON.parse(event.data)
                console.log(data)
                alert_typing()
                //setInterval(() => { target.querySelector(".typing").remove() }, 3000)
            } catch (error) {
                data = event.data
                session_expired() ? alert_token_expire() :
                target.appendChild(this.newMessage(data, "other"))
            }
        })
        // const infos = JSON.parse(event.data)
        // const target = document.querySelector(".chat-messages")

        // if (infos && infos.infos === "is typing...") {
        //     console.log("typing...")
        //     alert_typing()
        //     //setTimeout(() => { target.querySelector(".typing").remove() }, 1000)
        //     setInterval(() => { target.querySelector(".typing").remove() }, 3000)
        //     return
        // }

        // session_expired() ? alert_token_expire() :
        // target.appendChild(this.newMessage(event.data, "other"))
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
                <div class="discussion-close">
                    <i class="fa-solid fa-xmark"></i>
                </div>
            </div>
            <div class="chat-messages">

            </div>
            <div class="input">
                <i class="far fa-laugh-beam"></i><input id="message-text" placeholder="Type your message here!" type="text" />
            </div>
        `
        return elem
    }

    handleMessageDiscussion(user, targetElem) {
        targetElem.addEventListener("click", async (e) => {
            this.targetElement = document.querySelector(".right")
            this.render(user.Id)
            this.actualContact = user.Id

            const sendMessageCallback = (socket) => this.sendMessage(socket)
            const receiveMessageCallback = (event) => this.receiveMessage(event)
            this.apiMessage.initDiscussion(user.Id, sendMessageCallback, receiveMessageCallback)
            
            session_expired() ? alert_token_expire() : await this.onloadDiscussion(user.Id)
        })
    }

    render(contactId) {
        const contact = api.getUserById(contactId)
        for (let i = 0; i < this.targetElement.children.length; i++) {
            if (this.targetElement.children[i].className.includes("chat")) {
                this.targetElement.children[i].remove()
            }
        }
        this.targetElement.appendChild(this.createMessageHTML(contact))
    }
}