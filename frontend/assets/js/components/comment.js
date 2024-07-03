import api from "../../index.js"
import CommentAPI from "../api/comments.js"

export default class Comment {
    constructor() {
        this.apiComment = new CommentAPI()
        this.currentComment = null
        this.targetElement = null
    }

    bindInput() {
        const inputs = document.querySelectorAll(".all__input")
        console.log("input: ",inputs)
        inputs.forEach(input => {
            input.addEventListener('keypress', async (event) => {
                const id = event.target.id.split('-')[2]
                console.log(event.key)
                if (event.key === 'Enter') {
                    await this.addComment({
                        Comment: input.value,
                        Post_id: parseInt(id)
                    
                    }, id)
                    input.value = ""
                }
            })
        });
    }

    getInputComment() {
        this.input.addEventListener('keydown', (event) => {
            if (event.key === 'Enter') {
                
            }
        })
    }

    createCommentHTML(comment, lastId) {
        const commentId = comment.Id ? comment.Id : lastId
        const elem  = document.createElement("div")
        elem.classList.add("comment")
        elem.setAttribute("id", commentId)
        elem.innerHTML = `
            <img src="./frontend/assets/images/profile-1.jpg" alt="" />
			<span>
				${comment.Comment}
				<div class="desc">2m ago <span>Reply</span></div>
			</span>
			<i class="fa-regular fa-heart"></i>
        `
        return elem
    }

    async render(idPost) {
        await this.apiComment.getComments(idPost)
        .then(() => {
            this.targetElement = document.querySelector(`[data-comment-post-${idPost}]`)
            this.apiComment.comments[idPost].forEach(comment => {
            this.targetElement.appendChild(this.createCommentHTML(comment))
        })
        })
    }

    async addComment(comment, idPost) {
        const lastId = await this.apiComment.postComment(comment, idPost)
        const elem = this.createCommentHTML(comment, lastId)
        this.targetElement.appendChild(elem)
    }

}