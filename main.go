package main

import (
	Routeur "Forum/Controllers"
	DB "Forum/Controllers/DB"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("lala"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	db := DB.DBController{}
	db.INIT("db.db")
	Routeur.Routeur(db, store)
	fmt.Println("Server is running on port 8080 : " + "http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		print(err)
	}
}
