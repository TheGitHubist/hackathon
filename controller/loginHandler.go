package controller

import (
	"fmt"
	models "hackathon/models"
	"net/http"
	"text/template"
	"time"
)

type Data struct {
	Message string
}

var data Data

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("views/login.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	fmt.Print(username)
	password := r.Form.Get("password")
	fmt.Print(password)
	correct, userid := models.LoginCheck(username, password)

	fmt.Println("Ben")
	if correct {
		fmt.Println("rompich")
		expiresAt := time.Now().Add(120 * time.Minute)
		uuid := models.GetUserUUID(userid)

		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   uuid,
			Expires: expiresAt,
		})
	} else {
		fmt.Println("else")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

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
