package main

import (
	"fmt"
	controller "hackathon/controller"
	hackathon "hackathon/models"
	"net/http"
)

type PageData struct {
	Title   string
	Message string
}

func main() {
	// Gestion des fichiers statiques (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	hackathon.Init()
	// Définition des gestionnaires pour les pages
	http.HandleFunc("/home", controller.IndexColisHandler)

	// Démarrage du serveur
	fmt.Println("Le serveur est lancé sur le port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
