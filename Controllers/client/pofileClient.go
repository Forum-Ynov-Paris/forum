package Forum

import (
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"net/http"
)

func ProfileClient(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "forum")

		db.QUERY("SELECT * FROM user WHERE pseudo = ?", session.Values["username"].(string))

	})
}
