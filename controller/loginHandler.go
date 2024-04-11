package controller

import (
	models "hackathon/models"
	"net/http"
	"time"
)

type Data struct {
	Message string
}

var data Data

func loginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	models.LoginCheck(username, password)

}

func disconnectHandler(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(w, c)
	http.Redirect(w, r, "/profil", http.StatusSeeOther)
}
