import Page from './pages.js'
import { alert_icons_iframes, alert_infos } from '../utils/alert.js'

export default class Register extends Page {
    constructor() {
        super();
        this.setTitle('Register');
        this.UserInfos = {}
        this.formContainer = null
    }

    bindInputs() {
        this.formContainer = document.getElementById('login-up')

        this.formContainer.addEventListener('input', (e) => {
            if (e.target.classList.contains('login__input')) {
                this.UserInfos[e.target.name] = e.target.value
            }
        })

        this.formContainer.addEventListener('click', (e) => {
            if (e.target.id === 'sign-up') {
                this.register()
            }
        })
    }

    register() {
        this.FormContainer = document.getElementById('login-up')
        this.UserInfos.age = parseInt(this.UserInfos.age)
        fetch("http://localhost:3000/auth/public/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(this.UserInfos)
        })
        .then(response => response.json())
        .then(data => { 
            if (data.message == "Registering Successfuly") {
                if (alert_infos(data.message, alert_icons_iframes.success, this.FormContainer)) {
                    setTimeout(() => {
                        window.location.href = "/login"
                    }, 3000);
                }
            } else {
                alert_infos(data.message, alert_icons_iframes.failed, this.FormContainer)
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
                    <form action="" class="login__create" id="login-up">
                        <h1 class="login__title">Create Account</h1>

                        <div class="login__box">
                            <i class='bx bx-user login__icon'></i>
                            <input name="nickname" type="text" placeholder="Nickname" class="login__input">
                        </div>

                        <div class="login__box">
                            <i class='bx bx-calendar login__icon'></i>
                            <input name="age" type="text" placeholder="Age" class="login__input">
                        </div>

                        <div class="login__box">
                            <i class='bx bx-male login__icon'></i>
                            <input name="gender" type="text" placeholder="Gender" class="login__input">
                        </div>

                        <div class="login__box">
                            <i class='bx bx-id-card login__icon'></i>
                            <input name="firstName" type="text" placeholder="FirstName" class="login__input">
                        </div>

                        <div class="login__box">
                            <i class='bx bx-id-card login__icon'></i>
                            <input name="lastName" type="text" placeholder="LastName" class="login__input">
                        </div>

                        <div class="login__box">
                            <i class='bx bx-at login__icon'></i>
                            <input name="email" type="text" placeholder="Email" class="login__input">
                        </div>

                        <div class="login__box">
                            <i class='bx bx-lock-alt login__icon'></i>
                            <input name="password" type="password" placeholder="Password" class="login__input">
                        </div>

                        <div id="sign-up" class="login__button">Sign Up</div>

                        <div>
                            <span class="login__account">Already have an Account ?</span>
                            <span class="login__signup" href="/login" data-link>Sign In</span>
                        </div>

                    </form>
                </div>
            </div>
        </div>
        `
    }
}