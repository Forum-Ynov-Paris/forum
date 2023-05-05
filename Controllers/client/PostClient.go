package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	type data struct {
		post API.Article
	}

	Data := API.GetArticle(0) //changer + tard

	t, err := template.ParseFiles("./static/post.html")
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

func PostClient(db DB.DBController, store *sessions.CookieStore) {

}
