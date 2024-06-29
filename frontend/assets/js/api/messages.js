import api from "../../index.js";
export default class MessageAPI {
    constructor() {
        this.messages = []
        this.otherUser = {}
        this.socket = null
    }

    initDiscussion(idContact, callbackSend, callbackReceive) {
        console.log(`chat initialized with ${idContact}`)
        this.socket = new WebSocket(`ws://localhost:9090/chat/message/private/send/${api.client.Id}?user_id=${idContact}`)
        
        this.socket.onopen = () => { console.log("Socket open") }
        this.socket.onclose = () => { console.log("Socket closed") }
        
        callbackSend(this.socket)
        this.socket.onmessage = (event) => {
            callbackReceive(event)
        }
        
    }

    async getMessages(contactId) {
        this.messages = await api.get(`/chat/message/private/${api.client.Id}/${contactId}`)
    }

    async getOtherUser() {
        await api.get("/auth/getUsers").then(data => {
            this.otherUser = data.filter(user => user.Id != api.client.Id)
        })
    }

    getUserById(id) {
        return this.otherUser.find(user => user.Id === parseInt(id))
    }


}