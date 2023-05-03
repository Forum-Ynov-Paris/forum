package main

import (
	Routeur "Forum/Controllers"
	DB "Forum/Controllers/DB"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	db := DB.DBController{}
	db.INIT("db.db")
	Routeur.Routeur(db)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err)
	}
}
