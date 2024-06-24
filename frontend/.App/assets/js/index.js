import actions from './actions.js'
import Store from './lib/state/store.js'
import mutations from './mutations.js'
import state from './state.js'

const store = new Store({
    actions,
    mutations,
    state
})


export default store