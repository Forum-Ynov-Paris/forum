package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func HomeClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		session, err := store.Get(r, "forum")
		data := struct {
			Name  string
			Posts []API.Article
		}{}
		if session.Values["authenticated"] == true {
			data.Name = session.Values["username"].(string)
			data.Posts = API.GetArticles()
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
