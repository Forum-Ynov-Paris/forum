package Forum

import (
    "html/template"
    "net/http"
)

func createPost(w http.ResponseWriter, r *http.Request) {
    // Récupérer le nom d'utilisateur de la session
    session, _ := store.Get(r, "session-name")
    username := session.Values["username"].(string)

    // Afficher le formulaire avec le nom d'utilisateur prérempli
    tmpl, _ := template.ParseFiles("/static/createpost.html")
    tmpl.Execute(w, map[string]interface{}{
        "Username": username,
    })
}

func main() {
    http.HandleFunc("/static/createpost.html", createpost)
    http.ListenAndServe(":8080", nil)
}
