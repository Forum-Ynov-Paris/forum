package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
)

type data struct {
	Post API.Article
	Name string
}

var (
	Data    data
	id      int
	uid     int
	content string
)

func InitPostClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "forum")

		Data = data{
			API.GetArticle(id), //changer + tard
			"Guest",
		}
		id = 0
		row, err := db.QUERY("SELECT id FROM user WHERE pseudo = ?", session.Values["username"].(string))
		if err != nil {
			log.Fatal(err)
		} else {
			for row.Next() {
				err = row.Scan(&id)
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
			API.AddComment(id, content, uid)
		}

	})
}
