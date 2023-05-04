function showPassword() {
    var input = document.getElementById("myInput");
    var eye = document.getElementById("eye");
    var slash = document.getElementById("slash");

    if (input.type === 'password') {
        input.type = 'text';
        eye.style.display = "initial";
        slash.style.display = "none";
    } else {
        input.type = "password";
        eye.style.display = "none";
        slash.style.display = "initial"
    }
}

function showPasswordConfirm() {
    var input = document.getElementById("myInput2");
    var eye = document.getElementById("eye2");
    var slash = document.getElementById("slash2");

    if (input.type === 'password') {
        input.type = 'text';
        eye.style.display = "initial";
        slash.style.display = "none";
    } else {
        input.type = "password";
        eye.style.display = "none";
        slash.style.display = "initial"
    }
}