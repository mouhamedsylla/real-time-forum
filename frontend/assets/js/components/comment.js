import CommentAPI from "../api/comments.js"
import api from "../../index.js" 
import { session_expired } from "../utils/other.js"
import { alert_token_expire } from "../utils/alert.js"

export default class Comment {
    constructor() {
        this.apiComment = new CommentAPI()
        this.currentComment = null
        this.targetElement = null
    }

    bindInput() {
        const container = document.getElementById("app")

        container.addEventListener('keypress', async (event) => {
            if (event.target.classList.contains('all__input')) {
                if (session_expired()) {
                    alert_token_expire()
                    return
                } else if (event.key === 'Enter' && event.target.value.trim() !== ""){
                     const id = event.target.id.split('-')[2]
                    await this.addComment({
                        Comment: event.target.value,
                        Post_id: +id,
                        User_id: api.client.Id
                    }, +id)
                    event.target.value = ""
                }
            }
        })
    }

    createCommentHTML(comment, username, lastId) {
        const commentId = comment.Id ? comment.Id : lastId
        const elem  = document.createElement("div")
        elem.classList.add("Comment")
        elem.setAttribute("id", commentId)
        elem.innerHTML = `
            <img src="./frontend/assets/images/profile-1.jpg" alt="" />
			<span>
                <strong>${username}</strong>
				${comment.Comment}
			</span>
        `
        return elem
    }

    async render() {
        try {
            await this.apiComment.getComments()
            for (const comment of this.apiComment.comments) {
                const user = await this.apiComment.getUserByCommentId(comment.User_id)
                this.targetElement = document.querySelector(`[data-comment-post-${comment.Post_id}]`)
                this.targetElement.appendChild(this.createCommentHTML(comment, user.nickname))
            }
        } catch (error) {
            console.error("Error in render:", error)
        }
    }

    async addComment(comment, idPost) {
        const lastId = await this.apiComment.postComment(comment, idPost)
        this.targetElement = document.querySelector(`[data-comment-post-${idPost}]`)
        this.targetElement.appendChild(this.createCommentHTML(comment, api.client.nickname, lastId))
    }

}