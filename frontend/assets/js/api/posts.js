import api from "./main.js";
export default class PostAPI {
    constructor(targetElement, api) {
        this.api = api
        this.posts = []
        this.element = targetElement
        this.currentPost = {}
    }

    async getPosts() {
        return await this.api.get("posts/getAllPost")
    }

    async createPost(payload) {
        const formData = new FormData()
        formData.append("image", payload.Image)
        formData.append("title", payload.Title)
        formData.append("content", payload.Content)
        return await this.api.post(`posts/createdpost/${this.api.client.id}}`, formData)
    }

    addPost() {
        
    }
}