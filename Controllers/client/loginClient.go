package Forum

import (
	DB "Forum/Controllers/DB"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("static/connexion.html"))
	err := template.Execute(w, nil)
	if err != nil {
		print(err)
	}
}

func LoginPost(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		session, err1 := store.Get(r, "forum")
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		if r.Method == "POST" {
			pseudo := r.FormValue("pseudo")
			spassword := r.FormValue("password")
			fmt.Println("login yes")
			rows, _ := db.QUERY("SELECT password FROM user WHERE pseudo = ?", pseudo)
			for rows.Next() {
				var password string
				rows.Scan(&password)
				fmt.Println("login yes")
				err := DB.ComparePasswords(password, spassword)
				if err != nil {
					fmt.Println("Login failed")
					print(err)
				} else {
					fmt.Println("login yes")
					session.Values["authenticated"] = true
					session.Values["username"] = pseudo // remplacer avec le nom d'utilisateur réel
					err2 := session.Save(r, w)
					if err2 != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
			defer rows.Close()
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		session, err1 := store.Get(r, "forum")
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["authenticated"] = false
		session.Values["username"] = ""
		err2 := session.Save(r, w)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
