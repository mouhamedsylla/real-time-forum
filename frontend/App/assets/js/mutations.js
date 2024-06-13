export default {
   loadPosts(state, payload) {
      state.posts = payload
      return state
   },

   createPost(state) {
    return state
   },

   addPost(state, payload) {
    state.posts.push(payload)
    return state
   }
}