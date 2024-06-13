export default {
    async loadPosts(context) {
        try {
            const response = await fetch("http://localhost:3000/posts/getAllPost", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            const posts = await response.json();
            context.commit("loadPosts", posts);

        } catch (error) {
            console.error('Failed to fetch posts:', error);
        }     
    },

    async createPost(context, payload) {
        try {
            const idUser = context.state.user.id
            const formData = new FormData()
            formData.append("image", payload.Image)
            formData.append("title", payload.Title)
            formData.append("content", payload.Content)
            const response = await fetch(`http://localhost:3000/posts/createdpost/${idUser}`, {
                method: "POST",
                body: formData
            })
            const result = await response.json()
            context.commit("createPost", payload)
        } catch (error) {
            console.error('Failed to create post:', error);   
        }
    },

    addPost(context, payload) {
        context.commit("addPost", payload)
    }
}