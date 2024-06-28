import api from "../../index.js";
export default class CommentAPI {
    constructor() {
        this.comments = {}
    }

    async getComments(idPost) {
        this.comments[idPost] = await api.get(`/posts/${idPost}/getcomment`)
    }
}