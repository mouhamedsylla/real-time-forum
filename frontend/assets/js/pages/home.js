import Page from "./pages.js"
import api from "../../index.js"
import Post from "../components/post.js"
import Comment from "../components/comment.js"
import { parseJwt, alert_token_expire } from "../utils/utils.js"
import Discussion from "../components/discussions.js"
export default class Home extends Page {
    constructor() {
        super("home")
        this.posts = new Post()
        this.comments = new Comment()
        this.discussionsList = new Discussion()
    }

    async initClient() {
        const token_payload = parseJwt(document.cookie)
        if (!token_payload) {
            alert_token_expire()
            return false
        }
        await api.get("/auth/getUsers", { userId: token_payload.id}).then(data => {
            api.setClient(data)
        })
        return true
    }

    async renderComponents() {
        const postsTarget = document.querySelector(".feeds")
        this.posts.setElementTarget(postsTarget)
        await this.posts.render()
        .then(() => {
            this.posts.apiPost.posts.forEach(async post => {
                await this.comments.render(post.Id)
            })
            console.log("Niow")
            this.comments.bindInput()
        })
        .then(async () => {
            this.discussionsList.setTargetElement(document.querySelector(".discussions"))
            await this.discussionsList.render()
        })
    }

    async getHTML() {
        await this.initClient()
        return `
            <nav>
                <div class="container">
                    <h2 class="log">
                        Forum
                    </h2>
                    <div class="search-bar">
                        <i class="uil uil-search"></i>
                        <input type="search" name="" id="" placeholder="search posts">
                    </div>
                    <div class="create">
                        <label class="btn btn-primary" for="create-post">Create</label>
                        <div class="profile-photo">
                            <img src="./frontend/assets/profile.jpg" alt="profile">
                        </div>
                    </div>
                </div>
            </nav>
            <!-- -------------------- MAIN -------------------- -->
            <main>
                <div class="container">
                    <!-- -------------------- LEFT -------------------- -->
                    <div class="left">
                        <a href="" class="profile">
                            <div class="profile-photo">
                                <img src="./frontend/assets/profile2.jpg" alt="">
                            </div>
                            <div class="handle">
                                <h4>${api.client.firstName} ${api.client.lastName}</h4>
                                <p class="text-muted">
                                    @${api.client.nickname}
                                </p>
                            </div>
                        </a>
                        <!-- -------------------- SIDEBAR -------------------- -->
                        <div class="sidebar">
                            <a href="" class="menu-item active">
                                <span><i class="uil uil-home"></i></span><h3>Home</h3> 
                            </a>
                            <a href="" class="menu-item">
                                <span><i class="uil uil-envelope"><small class="notification-count">6</small></i></span><h3>Message</h3>  
                            </a>
                            <a href="" class="menu-item">
                                <span><i class="uil uil-bell"></i></span><h3>Notification</h3>
                            </a>
                            <a href="" class="menu-item">
                                <span><i class="uil uil-palette"></i></span><h3>Theme</h3> 
                            </a>
                        </div>
                        <input type="checkbox" name="" id="create-post" class="btn-check">
                        <label for="create-post" class="btn btn-primary">Create Post</label>
                    </div>
                    <!-- -------------------- MIDDLE -------------------- -->
                    <div class="middle">
                        <div class="feeds">
                            <!-- adding posts and comments HERE -->
                        </div>
                    </div>
                    <!-- -------------------- RIGHT -------------------- -->
                    <div class="right">
                        <div class="messages discussions">
                            <div class="heading">
                                <h4>Discussions</h4><i class="uil uil-edit"></i>
                            </div>
                            <!-- adding discussions HERE -->
                        </div>
                    </div>
                </div>
            </main>

            <div class="add-post close">
                <div class="card">
                    <div class="postAdd-head">
                        <h2>Add Post</h2>
                        <input type="checkbox" name="" id="post-close" class="btn-check" />
                        <label for="post-close" id="btn-close">
                            <i class="fa-solid fa-xmark"></i>
                        </label>
                    </div>

                    <label for="post-title">Title</label>
                    <input type="text" name="" id="post-title">
                    <label for="imageInput">Image</label>
                    <input type="file" id="post-image" name="image" accept="image/*" required>
                    <label for="post-content">Content</label>
                    <textarea name="" id="post-content"></textarea>
                    <label class="btn btn-primary" id="post-btn">Post</label>
                </div>
            </div>
        `
    }
}