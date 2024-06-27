export default class API {
    constructor() {
        this.baseURL
        this.client
    }

    setbaseURL(baseURL) {
        this.baseURL = baseURL
    }

    setClient(client) {
        this.client = Object.assign({}, client) 
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