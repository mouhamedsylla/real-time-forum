import api from "../../index.js";
export default class MessageAPI {
    constructor() {
        this.messages = {}
        this.otherUser = {}
    }

    async getMessages(idUser) {
        this.messages[idUser] = await api.get(`/chat/message/private/${api.client.Id}/${idUser}`)
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