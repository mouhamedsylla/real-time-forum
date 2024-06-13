import Component from "../lib/state/component";

export default class Comments extends Component {
    constructor(postId) {
        super({
            store,
            element: document.querySelector(`[data-comment-post-${postId}]`)
        })
    }

    render() {
        if (store.state.comments.length === 0) {
            this.element.innerHTML = `<h2>No comments yet!</h2>`
        }

        this.element.innerHTML = `
            ${store.state.comments.map(comment => {
                return `
                <div class="comment" id="${comment.id}">
                    <img src="" alt="" />
                    <span>
                        ${comment.comment}
                        <div class="desc">2m ago <span>Reply</span></div>
                    </span>
                    <i class="fa-regular fa-heart"></i>
                </div>
                `
            }).join("")}

        `
    }
}

// const commentobj = {
//     id: 1,
//     created_at: "2021-04-02T15:00:00",
//     comment: "This is a comment",
//     post_id: 1,
//     likes: 0,
//     dislikes: 0
// }