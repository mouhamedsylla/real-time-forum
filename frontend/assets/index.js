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
    document.body.addEventListener("click", (event) => {
        console.log("ok")
        event.preventDefault()
        console.log(event.target.getAttribute("href"))
        if (event.target.matches("[data-link]")) {
            navigateTo(event.target.getAttribute("href"))
        }
    })
})

console.log("hello")