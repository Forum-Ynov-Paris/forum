package Forum

//replace the package main by the name of your package and delete main function

import (
	DB "Forum/Controllers/DB"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Commentaire struct {
	Content string `json:"content"`
	Date    string `json:"date"`
	Upvote  int    `json:"upvote"`
	Uuid    int    `json:"uuid"`
}

type Article struct {
	Title       string        `json:"title"`
	Tag         string        `json:"tag"`
	Key         int           `json:"key"`
	Content     string        `json:"content"`
	Upvote      int           `json:"upvote"`
	Date        string        `json:"date"`
	Uuid        int           `json:"uuid"`
	Commentaire []Commentaire `json:"commentaire"`
}

var (
	Path     = "data.json"
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

func GetIndexByTitle(title string) Article {
	Get()
	for _, article := range articles {
		print(article.Title + "/" + title + "|\n")
		if article.Title == title {
			return article
		}
	}
	return Article{}
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

func Post() {
	//marshall the json file
	file, err := os.OpenFile(Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	//marshall the json file
	err = json.NewEncoder(file).Encode(&articles)
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

func AddPost(titre string, tag string, content string, date string, uuid int) {
	//Title       string        `json:"title"`
	//Tag         string        `json:"tag"`
	//Content     string        `json:"content"`
	//Upvote      int           `json:"upvote"`
	//Date        string        `json:"date"`
	//Uuid        int           `json:"uuid"

	// Charger le contenu JSON existant depuis un fichier
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Déclarer une variable pour stocker les données JSON
	var people []Article

	// Désérialiser le contenu JSON dans la variable
	err = json.Unmarshal(file, &people)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	key := rand.Int()
	// Ajouter un nouvel élément au tableau JSON
	newPerson := Article{
		Title:   titre,
		Tag:     tag,
		Content: content,
		Upvote:  0,
		Key:     key,
		Date:    date,
		Uuid:    uuid,
	}
	people = append(people, newPerson)

	// Sérialiser les données en JSON
	newData, err := json.MarshalIndent(people, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}

	// Écrire les données JSON dans le fichier
	err = ioutil.WriteFile("data.json", newData, os.ModePerm)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}
}

func AddComment(key int, date string, content string, uuid int) {
	//Content string `json:"content"`
	//Date    string `json:"date"`
	//Upvote  int    `json:"upvote"`
	//Uuid    int    `json:"uuid"`

	comment := Commentaire{
		Content: content,
		Date:    date,
		Upvote:  0,
		Uuid:    uuid,
	}
	var articless []Article

	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Désérialiser le contenu JSON dans la variable
	err = json.Unmarshal(file, &articless)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}

	var p int
	for m, article := range articless {
		if article.Key == key {
			p = m
		}
	}

	articless[p].Commentaire = append(articless[p].Commentaire, comment)

	// Sérialiser les données en JSON
	newData, err := json.MarshalIndent(articless, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}

	// Écrire les données JSON dans le fichier
	err = ioutil.WriteFile("data.json", newData, os.ModePerm)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}

}

func GetAPIWithKey(key string) Article {
	Get()
	tmp, _ := strconv.Atoi(key)
	for _, article := range articles {
		if article.Key == tmp {
			return article
		}
	}
	return Article{}
}

func UpVote(key int) {
	var articless []Article

	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Désérialiser le contenu JSON dans la variable
	err = json.Unmarshal(file, &articless)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}

	var p int
	for m, article := range articless {
		if article.Key == key {
			p = m
		}
	}

	articless[p].Upvote++

	// Sérialiser les données en JSON
	newData, err := json.MarshalIndent(articless, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}

	// Écrire les données JSON dans le fichier
	err = ioutil.WriteFile("data.json", newData, os.ModePerm)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}
}

func DownVote(key int) {
	var articless []Article

	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	// Désérialiser le contenu JSON dans la variable
	err = json.Unmarshal(file, &articless)
	if err != nil {
		fmt.Println("Erreur lors de la désérialisation JSON:", err)
		return
	}

	var p int
	for m, article := range articless {
		if article.Key == key {
			p = m
		}
	}

	articless[p].Upvote--

	// Sérialiser les données en JSON
	newData, err := json.MarshalIndent(articless, "", "  ")
	if err != nil {
		fmt.Println("Erreur lors de la sérialisation JSON:", err)
		return
	}

	// Écrire les données JSON dans le fichier
	err = ioutil.WriteFile("data.json", newData, os.ModePerm)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}
}
