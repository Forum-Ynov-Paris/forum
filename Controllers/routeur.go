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
	http.HandleFunc("/post", Client.Post)
	Client.LoginPost(db, store)
	Client.HomeClient(db, store)
	Client.RegisterPost(db)
	Client.PostClient(db, store)
	http.HandleFunc("/createpost", Client.CreatePost)
	Client.Search(db, store)
	Client.LoginPost(db, store)
	Client.HomeClient(db, store)
	Client.HandleCreatepost(db, store)
	Client.RegisterPost(db)
	Client.HandleCreatepost(db, store)
}
