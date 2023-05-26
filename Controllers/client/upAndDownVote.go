package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"net/http"
	"strconv"
)

func UpAndDownVote(db DB.DBController, store *sessions.CookieStore) {
	http.HandleFunc("/up/", func(w http.ResponseWriter, r *http.Request) {
		key, _ := strconv.Atoi(r.URL.String()[4:])
		API.UpVote(key)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
	http.HandleFunc("/down/", func(w http.ResponseWriter, r *http.Request) {
		key, _ := strconv.Atoi(r.URL.String()[4:])
		API.DownVote(key)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})
}
