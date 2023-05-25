package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

var (
	ID      int
	uid     int
	content string
)

func InitPostClient(db DB.DBController, store *sessions.CookieStore) {
	type data struct {
		Post API.Article
		Name string
	}
	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer la valeur du paramètre dans l'URL
		vars := mux.Vars(r)
		Title := vars["Title"]

		//Title := "r.URL.String()[:6]"

		fmt.Println(Title)

		ID = API.GetIndexByTitle(Title)

		session, _ := store.Get(r, "forum")

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
		t, err := template.ParseFiles("/static/post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		Data := data{
			API.GetArticle(ID), //changer + tard
			"Guest",
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
