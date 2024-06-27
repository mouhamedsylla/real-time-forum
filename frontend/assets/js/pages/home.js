import Page from "./pages.js"
import api from "../../index.js"
export default class Home extends Page {
    constructor() {
        super("home")
    }

    getHTML() {
        console.log(api.client)
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
                        <h4>Daenerys</h4>
                        <p class="text-muted">
                            @daeny
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
                <div class="messages">
                    <div class="heading">
                        <h4>Messages</h4><i class="uil uil-edit"></i>
                    </div>
                    <div class="message">
                        <div class="profile-photo">
                            <img src="./frontend/assets/images/profile-7.jpg" alt="">
                            <div class="active"></div>
                        </div>
                        <div class="message-body">
                            <h5>John Snow</h5>
                            <p class="text-muted">Just woke up bruh</p>
                        </div>
                    </div>
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