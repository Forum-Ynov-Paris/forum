package Forum

import (
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func HomeClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "forum")
		data := struct {
			Name string
		}{}
		if session.Values["authenticated"] == true {
			data.Name = session.Values["username"].(string)
		} else {
			data.Name = "Guest"
		}
		t, err := template.ParseFiles("./static/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
