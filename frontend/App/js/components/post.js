import Component from "../../lib/component.js"
import store from "../index.js"

export default class Post extends Component {
    constructor() {
        super({
            store,
            element: document.querySelector(".feeds"),
            events: {
                'loadPosts': () => this.render(),
                'createPost': () => this.addPost(this.post)
            }
        })
        this.post = null
        this.postBtn = null
    }

    createPostHTML(post) {
        const elem = document.createElement("div")
        elem.classList.add("feed")
        let post_component = `
                <div class="head">
                    <div class="user">
                        <div class="profile-photo">
                            <img src="./App/assets/profile2.jpg">
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
                    <span><img src="./App/assets/images/profile-10.jpg"></span>
                    <span><img src="./App/assets/images/profile-4.jpg"></span>
                    <span><img src="./App/assets/images/profile-15.jpg"></span>
                    <p>Liked by <b>Ahmed Sylla</b> and <b>253 others</b></p>
                </div>
                <div class="caption">
                    <p><b>Daenerys</b> ${post.Content} <span class="hash-tag">#lifestyle</span></p>
                </div>
                <div class="post-comment ">
                    <div class="head-comment">
                        <div class="name">Comments</div>
                        <label for="comment-${post.Id}" id="comment-btn3">
                          <i class="fa-solid fa-xmark"></i>
                        </label>
                    </div>
                    <div class="comment-container" data-comment-post-${post.Id}>

                    </div>
                    <div class="new-comment">
                        <img src="./App/assets/profile2.jpg" alt="" />
                        <input type="text" placeholder="Add a comment..." id="comment-input-${post.Id}" />
                    </div>
                </div>`
            elem.innerHTML = post_component 
            return elem
    }

    render() {
        if (store.state.posts.length === 0) { return }
        store.state.posts.forEach(post => {
            this.element.appendChild(this.createPostHTML(post))
        })

        store.state.posts.forEach(post => {
            (async (post) => {
                store.dispatch("loadComments", post.Id)
            })(post)
        });
    }

    addPost(post) {
        console.log("post added")
        this.element.appendChild(this.createPostHTML(post))
    }

    createPost(postButton, popup) {
        this.postBtn = postButton
        this.postBtn.addEventListener("click", () => {
            const Title = document.getElementById("post-title").value
            const Content = document.getElementById("post-content").value
            const Image = document.getElementById("post-image").files[0]

            const payload = { Title, Image, Content }

            const reader = new FileReader()
            reader.onload = () => {
                const Image = reader.result.replace(/^data:image\/[a-z]+;base64,/, "")
                this.post = { Title, Image, Content }
            }

            (async () => {
                store.dispatch("createPost", payload)
            })()

            reader.readAsDataURL(Image)
            popup.classList.add("close")
        })
    }
}