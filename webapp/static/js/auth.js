var spans = document.getElementsByClassName("close");

var loginModal = document.getElementById("loginModal");
var loginBtn = document.getElementById("loginButton");
var loginErrors = document.getElementById("loginErrors");
var loginClose = spans[0];

var registerModal = document.getElementById("registerModal");
var registerBtn = document.getElementById("registerButton");
var registerErrors = document.getElementById("registerErrors");
var registerClose = spans[1];

loginBtn.onclick = function() {
    loginModal.style.display = "block";
}

registerBtn.onclick = function() {
    registerModal.style.display = "block";
}

loginClose.onclick = function() {
    loginModal.style.display = "none";
}

registerClose.onclick = function() {
    registerModal.style.display = "none";
}

window.onclick = function(event) {
    if (event.target == loginModal) {
        loginModal.style.display = "none";
    } else if (event.target == registerModal) {
        registerModal.style.display = "none";
    }
}

document.getElementById("loginForm").onsubmit = function(event) {
    event.preventDefault();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    
    fetch("/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: username,
            password: password
        })
    }).then(async res => {
        if (res.status === 200) {
            location.reload();
        } else {
            loginErrors.textContent = (await res.json()).message;
        }
    }).catch(err => {
        console.error(err);
    });
}

document.getElementById("registerForm").onsubmit = function(event) {
    event.preventDefault();
    var username = document.getElementById("reg-username").value;
    var email = document.getElementById("reg-email").value;
    var password = document.getElementById("reg-password").value;
    
    fetch("/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            username: username,
            email: email,
            password: password
        })
    }).then(async res => {
        if (res.status === 201) {
            location.reload();
        } else {
            registerErrors.textContent = (await res.json()).message;
        }
    }).catch(err => {
        console.error(err);
    });
}