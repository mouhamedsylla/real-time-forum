import PubSub from "./pubsub.js"

export default class Store {
    constructor(params) {
        let self = this
        self.actions = params.actions || {}
        self.mutations = params.mutations || {}
        self.events = params.events
        self.eventDo = ""
        self.status = 'resting'
        self.state = {}
        self.pubsub = new PubSub()
        self.state = new Proxy((params.state || {}), {
            set (state, key, value) {
                if (!self.events[key].includes(self.eventDo)) {
                    console.warn(`No event ${self.eventDo} for ${key}`)
                    return true
                }

                state[key] = value
                self.pubsub.publish(self.eventDo, self.state)
                if (self.status !== "mutation") {
                    console.warn(`You should use a mutation to set ${key}`)
                }

                self.status = "resting"
                return true
            }
        })
    }

    // dispatch(actionKey, payload) {
    //     let self = this
    //     if (typeof self.actions[actionKey] !== 'function') {
    //         console.error(`Action ${actionKey} doesn't exist.`)
    //         return false
    //     }

    //     self.status = "action"
    //     self.actions[actionKey](self, payload)
    //     return true
    // }

    async dispatch(actionKey, payload) {
        if (typeof this.actions[actionKey] !== 'function') {
            console.error(`Action ${actionKey} doesn't exist.`);
            return false;
        }

        // Set status to indicate an action is being dispatched
        this.status = 'action';

        try {
            // Await the action to complete
            await this.actions[actionKey](this, payload);
        } catch (error) {
            console.error(`Error in action ${actionKey}:`, error);
        }

        // Reset status
        this.status = 'resting';
        return true;
    }

    commit(mutationKey, payload) {
        let self = this
        if (typeof self.mutations[mutationKey] !== 'function') {
            console.error(`Mutation ${mutationKey} doesn't exist`)
            return false
        }

        self.status = "mutation"
        let newState = self.mutations[mutationKey](self.state, payload)
        self.state = Object.assign(self.state, newState)
        return true
    }
}