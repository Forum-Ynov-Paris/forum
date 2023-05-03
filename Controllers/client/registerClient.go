package Forum

import (
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("static/Register.html"))
	err := template.Execute(w, nil)
	if err != nil {
		print(err)
	}
}
