import api from "../../index.js";
export default class MessageAPI {
    constructor() {
        this.messages = []
        this.socket = null
    }

    initDiscussion(idContact, callbackSend, callbackReceive) {
        this.socket = new WebSocket(`ws://localhost:9090/chat/message/private/send/${idContact}?user_id=${api.client.Id}`)
        
        callbackSend(this.socket)
        callbackReceive(this.socket)
        // this.socket.onmessage = (event) => {
        //     callbackReceive(event)
        // }
        
    }

    async getMessages(contactId) {
        this.messages = await api.get(`/chat/message/private/${api.client.Id}/${contactId}`)
    }

    async getMessagesPage(contactId, page, limit = 10) {
       try {
            const response = await api.get(`/chat/message/private/${api.client.Id}/${contactId}?page=${page}&limit=${limit}`)
            return response
       } catch (error) {
            console.error("Error while fetching messages: ", error)
       }
        
    }
}