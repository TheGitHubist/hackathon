package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	DB = LoadDataBase()
}

func LoadDataBase() *sql.DB {
	db, _ := sql.Open("sqlite3", "data.db")
	data, err := db.Prepare("CREATE TABLE IF NOT EXISTS user(id INTEGER PRIMARY KEY AUTOINCREMENT, email TEXT NOT NULL UNIQUE, password TEXT NOT NULL, salt INTEGER NOT NULL, uuid TEXT NOT NULL UNIQUE)")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec()
	data, err = db.Prepare("CREATE TABLE IF NOT EXISTS admin(id INTEGER PRIMARY KEY AUTOINCREMENT, adminPass TEXT NOT NULL, accountID BLOB NOT NULL, FOREIGN KEY (accountID) REFERENCES user(id))")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec()
	data, err = db.Prepare("CREATE TABLE IF NOT EXISTS backOfficer(id INTEGER PRIMARY KEY AUTOINCREMENT, backOffPass TEXT NOT NULL, accountID BLOB NOT NULL, salt INTEGER NOT NULL, FOREIGN KEY (accountID) REFERENCES user(id))")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec()
	data, err = db.Prepare("CREATE TABLE IF NOT EXISTS product(id INTEGER PRIMARY KEY AUTOINCREMENT, number TEXT NOT NULL, status INTEGER NOT NULL, officerID BLOB NOT NULL, ownerID BLOB NOT NULL, FOREIGN KEY (officerID) REFERENCES backOfficer(id), FOREIGN KEY (ownerID) REFERENCES user(id))")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec()
	return db
}

func NewUser(email string, password int, salt int, uuid string) int {

	// returns 0 if everything's fine, 1 for pseudo or uuid not unique, 2 for another db error
	db, err := sql.Open("sqlite3", "")
	if err != nil {
		return 2
	}
	rows, err := db.Query("SELECT email FROM user")
	if err != nil {
		return 2
	}
	defer rows.Close()
	for rows.Next() {
	}
	data, err := db.Prepare("INSERT INTO user(email, password, uuid) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec(email, password, salt, uuid)
	defer data.Close()
	return 0
}

func NewAdmin(password string, accountID int) int {
	// returns 0 if everything's fine, 1 for pseudo or uuid not unique, 2 for another db error
	db, err := sql.Open("sqlite3", "../databases/database.db")
	if err != nil {
		return 2
	}
	data, err := db.Prepare("INSERT INTO admin(password, accountID) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec(password, accountID)
	defer data.Close()
	return 0
}

func NewBackOfficer(password string, salt int, accountID int) int {
	// returns 0 if everything's fine, 1 for pseudo or uuid not unique, 2 for another db error
	db, err := sql.Open("sqlite3", "../databases/database.db")
	if err != nil {
		return 2
	}
	data, err := db.Prepare("INSERT INTO backOfficer(password, salt, accountID) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec(password, salt, accountID)
	defer data.Close()
	return 0
}

func NewProduct(number string, status int, officerID int, ownerID int) int {
	// returns 0 if everything's fine, 1 for pseudo or uuid not unique, 2 for another db error
	db, err := sql.Open("sqlite3", "../databases/database.db")
	if err != nil {
		return 2
	}
	data, err := db.Prepare("INSERT INTO user(number, status, officerID, ownerID) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	data.Exec(number, status, officerID, ownerID)
	defer data.Close()
	return 0
}
