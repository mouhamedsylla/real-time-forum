import store from "../index.js"
import Component from "../../lib/component.js"

export default class Discussions extends Component {
    constructor() {
        super({
            store,
            element: document.querySelector(".discussions"),
            events: {
                'loadDiscussions': () => this.render(),
            }
        })
    }

    createDiscussionHTML(discussion) {
        const elem = document.createElement("div")
        elem.classList.add("message")
        const content = `
                <div class="profile-photo">
                    <img src="./App/assets/images/profile-7.jpg" alt="">
                    <div class="active"></div>
                </div>
                <div class="message-body">
                    <h5>${discussion.firstName} ${discussion.lastName}</h5>
                </div>`
        elem.innerHTML = content
        return elem
    }

    render() {
        if (store.state.discussions.length === 0) {
            this.element.innerHTML = "<p>No discussions</p>"
        }
        store.state.discussions.forEach(user => {
            this.element.appendChild(this.createDiscussionHTML(user)) 
        });  
    }
}