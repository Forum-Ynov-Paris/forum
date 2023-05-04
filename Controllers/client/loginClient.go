package Forum

import (
	DB "Forum/Controllers/DB"
	"fmt"
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

func LoginPost(db DB.DBController) {
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			pseudo := r.FormValue("pseudo")
			spassword := r.FormValue("password")
			rows, _ := db.QUERY("SELECT password FROM user WHERE pseudo = ?", pseudo)
			for rows.Next() {
				var password string
				rows.Scan(&password)
				err := DB.ComparePasswords(password, spassword)
				if err != nil {
					fmt.Println("Login failed")
					print(err)
				} else {
					fmt.Println("Login success")
					//TODO: Create session
				}
			}

		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
