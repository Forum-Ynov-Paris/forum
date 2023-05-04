function closes() {
    var box = document.getElementById("box")
    box.style.display = "none"
    var h1 = document.getElementById("erreur")

    if (h1 == null) {
        box.style.display = "none"
    }
}

function closesPopup() {
    var box2 = document.getElementById("box_")
    box2.style.display = "none"
    var h1 = document.getElementById("erreur")

    if (h1 == null) {
        box2.style.display = "none"
    }
}

function openPopup() {
    var box = document.querySelector("#box_")
    box.style.display = "initial"
}

function openAndClose() {
    var popup = document.getElementById("wrapper")

    if (popup.style.display == "initial") {
        popup.style.display = "none"
    } else {
        popup.style.display = "initial"
    }
    window.onclick = function(event) {
        if (event.target == body || event.target == topic) {
            popup.setAttribute("style", "display:none;")
        }
        if (event.target == body || event.target == topic) {
            popup2.setAttribute("style", "display:none;")
        }
    }
}

function openAndClose_(id) {
    var popup = document.getElementById("wrapper_"+id)

    if (popup.style.display == "initial") {
        popup.style.display = "none"
    } else {
        popup.style.display = "initial"
    }
    var popup = document.getElementById("wrapper")
    var popup2 = document.getElementById("wrapper_"+id)
    var body = document.getElementById("body")
    var topic = document.getElementById("topic")
    
    window.onclick = function(event) {
        if (event.target == body || event.target == topic) {
            popup.setAttribute("style", "display:none;")
        }
        if (event.target == body || event.target == topic) {
            popup2.setAttribute("style", "display:none;")
        }
    }
}