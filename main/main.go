package main

<<<<<<< HEAD
import "net/http"

func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	//http.HandleFunc("/", indexPath)

	http.ListenAndServe("", nil)
=======
import (
	"fmt"
	controllers "hackathon/controller"
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	// Gestion des fichiers statiques (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("../styles"))
	http.Handle("/styles/", http.StripPrefix("/styles/", fs))
	controllers.Init()
	// Définition des gestionnaires pour les pages
	http.HandleFunc("/home", UserPageHandler)
	http.HandleFunc("/admin", adminHandler)
	http.HandleFunc("/backend", backHandler)

	// Démarrage du serveur
	fmt.Println("Le serveur est lancé sur le port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/admin.html")
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
	templ, err := template.ParseFiles("templates/back.html")
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
	templ, err := template.ParseFiles("templates/user.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
>>>>>>> e6ef7811a72afa7586d7e645f55222e06c2f9bca
}
