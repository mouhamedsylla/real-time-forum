import Page from "./pages.js";

export default class Login extends Page {

    constructor() {
        super("Login")
        this.credentials = {}
    }

    bindInputs() {
        const inputs = document.querySelectorAll(".login__input")
        inputs.forEach(input => {
            input.addEventListener("input", (e) => {
                this.credentials[e.target.name] = input.value
            })
        })

        const signIn = document.getElementById("sign-in")
        signIn.addEventListener("click", (e) => {
            console.log("Credentials: ", this.credentials)
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