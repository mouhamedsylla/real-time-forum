import Post from "./components/posts.js";
const posts = new Post()

posts.render()

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