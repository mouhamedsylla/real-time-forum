const like = "rgb(255, 87, 51)"
const dislike = "rgb(52, 152, 219)"
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

function session_expired() {
    const token_payload = parseJwt(document.cookie)
    if (!token_payload) {
        return true
    }
    return false
}


function throttle(fn, delay, { leading = false, trailing = true } = {}) {
    let last = 0
    let timer = null
    return function () {
        const now = +new Date()
        if (!last && leading === false) {
            last = now
        }
        if (now - last > delay) {
            if (timer) {
                clearTimeout(timer)
                timer = null
            }
            fn.apply(this, arguments)
            last = now
        } else if (!timer && trailing !== false) {
            timer = setTimeout(() => {
                fn.apply(this, arguments)
                last = +new Date()
                timer = null
            }, delay)
        }
    }
}


function modeSelect() {
    const toggleIcon = document.getElementById('toggleIcon');
    let isLightMode = true;

        toggleIcon.addEventListener('click', () => {
            if (isLightMode) {
                document.body.style.backgroundColor = 'hsl(252, 30%, 17%)';
                toggleIcon.classList.remove('uil-sun');
                toggleIcon.classList.add('uil-moon');
            } else {
                document.body.style.backgroundColor = 'hsl(252, 30%, 95%)';
                document.body.style.color = 'black';
                toggleIcon.classList.remove('uil-moon');
                toggleIcon.classList.add('uil-sun');
            }
            isLightMode = !isLightMode;
        });
}

function formatTimeAgo(created_at = getCurrentDateTime()) {
    const now = new Date();
    const createdDate = new Date(created_at);

    const diffInSeconds = Math.floor((now - createdDate) / 1000);

    const secondsInMinute = 60;
    const secondsInHour = 60 * secondsInMinute;
    const secondsInDay = 24 * secondsInHour;

    if (diffInSeconds < secondsInMinute) {
        return `DAKAR, ${diffInSeconds} SECONDS AGO`;
    } else if (diffInSeconds < secondsInHour) {
        const minutes = Math.floor(diffInSeconds / secondsInMinute);
        return `DAKAR, ${minutes} MINUTES AGO`;
    } else if (diffInSeconds < secondsInDay) {
        const hours = Math.floor(diffInSeconds / secondsInHour);
        return `DAKAR, ${hours} HOURS AGO`;
    } else {
        const days = Math.floor(diffInSeconds / secondsInDay);
        return `DAKAR, ${days} DAYS AGO`;
    }
}

function getCurrentDateTime() {
    const now = new Date();

    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0'); // Les mois commencent à 0
    const day = String(now.getDate()).padStart(2, '0');
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

let x;
function showToast(){
    let toast = document.getElementById("toast");
    clearTimeout(x);
    toast.style.transform = "translateX(0)";
    x = setTimeout(()=>{
        toast.style.transform = "translateX(400px)"
    }, 4000);
}

function prependChild(parent, newChild) {
    if (parent.firstChild) {
      parent.insertBefore(newChild, parent.firstChild);
    } else {
      parent.appendChild(newChild);
    }
  }
  

export {
        alert, alert_icons_iframes, 
        alert_token_expire,
        backToHome, 
        modeSelect, 
        session_expired, 
        parseJwt,
        throttle,
        alert_loading,
        formatTimeAgo,
        showToast,
        like,
        dislike,
        prependChild
}