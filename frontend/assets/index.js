import Router from "./js/router/router.js"
import Api from "./js/api/api.js"
import Login from "./js/pages/login.js"
import Register from "./js/pages/register.js"
import Home from "./js/pages/home.js"
import Error from "./js/pages/error.js"


const api = new Api()
const home = new Home()
const login = new Login()
const error = new Error()
const register = new Register()


const router = new Router({
    "/": login,
    "/home": home,
    "/error": error,
    "/login": login,
    "/register": register,
})

router.init()

export default api