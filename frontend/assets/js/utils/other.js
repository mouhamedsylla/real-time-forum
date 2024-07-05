export {
    backToHome,
    parseJwt,
    session_expired,
    throttle,
    showToast,
    formatTimeAgo,
    like,
    dislike,
    formatDate,
    modeSelect
}

const like = "rgb(255, 87, 51)"
const dislike = "rgb(52, 152, 219)"

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

function showChat() {
    const chatIcon = document.getElementById("chat-icon")

    if (window.innerWidth < 600) {
        chatIcon.style.display = "inline-block"
    } else {
        
    }
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

function formatDate(dateString) {
    let date;
    
    if (dateString) {
        // Convertir la chaîne de date en un objet Date
        date = new Date(dateString);
    } else {
        // Utiliser l'heure actuelle si aucune chaîne de date n'est fournie
        date = new Date();
    }

    // Récupérer les composants de la date
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Les mois sont indexés à partir de 0
    const year = date.getFullYear();

    // Formater la date selon le format désiré
    const formattedDate = `${hours}:${minutes} ${day}-${month}-${year}`;

    return formattedDate;
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