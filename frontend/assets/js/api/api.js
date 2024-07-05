export default class Api {
    constructor() {
        this.baseURL = 'http://localhost:3000'
        this.client = null
        this.otherClient = []
        this.discussionsUsers = []
    }

    setUserClient(client) {
        this.client = client
    }

    setOtherClient(clients) {
        this.otherClient = clients
    }

    sortUsers() {
        this.otherClient.sort((a, b) => {
            return a.firstName.localeCompare(b.firstName)
        })
    }

    getUserById(id) {
        return this.otherClient.find(user => user.Id === id) || this.discussionsUsers.find(user => user.Id === id)
    }

    getUsernameDiscussion(id) {
        const user = this.discussionsUsers.find(user => user.Id === id)   
        return user.nickname  || null
    }


    async get(endpoint, params = {}) {
        try {
            const url = new URL(`${this.baseURL}${endpoint}`)
            Object.keys(params).forEach(key => url.searchParams.append(key, params[key]))

            const response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                }
            })

            if (response.ok) {
                console.log('Promise resolved and HTTP status is successful')
                return await response.json()
            } else {
                if (response.status === 404) throw new Error('404 Not Found')
                if (response.status === 401) throw new Error('401 Unauthorized')
                if (response.status === 500) throw new Error('500 Internal Server Error')
                if (response.status === 405) throw new Error('405 Method Not Allowed')

                throw new Error(response.status)
            }
        } catch (error) {
            console.error('Promise rejected:', error)
        }
    }

    async post(endpoint, data = {}) {
        try {
            const url = new URL(`${this.baseURL}/${endpoint}`)
            const options = {
                method: 'POST',
                headers: {}
            }

            if (data instanceof FormData) {
                options.body = data
            } else {
                options.headers['Content-Type'] = 'application/json'
                options.body = JSON.stringify(data)
            }

            const response = await fetch(url, options)

            if (response.ok) {
                console.log('Promise resolved and HTTP status is successful')
                return await response.json()
            } else {
                if (response.status === 404) throw new Error('404 Not Found')
                if (response.status === 401) throw new Error('401 Unauthorized')
                if (response.status === 500) throw new Error('500 Internal Server Error')
                if (response.status === 405) throw new Error('405 Method Not Allowed')

                throw new Error(response.status)
            }
        } catch (error) {
            console.error('Promise rejected:', error)
        }
    }
}