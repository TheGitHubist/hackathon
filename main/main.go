package main

import (
	"html/template"
	"net/http"
)

func UserPageHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("../templates/user.html")
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

func main() {

}
