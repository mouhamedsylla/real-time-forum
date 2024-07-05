export {
    alert_infos,
    alert_icons_iframes,
    alert_token_expire,
    alert_loading,
    prependChild,
    alert_typing
}

const alert_icons_iframes = {
    success: `<iframe src="https://lottie.host/embed/48d63482-8c12-496e-9c21-3bebb982079b/uKmHVBxGMP.json" style="border: none; height: 50px;"></iframe>`,
    failed: `<iframe src="https://lottie.host/embed/647ace79-3cec-4dfe-982a-27ea2e5260f7/UQRJfTKbjp.json" style="border: none; height: 50px;"></iframe>`
}

function alert_infos(message, type, parentElement) {
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

async function alert_loading(target, bool) {
    const div = document.createElement("div")
    div.classList.add("loading")
    div.innerHTML = `
        <iframe src="https://lottie.host/embed/454fc9f7-9ba7-4eb6-b116-f52cb7ed59a4/soP2gQypD7.json" style="border: none; height: 60px;"></iframe>
    `
    if (bool && !target.querySelector(".loading")) {
        prependChild(target, div)
        await new Promise(resolve => setTimeout(resolve, 2000))
        target.removeChild(div)
    }
}

function alert_typing() {
    const container = document.querySelector(".chat-messages")
    const div = document.createElement("div")
    div.classList.add("message", "typing")
    div.innerHTML = `
                    <div class="typing typing-1"></div>
                    <div class="typing typing-2"></div>
                    <div class="typing typing-3"></div>
    `
    if (container.querySelector(".typing")) {
        container.removeChild(container.querySelector(".typing"))
    }
    container.appendChild(div)
    container.scrollTop = container.scrollHeight
}

function prependChild(parent, newChild) {
    if (parent.firstChild) {
      parent.insertBefore(newChild, parent.firstChild);
    } else {
      parent.appendChild(newChild);
    }
}