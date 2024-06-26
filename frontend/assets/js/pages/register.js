import Page from "./pages.js";

export default class Register extends Page {
    constructor() {
        super("Register")
        this.UserInfos = {}
    }

    bindInputs() {
        const inputs = document.querySelectorAll(".login__input")
        inputs.forEach(input => {
            input.addEventListener("input", (e) => {
                this.UserInfos[e.target.name] = input.value
            })
        })

        const signUp = document.getElementById("sign-up")
        signUp.addEventListener("click", (e) => {
            console.log("User Infos: ", this.UserInfos)
        })
    }

    getHTML() {
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
                            <input name="age" type="text" placeholder="Age" class="login__input">
                        </div>

                        <div class="login__box">
                            <input name="gender" type="text" placeholder="Gender" class="login__input">
                        </div>

                        <div class="login__box">
                            <input name="firstName" type="text" placeholder="FirstName" class="login__input">
                        </div>

                        <div class="login__box">
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

                        <!-- <div class="login__social">
                            <a href="#" class="login__social-icon"><i class='bx bxl-facebook' ></i></a>
                            <a href="#" class="login__social-icon"><i class='bx bxl-twitter' ></i></a>
                            <a href="#" class="login__social-icon"><i class='bx bxl-google' ></i></a>
                        </div> -->
                    </form>
                </div>
            </div>
        </div>
        `
    }
}