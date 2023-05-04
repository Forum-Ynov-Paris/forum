package Forum

import (
	DB "Forum/Controllers/DB"
	"fmt"
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("static/inscription.html"))
	err := template.Execute(w, nil)
	if err != nil {
		print(err)
	}
}

func RegisterPost(db DB.DBController) {
	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			firstName := r.FormValue("firstname")
			lastName := r.FormValue("lastname")
			email := r.FormValue("email")
			pseudo := r.FormValue("pseudo")
			password := r.FormValue("password")
			hash, _ := DB.HashPassword(password)
			fmt.Println(firstName, lastName, email, pseudo, password, hash)
			db.POST("INSERT INTO user(firstname, lastname, email, pseudo, password) VALUES (?, ?, ?, ?, ?)", firstName, lastName, email, pseudo, hash)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
