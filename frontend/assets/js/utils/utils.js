const alert_icons_iframes = {
    success: `<iframe src="https://lottie.host/embed/48d63482-8c12-496e-9c21-3bebb982079b/uKmHVBxGMP.json" style="border: none; height: 50px;"></iframe>`,
    failed: `<iframe src="https://lottie.host/embed/647ace79-3cec-4dfe-982a-27ea2e5260f7/UQRJfTKbjp.json" style="border: none; height: 50px;"></iframe>`
}

function alert(message, type, parentElement) {
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
}

export { alert, alert_icons_iframes }