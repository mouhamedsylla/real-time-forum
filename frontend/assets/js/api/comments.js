import api from "../../index.js";
export default class CommentAPI {
    constructor() {
        this.comments = []
    }

    async getComments(idPost) {
        try {
            const response = await api.get(`/posts/getAllcomment`);
            this.comments = await response || [];
        } catch (error) {
            console.error("Error getting comments:", error);
        }
    }
    

    async postComment(comment, idPost) {
       try {
            return await api.post(`posts/${idPost}/comment`, comment)
       } catch (error) {
            console.error("Error posting comment: ", error)
       }
    }

    async getUserByCommentId(id) {
        try {
            return await api.get(`/auth/getUsers?userId=${id}`)
        } catch (error) {
            console.error("Error getting user by post id:", error)
        }
    }
}