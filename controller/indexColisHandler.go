package controller

import (
	"net/http"
	"text/template"
)

func IndexColisHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("views/index.html")
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
