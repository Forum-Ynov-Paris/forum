function createDiv(idPost) {
    var del = document.getElementById('visible_'+idPost)
    if (del.getAttribute("style") == "display:none"){
        del.setAttribute("style", "display:null");
    } else {
        del.setAttribute("style", "display:none");
    }
}