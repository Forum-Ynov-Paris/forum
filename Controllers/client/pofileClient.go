package Forum

import (
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func ProfileClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "forum")

		profile := struct {
			Name  string
			Email string
		}{}

		row, err := db.QUERY("SELECT Email FROM user WHERE pseudo = ?", session.Values["username"].(string))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		} else {
			var Email string
			for row.Next() {
				err = row.Scan(&Email)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			profile.Name = session.Values["username"].(string)
			profile.Email = Email
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
}
