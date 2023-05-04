package main

import (
	Routeur "Forum/Controllers"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("lala"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	db := DB.DBController{}
	db.INIT("db.db")
	Routeur.Routeur(db, store)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err)
	}
}
