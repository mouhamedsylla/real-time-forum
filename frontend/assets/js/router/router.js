import { backToHome, modeSelect } from "../utils/other.js"

export default class Router {
    constructor(routes) {
        this.routes = routes
        this.app = document.getElementById('app')
    }

    init() {
        window.addEventListener("popstate", () => { this.renderView(window.location.pathname) })
        document.addEventListener("DOMContentLoaded", () => {
            backToHome(window.location.pathname)
            this.navigateTo(window.location.pathname)
            if (window.location.pathname !== "/home") {
                document.addEventListener("click", (event) => { this.handleLinkClick(event) })
            }
        })
    }

    handleLinkClick(event) {
        if (event.target.matches("[data-link]")) {
            event.preventDefault()
            const route = event.target.getAttribute("href")
            if (route) {
                backToHome(route)
                this.navigateTo(route)
            }
        }
    }

    async renderView(path) {
        const page = this.routes[path] || this.routes["/error"]
        this.app.innerHTML = await page.getHTML();
        if (path === "/home") {
            await page.renderComponents()
            modeSelect()
        }

        if (typeof page.bindInputs === "function") {
            page.bindInputs()
        }
    }

    navigateTo(path) {
        history.pushState(null, "", path)
        this.renderView(path)
    }
}