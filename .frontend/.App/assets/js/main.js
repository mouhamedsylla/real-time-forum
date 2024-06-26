import Post from "./components/posts.js";
import Comment from "./components/comments.js";
import store from "./index.js";
const posts = new Post()
const comments = new Comment()


const postCreate = document.getElementById("create-post")
const postClose = document.getElementById("post-close")
const postAdd_popup = document.querySelector(".add-post")
const postAdd = document.getElementById("post-btn")

console.log(postCreate.checked)

postCreate.addEventListener("change", () => {
    if (postCreate.checked) {
        postAdd_popup.classList.remove("close")
        posts.createPost(postAdd, postAdd_popup)
        postClose.checked = false
    }
})

postClose.addEventListener("change", () => {
    if (postClose.checked) {
        postAdd_popup.classList.add("close")
        postCreate.checked = false
    }
})


store.dispatch("loadPosts")

window.addEventListener("load", () => {
    setTimeout(() => {
        store.state.posts.forEach(post => {
            const commentInput = document.getElementById(`comment-input-${post.Id}`)
    
            commentInput.addEventListener("keypress", (e) => {
                if (e.key == 'Enter') {
                    const input = commentInput.value
                    comments.createComment(post.Id, input)
                }
            })
        })
    }, 0)
})
