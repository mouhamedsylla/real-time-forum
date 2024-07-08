import api from "../../index.js";
export default class PostAPI {
    constructor() {
        this.posts = []
        this.lastPostId = null
    }

    async getPosts() {
        try {
            this.posts = await api.get("/posts/getAllPost")
            return this.posts
        } catch (error) {
            console.error("Error getting posts:", error)
        }
    }

    async getReactionPost() {
        try {
            console.log(api.client.Id)
            return await api.get(`/posts/GetUserPostReactions/${api.client.Id}`)
        } catch (error) {
            console.error("Error getting reaction posts:", error)
        }
    }

    async createPost(payload) {
        const formData = new FormData()
        formData.append("image", payload.Image)
        formData.append("title", payload.Title)
        formData.append("content", payload.Content)
        formData.append("categories", payload.Categories)
        this.lastPostId = await api.post(`posts/createdpost/${api.client.Id}`, formData)
    }

    async getUserByPostId(id) {
        try {
            return await api.get(`/auth/getUsers?userId=${id}`)
        } catch (error) {
            console.error("Error getting user by post id:", error)
        }
    }

    async reactionPost(idPost, reaction) {
        try {
            return await api.post(`posts/ReactionPosts/${api.client.Id}/${idPost}`, reaction)
        } catch (error) {
            console.error("Error reacting to post: ", error)
        }
    }
}