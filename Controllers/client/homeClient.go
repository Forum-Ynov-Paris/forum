package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

type post struct {
	API.Article
	Username string
}

func HomeClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		articles := API.GetArticles()
		Posts := make([]post, len(articles))
		for i, article := range articles {
			Posts[i].Article = article
			Posts[i].Username = DB.GetUsername(db, article.Uuid)
		}

		session, err := store.Get(r, "forum")
		data := struct {
			Name      string
			Connected bool
			Vots      int
			Img       string
			Posts     []post
		}{}
		if session.Values["authenticated"] == true {
			data.Name = session.Values["username"].(string)
			data.Connected = true
			//data.Img = DB.Profil(db, store, r, w)
			data.Posts = Posts
		} else {
			data.Name = "Guest"
			data.Connected = false
			data.Posts = Posts
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
