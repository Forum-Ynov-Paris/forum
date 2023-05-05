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
<<<<<<< HEAD
	http.HandleFunc("/post", Client.Post)
	Client.LoginPost(db, store)
	Client.HomeClient(db, store)
	Client.RegisterPost(db)
	Client.PostClient(db, store)
=======
	http.HandleFunc("/createpost", Client.createPost)
	Client.LoginPost(db, store)
	Client.HomeClient(db, store)
	Client.RegisterPost(db)
	Client.handleCreatepost(db, store)

>>>>>>> 7aa84e417f315ae5f54899f84c44050c826e1d4c
}
