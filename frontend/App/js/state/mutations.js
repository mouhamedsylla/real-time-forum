export default {
    loadPosts(state, payload) {
        state.posts = payload
    },

    loadComments(state, payload) {
        let newComments = {...state.comments}
        let id = payload[0].Post_id
        newComments[id] = payload
        state.comments = newComments
  },

    loadDiscussions(state, payload) {
        state.discussions = payload
    },

    createPost(state, payload) {
        state.posts.push(payload)
        return state
    },
}