import Store from "./store.js"

export default class Component {
    constructor(props = {}) {
        this.render = this.render || function() {}

        if (props.store instanceof Store) {
            if (props.subscribe) {
                props.subscribe.forEach(event => {
                    props.store.events.subscribe(`stateChange:${event}`, () => this.render())
                })
            }
        }

        if (props.hasOwnProperty("element")) {
            this.element = props.element
        }
    }
}