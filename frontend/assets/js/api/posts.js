import api from "../../index.js";
export default class PostAPI {
    constructor() {
        this.posts = []
        this.lastPostId = null
    }

    async getPosts() {
        try {
            this.posts = await api.get("/posts/getAllPost")
        } catch (error) {
            console.error("Error getting posts:", error)
        }
    }

    async createPost(payload) {
        const formData = new FormData()
        formData.append("image", payload.Image)
        formData.append("title", payload.Title)
        formData.append("content", payload.Content)
        this.lastPostId = await api.post(`/posts/createdpost/${api.client.Id}`, formData)
    }

    async getUserByPostId(id) {
        try {
            return await api.get(`/auth/getUsers?userId=${id}`)
        } catch (error) {
            console.error("Error getting user by post id:", error)
        }
    }
}