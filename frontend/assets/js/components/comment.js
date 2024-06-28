import CommentAPI from "../api/comments.js"

export default class Comment {
    constructor() {
        this.apiComment = new CommentAPI()
        this.currentComment = null
        this.targetElement = null
    }

    createCommentHTML(comment) {
        const elem  = document.createElement("div")
        elem.classList.add("comment")
        elem.setAttribute("id", comment.Id)
        elem.innerHTML = `
            <img src="./frontend/assets/images/profile-${comment.Id}.jpg" alt="" />
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
}