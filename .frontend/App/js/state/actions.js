export default {
    async loadPosts(context) {
        console.log("loadPosts: ", context.eventDo)
        context.eventDo = "loadPosts"
        try {
            const response = await fetch("http://localhost:3000/posts/getAllPost", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            const posts = await response.json();
            console.log("posts: ", posts)
            context.commit("loadPosts", posts);

        } catch (error) {
            console.error('Failed to fetch posts:', error);
        }     
    },

    async loadComments(context, id) {
        context.eventDo = "loadComments"
        try {
            const response = await fetch(`http://localhost:3000/posts/${id}/getcomment`,{
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            })
            const comments = await response.json()
            context.commit("loadComments", comments)
        } catch (error) {
            
        }
    },

    async loadDiscussions(context) {
        console.log("loadDicussions: ", context.eventDo)
        context.eventDo = "loadDiscussions"
        try {
            const response = await fetch(`http://localhost:3000/auth/getGroupUser/${context.state.user.id}`, {
                method: "GET",
                headers: {
                    "Content-Type": "application/json"
                }
            });

            const discussions = await response.json();
            context.commit("loadDiscussions", discussions);
        } catch (error) {
            console.error('Failed to fetch discussions:', error);
        }
    },

    async createPost(context, payload) {
        context.eventDo = "createPost"
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
            context.commit("createPost", payload)
        } catch (error) {
            console.error('Failed to create post:', error);   
        }
    },
}