import Component from "../../lib/component.js";
import store from "../index.js";

export default class Comments extends Component {
	constructor() {
		super({
			store,
			events: {
                'loadComments': () => this.render(),
            }
		})
	}

	createComment(Post_id, Comment) {
		let payload = {
			Comment,
			Post_id,
			Like: 0,
			Dislike: 0
		}
		store.dispatch("createComment", payload)
		store.dispatch("addComment", payload)
	}

	render() {
        //console.log(store.state.comments)
		Object.entries(store.state.comments).forEach(([key, value]) => {
			let i = 1
			let IdPost = value[0].Post_id
			this.element = document.querySelector(`[data-comment-post-${IdPost}]`)

			this.element.innerHTML = `
			${value.map(comment => {
				let comment_cpnt = `
					<div class="comment" id="${comment.Id}">
			  		<img src="./assets/images/profile-${i}.jpg" alt="" />
			  		<span>
						${comment.Comment}
						<div class="desc">2m ago <span>Reply</span></div>
			  		</span>
			  		<i class="fa-regular fa-heart"></i>
					</div>
				`
				i++
				return comment_cpnt
			}).join("")}`
		})

	}
}

