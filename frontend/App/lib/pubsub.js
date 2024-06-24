export default class PubSub {
    constructor() {
        this.events = {}
    }

    subscribe(event, callback) {
        // check event property in events object and added if does't exist
        if (!this.events.hasOwnProperty(event)) {
            this.events[event] = []
        }
        // add callback to event array
        this.events[event].push(callback)
    }

    publish(event, data = {}) {
        if (!this.events.hasOwnProperty(event)) {
            console.log("event not exist", event)
            return []
        }

        console.log("event exist", event)
        console.log("events: ", this.events[event])

        return this.events[event].map(callback => callback(data))
    }
}