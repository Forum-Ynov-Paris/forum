package Forum

import (
    "html/template"
    "net/http"
	API "Forum/Controllers/API"
)

func createPost(w http.ResponseWriter, r *http.Request) {
    // Récupérer le nom d'utilisateur de la session
    session, _ := store.Get(r, "session-name")
    username := session.Values["username"].(string)

    t, err := template.ParseFiles("/static/createpost.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    err = t.Execute(w, Data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func handleCreatepost(db DB.DBController, store *sessions.CookieStore) {
	article := new(API.Article)
}
