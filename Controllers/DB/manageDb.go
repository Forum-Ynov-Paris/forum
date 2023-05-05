package Forum

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type DBController struct {
	IsInit   bool
	Database *sql.DB
}

func (dbc *DBController) GetUsername(uid int) string {
	row, _ := dbc.Database.Query("SELECT pseudo FROM user WHERE id = ?", uid)
	for row.Next() {
		var pseudo string
		row.Scan(&pseudo)
		return pseudo
	}
	return ""
}

func (dbc *DBController) INIT(databaseName string) error {
	var err error
	dbc.Database, err = sql.Open("sqlite3", databaseName)
	if err != nil {
		return err
	}
	if err = dbc.Database.Ping(); err != nil {
		return err
	}
	dbc.IsInit = true
	return nil
}

func (dbc *DBController) QUERY(query string, args ...interface{}) (*sql.Rows, error) {
	if !dbc.IsInit {
		return nil, errors.New("database not initialized")
	}
	rows, err := dbc.Database.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (dbc *DBController) POST(query string, args ...interface{}) error {
	if !dbc.IsInit {
		return errors.New("database not initialized")
	}
	stmt, err := dbc.Database.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePasswords(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
