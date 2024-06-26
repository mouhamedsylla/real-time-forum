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

   createPost(state) {
    return state
   },

   createComment(state) {
      return state
   },

   addPost(state, payload) {
    state.posts.push(payload)
    return state
   },

   addComment(state, payload) {
      let newComments = {...state.comments}
      let id = payload.Post_id
      newComments[id].push(payload)
      state.comments = newComments
      return state
   },

   setCurrentPost(state, payload) {
    state.currentPost = payload
    return state
   }
}