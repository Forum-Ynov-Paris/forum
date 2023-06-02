package Forum

import (
	DB "Forum/Controllers/DB"
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
			db.POST("INSERT INTO user(firstname, lastname, email, pseudo, password, profil) VALUES (?, ?, ?, ?, ?, ?)", firstName, lastName, email, pseudo, hash, "")
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})
}
