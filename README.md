# Real-Time Forum Project

## Objectives

This project focuses on the following key points:

1. User registration and login
2. Post creation
3. Commenting on posts
4. Private messaging

Your new forum will include five distinct components:

- **SQLite**: Data storage, similar to the previous forum
- **Golang**: Data handling and Websockets (Backend)
- **JavaScript**: Handling all Frontend events and client Websockets
- **HTML**: Structuring the page elements
- **CSS**: Styling the page elements

You will have a single HTML file, so all page changes need to be managed in JavaScript (Single Page Application).

## Project Architecture

This project employs a microservices architecture. Each service is responsible for a specific functionality of the application. Here is the overall project structure:

```
.
├── backend
│ ├── cmd
│ │ ├── app
│ │ │ └── main.go
│ │ └── bootservices
│ │ └── main.go
│ ├── orm
│ ├── server
│ │ ├── gateway
│ │ ├── microservices
│ │ ├── middleware
│ │ └── router
│ ├── services
│ │ ├── auth
│ │ │ ├── controllers
│ │ │ ├── database
│ │ │ │ └── migrates
│ │ │ └── models
│ │ ├── chat
│ │ │ ├── controllers
│ │ │ ├── database
│ │ │ │ └── migrates
│ │ │ └── models
│ │ ├── notification
│ │ │ ├── controllers
│ │ │ ├── database
│ │ │ │ └── migrates
│ │ │ └── models
│ │ └── posts
│ │ ├── controllers
│ │ ├── database
│ │ │ └── migrates
│ │ └── models
│ └── utils
│ ├── jwt
│ ├── key
│ └── validation
│ └── test
├── frontend
│ └── assets
│ ├── css
│ │ └── img
│ ├── images
│ └── js
│ ├── api
│ ├── components
│ ├── pages
│ ├── router
│ └── utils
```


### Architecture Details

1. **Microservices**
    - Each microservice is autonomous and responsible for a specific part of the application (authentication, chat, notifications, posts).
    - Microservices are defined in the `backend/services` directory.
    - Interaction between microservices is done via REST APIs.

2. **Gateway**
    - A single entry point for all requests is managed by the gateway located in `backend/server/gateway`.
    - The gateway routes requests to the appropriate microservices based on the URL.

3. **Websockets**
    - Used to handle real-time messaging in the chat and notifications.
    - Implemented in the backend with Golang and the frontend with JavaScript.

4. **JWT Authentication**
    - Using JSON Web Tokens (JWT) to secure endpoints.
    - Custom JWT implementation in `backend/utils/jwt`.

### Services

- **Auth**: Manages user registration, login, and logout.
- **Chat**: Manages private messages between users.
- **Notification**: Manages real-time notifications.
- **Posts**: Manages post creation and comments.

## Usage

To start the application, follow these steps:

1. **Clone the repository**

```bash
git clone <repo-url>
cd <repo-name>
```

2. **Start the backend services**

```
cd backend/cmd/app
go run main.go
```

Then, start the services in **cmd/bootservices** :

```
cd ../bootservices
go run main.go
```

Open your browser and go to http://localhost:3000

## Conclusion
This project uses a microservices architecture to separate responsibilities and facilitate maintenance and scalability. Each service is autonomous and communicates with others via REST APIs. Websockets are used for real-time communication for chat and notification functionalities. Authentication is secured using JWT.