import Component from '../lib/state/component.js'
import store from '../index.js'
import Comments from './comments.js'

export default class Post extends Component {
    constructor() {
        super({
            store,
            subscribe: ["posts"],
            element: document.querySelector(".feeds")
        })
        this.postBtn = null
    }

    createPost(postButton, popup) {
        this.postBtn = postButton
        this.postBtn.addEventListener("click", () => {
            const Title = document.getElementById("post-title").value
            const Content = document.getElementById("post-content").value
            const Image = document.getElementById("post-image").files[0]

            store.dispatch("createPost", { Title, Image, Content })

            const reader = new FileReader()
            reader.onload = () => {
                const Image = reader.result.replace(/^data:image\/[a-z]+;base64,/, "")
                store.dispatch("addPost", { Title, Image, Content })
            }

            reader.readAsDataURL(Image)
            popup.classList.add("close")
        })
    }

    render() {
        if (store.state.posts.length === 0) {
            return
        }

        this.element.innerHTML = `
            ${store.state.posts.map(post => {
            let post_component = `
                <div class="feed">
                <div class="head">
                    <div class="user">
                        <div class="profile-photo">
                            <img src="./assets/profile2.jpg">
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
                    <span><img src="./assets/images/profile-10.jpg"></span>
                    <span><img src="./assets/images/profile-4.jpg"></span>
                    <span><img src="./assets/images/profile-15.jpg"></span>
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
                        <img src="./assets/profile2.jpg" alt="" />
                        <input type="text" placeholder="Add a comment..." id="comment-input-${post.Id}" />
                    </div>
                </div>    
            </div> `
            return post_component
        }).join("")}`


        store.state.posts.forEach(post => {
            store.dispatch("loadComments", post.Id)
        });

    }


}