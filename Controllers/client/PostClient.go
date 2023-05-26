package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"fmt"
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
	type commentary struct {
		C        API.Commentaire
		Username string
	}

	type data struct {
		Post        API.Article
		Commentates []commentary
		Name        string
	}
	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {

		Title := r.URL.String()[6:]

		fmt.Println(Title)

		if ID == -1 {
			http.Redirect(w, r, "/", http.StatusNotFound)
		}

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

		Data := data{
			API.GetIndexByTitle(Title), //changer + tard
			make([]commentary, 0),
			session.Values["username"].(string),
		}

		for _, c := range Data.Post.Commentaire {
			Data.Commentates = append(Data.Commentates, commentary{c, db.GetUsername(c.Uuid)})
		}

		fmt.Println(len(Data.Commentates))

		content = r.FormValue("newComment")
		t, err := template.ParseFiles("./static/post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if r.Method == "POST" {
			API.AddComment(ID, content, uid)
		}

	})
}
