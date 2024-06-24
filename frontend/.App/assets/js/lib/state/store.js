import PubSub from "./pubsub.js"

export default class Store {
    constructor(params) {
        let self = this
        self.events = new PubSub()
        self.state = {}
        self.actions = params.actions || {}
        self.mutations = params.mutations || {}
        self.status = 'resting'

        self.state = new Proxy(params.state || {}, {
            set: function(state, key, value) {
                state[key] = value

                console.log(`stateChange: ${key}: ${value}`)
            
                self.events.publish(`stateChange:${key}`, self.state)

                if (self.status !== 'mutation') {
                    console.log(`You should use a mutation to set ${key}`)
                }

                self.status = 'resting'
                return true
            }
        })
    }

    dispatch(actionKey, payload) {
        let self = this
        if (typeof self.actions[actionKey] !== 'function') {
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