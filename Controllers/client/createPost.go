package Forum

import (
	Forum "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
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
			Name string
			//Img  string
		}{}
		if session.Values["authenticated"] == true {
			data.Name = session.Values["username"].(string)
			//row, _ := db.QUERY("SELECT profil FROM user WHERE pseudo = ?", session.Values["username"].(string))
			//for row.Next() {
			//	var img string
			//	err = row.Scan(&img)
			//	if err != nil {
			//		http.Error(w, err.Error(), http.StatusInternalServerError)
			//		return
			//	}
			//	if img != "" {
			//		data.Img = img
			//	} else {
			//		data.Img = ""
			//	}
			//}
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
					var date string
					for _, i := range currentTime.String() {
						if i != '.' {
							date += string(i)
						} else {
							break
						}
					}
					Forum.AddPost(title, tag, content, date, id)
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
