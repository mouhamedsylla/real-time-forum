const alert_icons_iframes = {
    success: `<iframe src="https://lottie.host/embed/48d63482-8c12-496e-9c21-3bebb982079b/uKmHVBxGMP.json" style="border: none; height: 50px;"></iframe>`,
    failed: `<iframe src="https://lottie.host/embed/647ace79-3cec-4dfe-982a-27ea2e5260f7/UQRJfTKbjp.json" style="border: none; height: 50px;"></iframe>`
}

function alert(message, type, parentElement) {
    let result = alert_icons_iframes.success == type
    if (parentElement) {
        for (let i = 0; i < parentElement.children.length; i++) {
            if (parentElement.children[i].className.includes("alert")) {
                parentElement.children[i].remove()
            }
        }
    }
    

    const div = document.createElement("div")
    div.className = "alert"
    div.innerHTML = `
        ${type}
        <p>${message}</p>
        
    `
    parentElement.appendChild(div)
    return result
}

function alert_token_expire() {
    const body = document.querySelector("body")
    const div = document.createElement("div")
    div.classList.add("token-alert")
    div.innerHTML = `
        <div class="card">
            <h2>token not found or expired</h2>
            <br>
            <label class="btn btn-primary" id="token-lost-alert">Back to login</label>
        </div>`
    body.appendChild(div)
    const btn = document.getElementById("token-lost-alert")
    btn.addEventListener("click", () => {
        window.location.href = "/"
    })
}

function parseJwt(token) {
    try {
        return JSON.parse(atob(token.split('.')[1]));
    } catch (e) {
        return null;
    }
}

function backToHome(route) {
    if (route == "/login" || route == "/") {
        const token_payload = parseJwt(document.cookie)
        if (token_payload) {
            window.location.href = "/home"
        }
    }
}

function modeSelect() {
    const toggleIcon = document.getElementById('toggleIcon');
    let isLightMode = true;

        toggleIcon.addEventListener('click', () => {
            console.log('clicked');
            if (isLightMode) {
                document.body.style.backgroundColor = 'hsl(252, 30%, 17%)';
                toggleIcon.classList.remove('uil-sun');
                toggleIcon.classList.add('uil-moon');
            } else {
                document.body.style.backgroundColor = 'white';
                document.body.style.color = 'black';
                toggleIcon.classList.remove('uil-moon');
                toggleIcon.classList.add('uil-sun');
            }
            isLightMode = !isLightMode;
        });
}

export { alert, alert_icons_iframes, alert_token_expire, parseJwt, backToHome, modeSelect }