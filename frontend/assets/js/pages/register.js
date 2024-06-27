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
                this.UserInfos[e.target.name] = e.target.name == "age" ? parseInt(input.value) : input.value 
            })
        })

        const signUp = document.getElementById("sign-up")
        signUp.addEventListener("click", () => {
            // console.log("User Infos: ", this.UserInfos)
            this.register()
        })
    }

    register() {
        fetch("http://localhost:3000/auth/register", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(this.UserInfos)
        })
        .then(response => response.json())
        .then(data => { 
            console.log(data)
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