package handler

import (
	"html/template"
	"log"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../frontend/static/auth.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../frontend/static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}

func Registration(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../frontend/static/registration.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}
