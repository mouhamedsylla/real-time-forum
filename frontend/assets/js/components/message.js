import MessageAPI from "../api/messages.js"

export default class Message {
    constructor() {
        this.targetElement = null
    }

    setTargetElement(element) {
        this.targetElement = element
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
                <i class="far fa-laugh-beam"></i><input placeholder="Type your message here!" type="text" />
            </div>
        `
        return elem
    }

    render(idContact, apiMessage) {
        const contact = apiMessage.getUserById(idContact)
        for (let i = 0; i < this.targetElement.children.length; i++) {
            if (this.targetElement.children[i].className.includes("chat")) {
                this.targetElement.children[i].remove()
            }
        }
        this.targetElement.appendChild(this.createMessageHTML(contact))
    }
}