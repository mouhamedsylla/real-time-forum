import actions from "./state/actions.js";
import Store from "../lib/store.js";
import mutations from "./state/mutations.js";
import state from "./state/state.js";

const events = {
    user: [],
    posts: ["loadPosts", "createPost"],
    comments: ["loadComments"],
    discussions: ["loadDiscussions"]
}

const store = new Store({
    state,
    actions,
    mutations,
    events
})

export default store