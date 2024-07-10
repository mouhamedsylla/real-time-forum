import Page from "./pages.js";
import api from "../../index.js";
import { alert_token_expire } from "../utils/alert.js";
import { session_expired, parseJwt } from "../utils/other.js";
import Notification from "../components/notification.js";
import Message from "../components/message.js";
import Post from "../components/post.js";
import Comment from "../components/comment.js";



export default class Home extends Page {
    constructor() {
        super();
        this.setTitle('Home');
        this.posts = new Post()
        this.comments = new Comment()
        this.notifications = new Notification()
        this.message = new Message()
    }

    async initClients() {
        if (session_expired()) {
            alert_token_expire()
            return false
        }

        const token_payload = parseJwt(document.cookie)
        await api.get("/auth/getUsers", { userId: token_payload.id }).then(data => {
            api.setUserClient(data)
        })        

        
        api.otherClient.forEach(user => user.status = "offline")
        return true
    }

    sittengLogout() {
        const logout = document.getElementById("logout-pointer")
        logout.addEventListener("click", () => {
            document.cookie = 'forum=; Max-Age=0; path=/;'
            window.location.href = "/login"
        })
    }

    createDiscussionProfile(user) {
        const elem = document.createElement("div")
        elem.classList.add("message")
        elem.setAttribute("id", `discussion-${user.Id}`)
        const content = `
                <div class="profile-photo">
                    <img src="./frontend/assets/images/profile-${user.Id}.jpg" alt="">
                    <div class="status" id="${user.Id}"></div>
                </div>
                <div class="message-body">
                    <h5>${user.firstName} ${user.lastName}</h5>
                </div>`
        elem.innerHTML = content
        this.message.handleMessageDiscussion(user, elem)
        return elem
    }

    async renderDiscussionProfile() {
        try {
            api.discussionsUsers = await api.get(`/auth/getGroupUser/${api.client.Id}`) || []
            
            const allUsersData = await api.get("/auth/getUsers");
            const discussionUserIds = new Set(api.discussionsUsers.map(user => user.Id))
    
            const otherUsers = allUsersData
                .filter(user => !discussionUserIds.has(user.Id) && user.Id !== api.client.Id)
    
            api.setOtherClient(otherUsers)
    
            const discussions = document.querySelector(".discussions")
            discussions.innerHTML = ''
            api.sortUsers()
            api.discussionsUsers.concat(api.otherClient)
                .forEach(user => { 
                    discussions.appendChild(this.createDiscussionProfile(user))
                })
    
        } catch (error) {
            console.error("Error in renderDiscussionProfile:", error)
        }
    }

    
    async renderComponents() {
        await this.renderDiscussionProfile()
        await this.notifications.initNotification()

        this.posts.setElementTarget(document.querySelector(".feeds"))
        this.posts.bindButton()
        try {
            await this.posts.render()
            this.comments.bindInput()
            this.sittengLogout()
        } catch (error) {
            console.error("Error rendering posts or comments:", error)
        }

        try {
            await this.comments.render()
        } catch (error) {
            console.error("Error rendering comments:", error)
        }
    }


    async getHTML() {
        await this.initClients()
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
                        
                        <i class="uil uil-sun icon" id="toggleIcon"></i>

                        <i class="fas fa-comments" id=chat-icon></i>
                        
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
                        <div class="logout">
                            <iframe src="https://lottie.host/embed/0d3e821e-608e-4c51-9ccc-f25d052d11a5/vYSfT4mQqm.json"></iframe>
                            <h4 id="logout-pointer">Logout</h4>
                        </div>
                        <!-- -------------------- SIDEBAR -------------------- -->
                        <div class="sidebar">
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
                        <div class="right-profile">
                            <a href="" class="profile">
                                <div class="handle">
                                    <h4>${api.client.firstName} ${api.client.lastName}</h4>
                                    <p class="text-muted">
                                        @${api.client.nickname}
                                    </p>
                                </div>
                            </a>
                        </div>
                        <div class="heading">
                                <h4>Discussions</h4><i class="uil uil-edit"></i>
                        </div>
                        <div class="messages discussions">
                            
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
                    <input type="file" id="post-image" name="image" accept="image/jpeg" required>
                    <label for="post-content">Content</label>
                    <textarea name="" id="post-content"></textarea>

                    <label for="list">Categories</label>
                    <div class="list">
                        <input class="categorie__check" type="checkbox" name="role" id="opt1" />
                        <label for="opt1"> Science </label>
                        <input class="categorie__check" type="checkbox" name="role" id="opt2" />
                        <label for="opt2"> Informatique </label>
                        <input class="categorie__check" type="checkbox" name="role" id="opt3" />
                        <label for="opt3"> Litterature </label>
                        <input class="categorie__check" type="checkbox" name="role" id="opt4" />
                        <label for="opt4"> Religion </label>
                    </div>

                    <label class="btn btn-primary" id="post-btn">Post</label>
                </div>
            </div>
        `
    }
}