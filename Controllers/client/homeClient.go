package Forum

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("static/templates/home.html"))
	err := template.Execute(w, nil)
	if err != nil {
		print(err)
	}
}
