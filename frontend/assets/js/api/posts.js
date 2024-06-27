export default class Posts {
    constructor(targetElement) {
        this.posts = []
        this.element = targetElement
        this.currentPost = {}
    }

    async getPosts() {
        try {
            const response = await fetch("http://localhost:3000/posts/getAllPost", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            this.posts = await response.json();
        } catch (error) {
            console.error('Failed to fetch posts:', error);
        }
    }

    async createPost(payload) {
        try {
            const formData = new FormData()
            formData.append("image", payload.Image)
            formData.append("title", payload.Title)
            formData.append("content", payload.Content)
            const response = await fetch(`http://localhost:3000/posts/createdpost/${context.state.user.id}`, {
                method: "POST",
                body: formData
            })
            const result = await response.json()
        } catch (error) {
            console.error('Failed to create post:', error);   
        }
    }

    addPost() {
        
    }
}