import api from "../../index.js";
export default class PostAPI {
    constructor() {
        this.posts = []
    }

    async getPosts() {
        this.posts = await api.get("/posts/getAllPost")
    }

    async createPost(payload) {
        const formData = new FormData()
        formData.append("image", payload.Image)
        formData.append("title", payload.Title)
        formData.append("content", payload.Content)
        return await api.post(`/posts/createdpost/${api.client.Id}`, formData)
    }
}