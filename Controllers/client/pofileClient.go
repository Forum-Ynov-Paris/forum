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

func ProfileClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "forum")

		articles := API.SearchArticles(session.Values["username"].(string), db)
		Post := make([]post, len(articles))
		for i, article := range articles {
			Post[i].Article = article
			Post[i].Username = DB.GetUsername(db, article.Uuid)
		}

		profile := struct {
			Name  string
			Email string
			Fname string
			Lname string
			Img   string
			Posts []post
		}{}

		row, err := db.QUERY("SELECT Email, Lastname, Firstname FROM user WHERE pseudo = ?", session.Values["username"].(string))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			var Email string
			for row.Next() {
				err = row.Scan(&Email, &profile.Lname, &profile.Fname)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			profile.Img = DB.Profil(db, store, r, w)
			profile.Name = session.Values["username"].(string)
			profile.Email = Email
			profile.Posts = Post
		}
		t, err := template.ParseFiles("./static/profile.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = t.Execute(w, profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	http.HandleFunc("/profile/edit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			img := r.FormValue("profileimg")
			fmt.Println(img)
			session, _ := store.Get(r, "forum")
			//_, err := db.Database.Query("UPDATE user SET profil = ? WHERE pseudo = ?", img, session.Values["username"].(string))
			//if err != nil {
			//	http.Error(w, err.Error(), http.StatusInternalServerError)
			//}
			stmt, err := db.Database.Prepare("UPDATE user SET profil = ? WHERE pseudo = ?")
			if err != nil {
				log.Fatal(err)
			}

			// Met à jour le nom de l'utilisateur avec l'ID 1
			_, err = stmt.Exec(img, session.Values["username"].(string))
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		}
	})
}
