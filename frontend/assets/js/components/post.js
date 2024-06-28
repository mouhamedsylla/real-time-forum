import PostAPI from "../api/posts.js";

export default class Post {
    constructor() {
        this.apiPost = new PostAPI()
        this.currentPost = null
        this.elementTarget = {}
    }

    setElementTarget(element) {
        this.elementTarget = element
    }

    bindButton() {
        const postCreate = document.getElementById("create-post")
        const postClose = document.getElementById("post-close")
        const postAdd_popup = document.querySelector(".add-post")
        const postAdd = document.getElementById("post-btn")

        postAdd.addEventListener("click", async () => {
            await this.addPost(this.getPostInput())
            postAdd_popup.classList.add("close")
        })


        postCreate.addEventListener("change", () => {
            if (postCreate.checked) {
                postAdd_popup.classList.remove("close")
                postClose.checked = false
            }
        })

        postClose.addEventListener("change", () => {
            if (postClose.checked) {
                postAdd_popup.classList.add("close")
                postCreate.checked = false
            }
        })
    }

    getPostInput() {
        const Image = document.getElementById("post-image").files[0]
        const Title = document.getElementById("post-title").value
        const Content = document.getElementById("post-content").value

        const reader = new FileReader()
        reader.onload = () => {
            const Image = reader.result.replace(/^data:image\/[a-z]+;base64,/, "")
            this.currentPost = { Title, Image, Content }
        }
        reader.readAsDataURL(Image)
        return { Title, Image, Content }
    }

    createPostHTML(post) {
        const elem = document.createElement("div")
        elem.classList.add("feed")
        let post_component = `
                <div class="head">
                    <div class="user">
                        <div class="profile-photo">
                            <img src="./frontend/assets/profile2.jpg">
                        </div>
                        <div class="ingo">
                            <h3>Daenerys</h3>
                            <small>New York, 15 MINUTES AGO</small>
                        </div>
                    </div>
                </div>
                <div class="photo">
                    <img src="data:image/jpeg;base64,${post.Image}">
                </div>
                <input type="checkbox" name="" id="like-${post.Id}" class="btn-check like-check" />
                <input type="checkbox" name="" id="comment-${post.Id}" class="btn-check comment-check" />
                <div class="action-buttons">
                    <div class="interaction-buttons">
                        <label for="like-${post.Id}" id="like-btn"></label>
                        <label for="comment-${post.Id}" id="comment-btn"></label>
                        <span><i class="uil uil-share-alt"></i></span>
                    </div>
                    <div class="bookmark">
                        <span><i class="uil uil-bookmark-full"></i></span>
                    </div>
                </div>
                <div class="liked-by">
                    <span><img src="./frontend/assets/images/profile-10.jpg"></span>
                    <span><img src="./frontend/assets/images/profile-4.jpg"></span>
                    <span><img src="./frontend/assets/images/profile-15.jpg"></span>
                    <p>Liked by <b>Ahmed Sylla</b> and <b>253 others</b></p>
                </div>
                <div class="caption">
                    <p><b>Daenerys</b> ${post.Content} <span class="hash-tag">#lifestyle</span></p>
                </div>
                <div class="post-comment">
                    <div class="head-comment">
                        <div class="name">Comments</div>
                        <label for="comment-${post.Id}" id="comment-btn3">
                          <i class="fa-solid fa-xmark"></i>
                        </label>
                    </div>
                    <div class="comment-container" data-comment-post-${post.Id}>

                    </div>
                    <div class="new-comment">
                        <img src="./frontend/assets/profile2.jpg" alt="" />
                        <input type="text" placeholder="Add a comment..." id="comment-input-${post.Id}" />
                    </div>
                </div>`
        elem.innerHTML = post_component
        return elem
    }

    async render() {
        await this.apiPost.getPosts()
        this.apiPost.posts.forEach(post => {
            this.elementTarget.appendChild(this.createPostHTML(post))
        })
        this.bindButton()
    }

    async addPost(post) {
        await this.apiPost.createPost(post)
        this.elementTarget.appendChild(this.createPostHTML(this.currentPost))
        this.currentPost = null
    }
}