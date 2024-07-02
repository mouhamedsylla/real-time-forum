import MessageAPI from "../api/messages.js";
import Message from "./message.js";
import { session_expired, alert_token_expire } from "../utils/utils.js";
import api from "../../index.js";

export default class Discussion {
    constructor() {
        this.apiMessage = new MessageAPI()
        this.message = new Message()
        this.targetElement = null
    }

    setTargetElement(element) {
        this.targetElement = element
    }

    createDiscussionHTML(discussion) {
        const elem = document.createElement("div")
        elem.classList.add("message")
        elem.setAttribute("id", discussion.Id)
        const content = `
                <div class="profile-photo">
                    <img src="./frontend/assets/images/profile-${discussion.Id}.jpg" alt="">
                    <div class="status" id="${discussion.Id}"></div>
                </div>
                <div class="message-body">
                    <h5>${discussion.firstName} ${discussion.lastName}</h5>
                </div>`
        elem.innerHTML = content
        return elem
    }

    render() {
        if (session_expired()) {
            alert_token_expire()
        } else {
            api.sortUsers()
            api.otherUser.forEach(user => {
                const currentElement = this.createDiscussionHTML(user)
                this.targetElement.appendChild(currentElement)
                currentElement.addEventListener("click", async () => {
    
                    this.message.setTargetElement(document.querySelector(".right"))
                    this.message.render(user.Id, this.apiMessage)
    
                    const sendMessageCallback = (socket) => this.message.sendMessage(socket)
                    const receiveMessageCallback = (event) => this.message.receiveMessage(event)
                    this.apiMessage.initDiscussion(user.Id, sendMessageCallback, receiveMessageCallback)
    
                    session_expired() ? alert_token_expire() : await this.message.onloadDiscussion(user.Id, this.apiMessage)
                })
            })
        }  
    }
}