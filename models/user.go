package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID            int
	Username          string
	Prenom            string
	Nom               string
	Age               int
	Telephone         string
	NombreCommandes   string
	NombreCovoiturage string
	GoFast            bool
	KmParcouru        int
	Pays              string
	Adresse           string
	Ville             string
	CodePostal        string
}

func Register(username string, passwd string, email string) {
	newpasswd := HashString(passwd)
	uuid := uuid.NewString()
	NewUser(username, email, newpasswd, uuid)
}

func LoginCheck(username string, passwd string) (bool, int) {
	var id int
	var password string
	DB.QueryRow("SELECT password FROM user WHERE email = ?", username).Scan(&password)
	DB.QueryRow("SELECT id FROM user WHERE email = ?", username).Scan(&id)

	if HashString(passwd) == password {
		return true, id
	}
	DB.QueryRow("SELECT password FROM user WHERE username = ?", username).Scan(&password)
	DB.QueryRow("SELECT id FROM user WHERE username = ?", username).Scan(&id)
	if HashString(passwd) == password {
		return true, id
	}
	return false, 0
}

func GetUserUUID(id int) string {
	var uuid string
	DB.QueryRow("SELECT uuid, id FROM user WHERE id = ?", id).Scan(&uuid)
	return uuid
}

func LoadUserApi(uuid string) User {
	var user User
	DB.QueryRow("SELECT username, id, FROM user WHERE uuid = ?", uuid).Scan(&user.Username, &user.UserID)
	return user
}
