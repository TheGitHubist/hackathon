package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	http.HandleFunc("/home", UserPageHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/backend", backHandler)
	fmt.Println("le serveur est lancer sur le port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fs := http.FileServer(http.Dir("../styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("../templates/admin.html")
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

func backHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("../templates/back.html")
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
