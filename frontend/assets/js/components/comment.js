import CommentAPI from "../api/comments.js"
import api from "../../index.js"   

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
                const id = event.target.id.split('-')[2]
                if (event.key === 'Enter') {
                    await this.addComment({
                        Comment: event.target.value,
                        Post_id: parseInt(id),
                        User_id: api.client.Id
                    }, id)
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
                <b>${username}</b>
				${comment.Comment}
			</span>
        `
        return elem
    }

    async render() {
        try {
            await this.apiComment.getComments()
            this.apiComment.comments.forEach(async comment => {
                this.targetElement = document.querySelector(`[data-comment-post-${comment.Post_id}]`)
                const user = await this.apiComment.getUserByCommentId(comment.User_id)
                this.targetElement.appendChild(this.createCommentHTML(comment, user.nickname))
            })
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