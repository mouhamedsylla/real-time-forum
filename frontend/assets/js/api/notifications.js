import api from "../../index.js"; 
export default class NotificationAPI {
    constructor() {
        this.connectUser = []
        this.socket = null
    }

    connectToNotificationService(handleCallback) {
        this.socket = new WebSocket(`ws://localhost:9191/notification/sendNotification/${api.client.Id}`)
        this.socket.onopen = () => { console.log("Socket open") }
        this.socket.onmessage = (event) => {
            handleCallback(event)
        }
    }

    async getConnectedUsers() {
        try {
            this.connectUser = await api.get(`/chat/message/private/getConnectedUser/${api.client.Id}`)
        } catch (error) {
            console.error("Error getting connected users:", error)
        }
    }


}