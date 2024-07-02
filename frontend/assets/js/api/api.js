export default class API {
    constructor() {
        this.baseURL
        this.client
        this.otherUser = {}
    }

    setbaseURL(baseURL) {
        this.baseURL = baseURL
    }

    setClient(client) {
        this.client = Object.assign({}, client) 
    }

    sortUsers() {
        return this.otherUser.sort((a, b) => {
          if (a.status === "online" && b.status === "offline") {
            return -1;
          } else if (a.status === "offline" && b.status === "online") {
            return 1;
          } else if (a.status === "offline" && b.status === "offline") {
            return a.firstName.localeCompare(b.firstName);
          }
          return 0;
        })
    }

    getUserById(id) {
        return this.otherUser.find(user => user.Id === parseInt(id))
    }

    setUserStatus(id, status) {
        this.otherUser.find(user => user.Id === parseInt(id)).status =  status
    }

    async get(endpoint, params = {}) {
        try {
            const url = new URL(`${this.baseURL}${endpoint}`);
            Object.keys(params).forEach(key => url.searchParams.append(key, params[key]));
      
            const response = await fetch(url, {
              method: 'GET',
              headers: {
                'Content-Type': 'application/json'
              }
            });
      
            if (!response.ok) {
              throw new Error(`Error: ${response.statusText}`);
            }
      
            return await response.json();
          } catch (error) {
            console.error('GET request failed:', error);
            throw error;
          }
    }

    async post(endpoint, data = {}) {
        try {
            let options = {
                method: "POST",
                headers: {}
            }

            if (data instanceof FormData) {
                options.body = data
            } else {
                options.headers = {
                    "Content-Type": "application/json"
                }
                options.body = JSON.stringify(data)
            }

            const response = await fetch(`${this.baseURL}${endpoint}`, options);

            if (!response.ok) {
                throw new Error(`Error: ${response.statusText}`);
              }

            return await response.json();
        } catch (error) {
            console.error('POST request failed:', error);
            throw error;
        }
    }
}