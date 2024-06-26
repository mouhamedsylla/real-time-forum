import Login from "./js/pages/login.js"
import Register from "./js/pages/register.js"

const app = document.getElementById("app")

const register = new Register()
const login = new Login()

const pages = {
    "/register": register,
    "/login": login
}

function getPage(path) {
    var page = ""
    Object.entries(pages).forEach(([key, value]) => {
        if (key == path) {
            page = value.getHTML()
        }
    })
    return page
}

function navigateTo(path) {
    history.pushState(null, "", path)
    app.innerHTML = getPage(path)
    

    if (path == "/login") {
        login.bindInputs()
    }

    if (path == "/register") {
        register.bindInputs()
    }
}

document.addEventListener("DOMContentLoaded", (e) => {
    //const matchLinks = document.querySelectorAll("[data-link]") 
    document.body.addEventListener("click", (event) => {
        event.preventDefault()
        if (event.target.matches("[data-link]")) {
            let route = event.target.getAttribute("href")
            if (route) {
                navigateTo(route)
            }
        }
    })
})

console.log("hello")