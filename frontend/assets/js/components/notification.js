import NotificationAPI from "../api/notifications.js"
import { session_expired, alert_token_expire } from "../utils/utils.js"

export default class Notification {
    constructor() {
        this.apiNotification = new NotificationAPI()
        this.contactStatus = null
        this.targetElement = null
    }

    setStatusContact(idContact, bool) {
        this.contactStatus.forEach(contact => {
            const id = contact.getAttribute("id")
            if ((parseInt(id) === idContact)) {
                bool ? contact.classList.add("active") : contact.classList.remove("active")
            }
        });
    }

    notificationHandler(event) {
        const notification = JSON.parse(event.data)
        if (notification.type === "user_status") {
            notification.status === "online" ? this.setStatusContact(notification.id, true) : this.setStatusContact(notification.id, false)
        }
    }

    async initConnectedUser() {
        this.contactStatus = document.querySelectorAll(".status")
        try {
            session_expired() ? alert_token_expire() :
            await this.apiNotification.getConnectedUsers().then(() => {
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