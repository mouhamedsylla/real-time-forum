import Login from "./js/pages/login.js"
import Register from "./js/pages/register.js"
import Error from "./js/pages/error.js"
import Home from "./js/pages/home.js"
import API from './js/api/api.js';

const api = new API();
api.setbaseURL("http://localhost:3000");

export default api;

var found = false

const app = document.getElementById("app")
const register = new Register()
const login = new Login()
const error = new Error()
const home = new Home()


const pages = {
    "/": login, // Default page is login
    "/home": home,
    "/register": register,
    "/login": login,
    "/error": error
}

function renderView(path) {
    Object.entries(pages).forEach(([key, value]) => {
        if (key == path) {
            app.innerHTML = value.getHTML()
            if (typeof value.bindInputs === "function") {
                found = true
                value.bindInputs()
            }
        }
    })
    if (!found && (path == "/login" || path == "/register")) {
        app.innerHTML = error.getHTML()
    }
}

function navigateTo(path) {
    history.pushState(null, "", path)
    renderView(path)
}

document.addEventListener("DOMContentLoaded", (e) => {
    navigateTo(window.location.pathname)
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
