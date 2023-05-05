package Forum

//replace the package main by the name of your package and delete main function

import (
	DB "Forum/Controllers/DB"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

type Commentaire struct {
	Content string `json:"content"`
	Uuid    int    `json:"uuid"`
}

type Article struct {
	Title       string        `json:"title"`
	Tag         string        `json:"tag"`
	Content     string        `json:"content"`
	Upvote      int           `json:"upvote"`
	Date        string        `json:"date"`
	Uuid        int           `json:"uuid"`
	Commentaire []Commentaire `json:"commentaire"`
}

var (
	Path     = "./data.json"
	articles []Article
)

func GetArticles() []Article {
	Get()
	return articles
}

func GetArticle(id int) Article {
	Get()
	return articles[id]
}

func Get() {
	//open ./data.json and unmarshall it
	file, err := os.Open(Path)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//unmarshall the json file
	err = json.NewDecoder(file).Decode(&articles)
	if err != nil {
		log.Fatal(err)
	}
}

func PostArticle(article Article) {
	Get()
	articles = append(articles, article)

	jsonData, err := json.Marshal(articles)
	if err != nil {
		panic(err)
	}

	//open ./data.json and write the new json
	file, err := os.OpenFile(Path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
}

func SearchArticles(search string, db DB.DBController) []Article {
	Get()
	var articlesSearch []Article
	for _, article := range articles {
		if article.Title == search || article.Tag == search || strings.Contains(article.Title, search) {
			articlesSearch = append(articlesSearch, article)
		}
	}
	row, err := db.QUERY("SELECT id FROM user WHERE pseudo = ?", search)
	if err != nil {
		log.Fatal(err)
	} else {
		for row.Next() {
			var id int
			err = row.Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
			for _, article := range articles {
				if article.Uuid == id {
					articlesSearch = append(articlesSearch, article)
				}
			}
		}
	}
	return articlesSearch
}
