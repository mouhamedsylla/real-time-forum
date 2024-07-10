import api from "../../index.js";
import PostAPI from "../api/posts.js";
import { alert_token_expire } from "../utils/alert.js";
import { session_expired, formatTimeAgo, like, dislike, escapeHtml } from "../utils/other.js";
export default class Post {
    constructor() {
        this.apiPost = new PostAPI()
        this.currentPost = null
        this.elementTarget = {}
        this.postReaction = null
    }

    setElementTarget(element) {
        this.elementTarget = element
    }

    bindButton() {
        const postCreate = document.getElementById("create-post");
        const postClose = document.getElementById("post-close");
        const postAdd_popup = document.querySelector(".add-post");
        const postAdd = document.getElementById("post-btn");

        postAdd.addEventListener("click", async () => {
            const newPost = this.getPostInput()
            if (!newPost) return
            session_expired() ? alert_token_expire() : await this.addPost(newPost)
            postAdd_popup.classList.add("close")
        })

        postCreate.addEventListener("change", () => {
            postAdd_popup.classList.toggle("close", !postCreate.checked);
            postClose.checked = !postCreate.checked;
        });

        postClose.addEventListener("change", () => {
            postAdd_popup.classList.toggle("close", postClose.checked);
            postCreate.checked = !postClose.checked;
        });
    }

    getPostInput() {
        const ImageInput = document.getElementById("post-image");
        const Image = ImageInput.files[0];
        const Title = document.getElementById("post-title").value.trim();
        const Content = document.getElementById("post-content").value.trim();
        let Categories = [];

        if (!Image) {
            alert("Please select an image.");
            return null;
        }
        
        // Check that the title is not empty
        if (!Title) {
            alert("Please enter a title.");
            return null;
        }
    
        // Check that the content is not empty
        if (!Content) {
            alert("Please enter content.");
            return null;
        }
    
        // Check that at least one category is selected
        document.querySelectorAll(".categorie__check").forEach(checkBox => {
            if (checkBox.checked) {
                const label = document.querySelector(`label[for="${checkBox.id}"]`);
                Categories.push(label.textContent);
            }
        });
    
        if (Categories.length === 0) {
            alert("Please select at least one category.");
            return null;
        }
    
        Categories = Categories.map(item => `#${item.trim()}`).join(' ');
    
        const reader = new FileReader();
        reader.onload = () => {
            const Image = reader.result.replace(/^data:image\/[a-z]+;base64,/, "");
            this.currentPost = { Title, Image, Content, Categories };
        };
        
        reader.readAsDataURL(Image);

    
        return { Title, Image, Content, Categories };
    }
    

    handleReaction() {
        const likeBtn = document.querySelectorAll("[reaction]")

        likeBtn.forEach(element => {
            const observer = new MutationObserver((mutationsList) => {
                for (const mutation of mutationsList) {
                    if (mutation.type === 'attributes' && mutation.attributeName === 'style') {
                        const newStyle = window.getComputedStyle(element).color;
                        const nbView = document.getElementById(`likePost-${element.id}`)
                        if (newStyle === like) {
                            this.apiPost.reactionPost(element.id, { Value: "like" })
                            nbView.textContent = +nbView.textContent + 1
                        } else if (newStyle === dislike) {
                            this.apiPost.reactionPost(element.id, { Value: "dislike" })
                            nbView.textContent = +nbView.textContent - 1
                        }
                    }
                }
            })
            observer.observe(element, { attributes: true })
            element.addEventListener("click", () => {
               element.style.color = element.style.color === 'rgb(255, 87, 51)' ? '#3498db' : '#ff5733';
            });
        })
    }

    createPostHTML(post, userPosted, lastId) {
        const postId = post.Id ? post.Id : lastId
        const elem = document.createElement("div")

        const escapedContent = escapeHtml(post.Content)
        const escapedCategories = escapeHtml(post.Categories)
        elem.classList.add("feed")
        let post_component = `
                <div class="head">
                    <div class="user">
                        <div class="profile-photo">
                            <img src="./frontend/assets/images/profile-${userPosted.Id}.jpg">
                        </div>
                        <div class="ingo">
                            <h3>${userPosted.firstName} ${userPosted.lastName}</h3>
                            <small>${formatTimeAgo(post.CreatedAt)}</small>
                        </div>
                    </div>
                </div>
                <div class="photo">
                    <img src="data:image/jpeg;base64,${post.Image}">
                </div>
                <input type="checkbox" name="" id="like-${postId}" class="btn-check like-check"/>
                <input type="checkbox" name="" id="comment-${postId}" class="btn-check comment-check" />
                <div class="action-buttons">
                    <div class="interaction-buttons">
                        <label for="like-${postId}" class="like-btn" id="${postId}" reaction></label>
                        <label for="comment-${postId}" id="comment-btn"></label>
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
                    <p>Liked by <b id="likePost-${postId}">${post.nbLike || 0 }</b> persons</p>
                </div>
                <div class="caption">
                    <p><b>Daenerys</b> ${escapedContent} <span class="hash-tag">${escapedCategories}</span></p>
                </div>
                <div class="post-comment">
                    <div class="head-comment">
                        <div class="name">Comments</div>
                        <label for="comment-${postId}" id="comment-btn3">
                          <i class="fa-solid fa-xmark"></i>
                        </label>
                    </div>
                    <div class="comment-container" data-comment-post-${postId}>

                    </div>
                    <div class="new-comment">
                        <img src="./frontend/assets/profile2.jpg" alt="" />
                        <input class="all__input" type="text" placeholder="Add a comment..." id="comment-input-${postId}"/>
                    </div>
                </div>`
        elem.innerHTML = post_component
        return elem
    }

    async render() {
        try {
            const posts = await this.apiPost.getPosts()
            for (const post of posts) {
                const userPosted = await this.apiPost.getUserByPostId(post.UserId);
                this.elementTarget.appendChild(this.createPostHTML(post, userPosted));
            }

            const reaction = await this.apiPost.getReactionPost()
            const likeBtn = document.querySelectorAll(".like-btn")
            likeBtn.forEach(element => {
                const reactionPost = reaction.find(reaction => reaction.PostId == element.id)
                if (reactionPost) {
                    if (reactionPost.value === "like") {
                        element.style.color = '#ff5733'
                    } else if (reactionPost.value === "dislike") {
                        element.style.color = '#3498db'
                    }
                }
            })
            this.handleReaction();
        } catch (error) {
            console.error("Error while rendering posts: ", error);
        }
    }



    async addPost(post) {
        if (!post) return;
        try {
            await this.apiPost.createPost(post)
            this.elementTarget.appendChild(this.createPostHTML(this.currentPost, api.client, this.apiPost.lastPostId.lastId))
            this.handleReaction()
            this.currentPost = null
        } catch (error) {
            console.error("Error while adding post: ", error)
        }
    }

}