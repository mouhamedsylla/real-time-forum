import Store from "./store.js"

export default class Component {
    constructor(props = {}) {

        if (props.store instanceof Store) {
            if (props.hasOwnProperty("events")) {
                Object.entries(props.events).forEach(([event, callback]) => {
                    props.store.pubsub.subscribe(event, callback)
                })
            }
        }

        console.log("props: ", props.store.pubsub)

        if (props.hasOwnProperty("element")) {
            this.element = props.element
        }
    }
}