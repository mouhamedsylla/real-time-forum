import NotificationAPI from "../api/notifications.js";
import { showToast, session_expired } from "../utils/other.js";
import { alert_token_expire } from "../utils/alert.js";

export default class Notification {
    constructor() {
        this.apiNotification = new NotificationAPI()
        this.contactStatus = null
        this.targetElement = document.querySelector("body")
    }

    createNotifcationHTML(userSended) {
        const div = document.createElement("div")
        div.classList.add("wrapper")
        div.innerHTML = `
            <div id="toast">
            <div class="container-1">
                <i class="fas fa-check-square"></i>
            </div>
            <div class="container-2">
                <p>Notification</p>
                <p><b>${userSended.nickname}</b> send you a new message</p>
            </div>
            <button id="close" onclick="closeToast()">
                &times;
            </button>
        </div>`

        this.targetElement.appendChild(div)
    }

    setStatusContact(idContact, bool) {
        this.contactStatus.forEach(contact => {
            const id = contact.getAttribute("id")
            if ((+id === idContact)) {
                bool ? contact.classList.add("active") : contact.classList.remove("active")
            }
        })
    }


    async notificationHandler(event) {
        const notification = JSON.parse(event.data)
        if (notification.type === "user_status") {
            notification.status === "online" ? this.setStatusContact(notification.id, true) : this.setStatusContact(notification.id, false)
        }

        if (notification.type === "notification") {
            const user = await this.apiNotification.getUserByNotificationId(notification.id)
            this.createNotifcationHTML(user)
            showToast()
            this.upadateDiscussionProfile(notification.id)
        }
    }

    upadateDiscussionProfile(idUser) {
        const discussionProfile = document.getElementById(`discussion-${idUser}`)
        const container = document.querySelector(".discussions")
        container.insertBefore(discussionProfile, container.firstChild)
    }

    async initNotification() {
        this.contactStatus = document.querySelectorAll(".status")
        try {
            session_expired() ? alert_token_expire() :
                await this.apiNotification.getConnectedUsers().then(() => {
                    if (!this.apiNotification.connectUser) return
                    this.apiNotification.connectUser.forEach(user => {
                        this.setStatusContact(user.id, true)
                    })
                })
        } catch (error) {
            console.error("Error getting connected users:", error)
        }

        const handler = (event) => this.notificationHandler(event)
        session_expired() ? alert_token_expire() : this.apiNotification.connectToNotificationService(handler)
    }
}