package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int
	Username string
	Email    string `gorm:"typevarchar(100)"`
	Password string
}

type UserDetails struct {
	ID      int
	Name    string
	Address string
	Phone   string `gorm:"unique_index"`
	UserID  int
}

var Db *sql.DB

func Init() *sql.DB {
	dbURL := "postgres://sarisha:sarisha.123@localhost:5432/goauthentication?sslmode=disable"
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Connected to database")
	}
	Db = db
	return db
}
