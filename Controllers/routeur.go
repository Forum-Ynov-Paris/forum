package Forum

import (
	DB "Forum/Controllers/DB"
	Client "Forum/Controllers/client"
	"net/http"
)

func Routeur(db DB.DBController) {
	http.HandleFunc("/", Client.Home)
	http.HandleFunc("/login", Client.Login)
	http.HandleFunc("/register", Client.Register)
	Client.LoginPost(db)
}
