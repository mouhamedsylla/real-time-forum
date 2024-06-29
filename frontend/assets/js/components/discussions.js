import MessageAPI from "../api/messages.js";
import Message from "./message.js";

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
                    <div class="active"></div>
                </div>
                <div class="message-body">
                    <h5>${discussion.firstName} ${discussion.lastName}</h5>
                </div>`
        elem.innerHTML = content
        return elem
    }

    async render() {
        await this.apiMessage.getOtherUser()
        .then(() => {
            this.apiMessage.otherUser.forEach(user => {
                const currentElement = this.createDiscussionHTML(user)
                this.targetElement.appendChild(currentElement)
                currentElement.addEventListener("click", async () => {

                    this.message.setTargetElement(document.querySelector(".right"))
                    this.message.render(user.Id, this.apiMessage)

                    const sendMessageCallback = (socket) => this.message.sendMessage(socket)
                    const receiveMessageCallback = (event) => this.message.receiveMessage(event)
                    this.apiMessage.initDiscussion(user.Id, sendMessageCallback, receiveMessageCallback)
                    
                    await this.message.onloadDiscussion(user.Id, this.apiMessage)
                })
            })
        })
    }
}