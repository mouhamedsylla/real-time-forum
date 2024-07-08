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
        this.actualContact = {}
        this.typingProgress = {
            id: this.actualContact.id,
            username: "",
            infos: "is typing..."
        }
        this.typingTimeout
        this.typingFlag = false
    }

    newMessage(content, type, username, date) {
        const message = document.createElement("div")
        message.classList.add("message", type)

        const div2 = document.createElement("div")
        div2.classList.add("message-body")
        div2.innerText = content

        const div3 = document.createElement("div")
        div3.classList.add("message-time")
        div3.innerHTML = `
            <span>${username}  ${formatDate(date)}</span>
        `
        message.appendChild(div2)
        message.appendChild(div3)
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
            if (session_expired()) { 
                alert_token_expire()
            } else {
                this.pointer = 1
                this.messages = await this.apiMessage.getMessagesPage(contactId, this.pointer, this.amount)
                this.addMessagesInTop()
                target.scrollTop = target.scrollHeight
                this.doLoad = true
                this.handleScroll(contactId)
            }
            
        } catch (error) {
            console.error("Error while loading messages: ", error)
        }
    }

    receiveMessage(socket) {
        socket.addEventListener("message", (event) => {
            const target = document.querySelector(".chat-messages");
            let data;
    
            try {
                data = JSON.parse(event.data);

                if (!this.typingFlag) {
                    alert_typing();
                    this.typingFlag = true;
                    setTimeout(() => {
                        this.typingFlag = false;
                    }, 2500); // Délai de 2 secondes avant de permettre un autre appel
                }
    
                // Réinitialiser le timeout pour enlever l'animation après 3 secondes d'inactivité
                clearTimeout(this.typingTimeout);
                this.typingTimeout = setTimeout(() => {
                    if (target.querySelector(".typing")) {
                        target.querySelector(".typing").remove();
                    }
                }, 3000);
    
            } catch (error) {
                data = event.data;
                if (session_expired()) { 
                    alert_token_expire() 
                } else {
                    target.appendChild(this.newMessage(data, "other", this.actualContact.username));
                    target.scrollTop = target.scrollHeight
                }
            }
        });
    }

    sendMessage(socket) {
        const input = document.querySelector("#message-text")
        const target = document.querySelector(".chat-messages")
        input.addEventListener("keypress", (event) => {

            if (event.key === "Enter") {
                if (session_expired()) { 
                    alert_token_expire()
                } else {
                    const message = input.value
                    if (message.trim() === "") return
                    target.appendChild(this.newMessage(message, "user", api.client.nickname))
                    target.scrollTop = target.scrollHeight
                    socket.send(message)
                    this.upadateDiscussionProfile(this.actualContact.id)
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
                    <i class="fa-solid fa-xmark" id="close-discussion"></i>
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
            this.actualContact.id = user.Id
            this.actualContact.username = user.nickname

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
        const box_discussion = this.createMessageHTML(contact)
        this.targetElement.appendChild(box_discussion)
        const discussionClose = document.querySelector("#close-discussion")
            discussionClose.addEventListener("click", () => {
                box_discussion.remove()
            })
    }
}