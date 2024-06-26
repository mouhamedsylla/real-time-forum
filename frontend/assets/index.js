import Login from "./js/pages/login.js"
import Register from "./js/pages/register.js"

const app = document.getElementById("app")

const pages = {
    "/register": new Register(),
    "/login": new Login()
}

function getPage(path) {
    var page = ""
    Object.entries(pages).forEach(([key, value]) => {
        if (key == path) {
            page = value.getHTML()
            console.log(page)
        }
    })
    return page
}

function navigateTo(path) {
    history.pushState(null, "", path)
    app.innerHTML = getPage(path)
}

window.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll("a[data-link]").forEach(a => {
        a.addEventListener("click", (event) => {
            event.preventDefault()
            console.log('ok')
            navigateTo(event.target.getAttribute("href"))
        })
    })
})