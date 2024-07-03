import api from "../../index.js";
export default class CommentAPI {
    constructor() {
        this.comments = {}
    }

    async getComments(idPost) {
        try {
            const response = await api.get(`/posts/${idPost}/getcomment`);
            this.comments[idPost] = response;
        } catch (error) {
            console.error("Error getting comments:", error);
        }
    }
    

    async postComment(comment, idPost) {
       return await api.post(`/posts/${idPost}/comment`, comment)
    }
}