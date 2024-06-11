import PubSub from "./pubsub.js"

export default class Store {
    constructor(params) {
        this.events = new PubSub()
        this.state = {}
        this.actions = params.actions || {}
        this.mutations = params.mutations || {}
        this.status = 'resting'

        this.state = new Proxy(params.state || {}, {
            set: function(state, key, value) {
                let self = this
                state[key] = value
                self.events.publish('stateChange', self.state)

                if (self.status !== 'mutation') {
                    console.error(`You should use a mutation to set ${key}`)
                }

                self.status = 'resting'
                return true
            }
        })
    }

    dispatch(actionKey, payload) {
        let self = this
        if (self.actions[actionKey] !== 'function') {
            console.error(`Action ${actionKey} doesn't exist.`)
            return false
        }

        self.status ='action'
        self.actions[actionKey](self, payload)
        return true
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