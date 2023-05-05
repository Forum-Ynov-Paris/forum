package Forum

import (
	DB "Forum/Controllers/DB"
	Client "Forum/Controllers/client"
	"github.com/gorilla/sessions"
	"net/http"
)

func Routeur(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/login", Client.Login)
	http.HandleFunc("/register", Client.Register)
	Client.Search(db, store)
	Client.LoginPost(db, store)
	Client.HomeClient(db, store)
	Client.CreatePost(store)
	Client.RegisterPost(db)
	Client.HandleCreatepost(db, store)

}
