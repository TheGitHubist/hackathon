package main

import (
	"fmt"
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
	http.HandleFunc("/home", UserPageHandler)
	fmt.Println("le serveur est lancer sur le port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	fs := http.FileServer(http.Dir("../styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))
}
