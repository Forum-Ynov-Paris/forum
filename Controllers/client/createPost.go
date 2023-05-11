package Forum

import (
	Forum "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	"time"
)

func CreatePost(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/newpost", func(w http.ResponseWriter, r *http.Request) {
		// Récupérer le nom d'utilisateur de la session
		//session, _ := store.Get(r, "session-name")
		//username := session.Values["username"].(string)

		session, err := store.Get(r, "forum")
		data := struct {
			Name  string
			Posts []post
		}{}
		if session.Values["authenticated"] == true {
			data.Name = session.Values["username"].(string)
		} else {
			data.Name = "Guest"
		}
		t, err := template.ParseFiles("./static/newpost.html")
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

	http.HandleFunc("/api/newpost", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "forum")
		username := session.Values["username"].(string)

		if r.Method == "POST" {
			title := r.FormValue("title")
			content := r.FormValue("content")
			tag := r.FormValue("tag")
			row, err := db.QUERY("SELECT id FROM user WHERE pseudo = ?", username)
			if err != nil {
				log.Fatal(err)
			} else {
				for row.Next() {
					var id int
					err = row.Scan(&id)
					if err != nil {
						log.Fatal(err)
					}
					currentTime := time.Now()
					fmt.Println(content)
					Forum.AddPost(title, tag, content, currentTime.String(), id)
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
