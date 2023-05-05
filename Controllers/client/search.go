package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

type postSearch struct {
	API.Article
	Username string
}

func SearchClient(template *template.Template, data interface{}) {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		err := template.Execute(w, data)
		if err != nil {
			print(err)
		}
	})
}

func Search(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			search := r.FormValue("search")
			articles := API.SearchArticles(search, db)
			Posts := make([]postSearch, len(articles))
			for i, article := range articles {
				Posts[i].Article = article
				Posts[i].Username = db.GetUsername(article.Uuid)
			}

			session, err := store.Get(r, "forum")
			data := struct {
				Name  string
				Posts []postSearch
			}{}
			if session.Values["authenticated"] == true {
				data.Name = session.Values["username"].(string)
				data.Posts = Posts
			} else {
				data.Name = "Guest"
				data.Posts = Posts
			}
			t, err := template.ParseFiles("./static/search.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = t.Execute(w, data)
			if err != nil {
				print(err)
			}
		}

	})
}
