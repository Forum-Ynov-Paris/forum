package Forum

import (
	API "Forum/Controllers/API"
	DB "Forum/Controllers/DB"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	ID      int
	uid     int
	content string
)

func InitPostClient(db DB.DBController, store *sessions.CookieStore) {
	type commentary struct {
		C        API.Commentaire
		Username string
	}

	type data struct {
		Post        API.Article
		Connected   bool
		Commentates []commentary
		Name        string
	}
	http.HandleFunc("/post/", func(w http.ResponseWriter, r *http.Request) {

		Title := r.URL.String()[6:]

		session, _ := store.Get(r, "forum")

		//row, err := db.QUERY("SELECT id FROM user WHERE pseudo = ?", session.Values["username"].(string))
		//if err != nil {
		//	log.Fatal(err)
		//} else {
		//	for row.Next() {
		//		err = row.Scan(&ID)
		//		if err != nil {
		//			log.Fatal(err)
		//		}
		//	}
		//}

		Data := data{
			API.GetAPIWithKey(Title), //changer + tard
			true,
			make([]commentary, 0),
			session.Values["username"].(string),
		}

		for _, c := range Data.Post.Commentaire {
			Data.Commentates = append(Data.Commentates, commentary{c, db.GetUsername(c.Uuid)})
		}

		t, err := template.ParseFiles("./static/post.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, Data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	http.HandleFunc("/api/post/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			Title := r.URL.String()[10:]
			comment := r.FormValue("newComment")
			//session, _ := store.Get(r, "forum")

			session, _ := store.Get(r, "forum")
			username := session.Values["username"].(string)

			row, err := db.QUERY("SELECT id FROM user WHERE pseudo = ?", username)
			if err != nil {
				log.Fatal(err)
			} else {
				for row.Next() {
					var id int
					err = row.Scan(&id)
					if err != nil {
						log.Fatal(err)
					}
					currentTime := time.Now()
					var date string
					for _, i := range currentTime.String() {
						if i != '.' {
							date += string(i)
						} else {
							break
						}
					}
					key, _ := strconv.Atoi(Title)
					API.AddComment(key, date, comment, id)
				}
			}
		}
		red := "/post/" + r.URL.String()[10:]
		http.Redirect(w, r, red, http.StatusSeeOther)
	})
}
