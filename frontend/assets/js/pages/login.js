import Page from './pages.js'
import {alert_infos, alert_icons_iframes} from '../utils/alert.js'

export default class Login extends Page {
    constructor() {
        super();
        this.setTitle('Login');
        this.credentials = {}
        this.formContainer = null
    }

    bindInputs() {
        this.formContainer = document.getElementById('login-in')

        this.formContainer.addEventListener('input', (e) => {
            if (e.target.classList.contains('login__input')) {
                this.credentials[e.target.name] = e.target.value
            }
        })

        this.formContainer.addEventListener('click', (e) => {
            if (e.target.id === 'sign-in') {
                this.login()   
            }
        })
    }

    login() {
        fetch("http://localhost:3000/auth/public/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(this.credentials)
        })
        .then(response => response.json())
        .then(data => {
            return data.message == "login successfull" ?
                alert_infos(data.message, alert_icons_iframes.success, this.formContainer) :
                alert_infos(data.message, alert_icons_iframes.failed, this.formContainer)
        })
        .then(result => {
            if (result) {
                setTimeout(() => {
                    window.location.href = "/home"
                }, 3000)
            }
        })
    }

    async getHTML() {
        return `
        <div class="login">
            <div class="login__content">
                <div class="login__img">
                    <img src="/frontend/assets/css/img/login__img.svg" alt="">
                </div>

                <div class="login__forms">
                    <form action="" class="login__registre" id="login-in">
                        <h1 class="login__title">Sign In</h1>
    
                        <div class="login__box">
                            <i class='bx bx-user login__icon'></i>
                            <input type="text" name="identifier" placeholder="Nickname or Email" class="login__input">
                        </div>
    
                        <div class="login__box">
                            <i class='bx bx-lock-alt login__icon'></i>
                            <input type="password" name="password" placeholder="Password" class="login__input">
                        </div>

                        <a href="#" class="login__forgot">Forgot password?</a>

                        <div id="sign-in" class="login__button" >Sign In</div>

                        <div>
                            <span class="login__account">Don't have an Account ?</span>
                            <span class="login__signin" id="sign-up" href="/register" data-link>Sign Up</span>
                        </div>
                    </form>
                </div>
            </div>
        </div>
        `
    }

}