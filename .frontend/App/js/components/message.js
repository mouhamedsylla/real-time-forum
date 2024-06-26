import store from "../index.js";

export default class Message extends Component {
    constructor() {
        super({
            store,
            element: document.querySelector(".right"),
            events: {
                'loadMessages': () => this.render(),
            }
        })
        this.users = {}
    }

    setUsers(sender, receiver) {
        this.users = { sender, receiver }
    }

    createMessageHTML(message) {
        
    }

    render() {
        const elem = document.createElement("div")
        elem.classList.add("chat")
        elem.innerHTML = `
            <div class="contact bar">
                <div class="profile-photo pic">
                    <img src="./App/assets/images/profile-7.jpg" alt="">
                    <div class="active"></div>
                </div>
                <div class="name">
                    John Snow
                </div>
                <div class="seen">
                    Today at 12:56
                </div>
            </div>
            <div class="chat-messages">
                <div class="time">
                    Today at 11:41
                </div>
            </div>
            <div class="input">
                <i class="far fa-laugh-beam"></i><input placeholder="Type your message here!" type="text" />
            </div>
        `
    }
}