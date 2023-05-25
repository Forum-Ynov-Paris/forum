package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

type dataa struct {
	Post API.Article
	Name string
}

var (
	Data    dataa
	ID      int
	uid     int
	content string
)

func InitPostClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer la valeur du paramètre dans l'URL
		vars := mux.Vars(r)
		Title := vars["Title"]

		session, _ := store.Get(r, "forum")

		Data = dataa{
			API.GetArticle(ID), //changer + tard
			"Guest",
		}
		ID = API.GetIndexByTitle(Title)
		row, err := db.QUERY("SELECT id FROM user WHERE pseudo = ?", session.Values["username"].(string))
		if err != nil {
			log.Fatal(err)
		} else {
			for row.Next() {
				err = row.Scan(&ID)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		content = r.FormValue("newComment")
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

		if r.Method == "POST" {
			API.AddComment(ID, content, uid)
		}

	})
}
