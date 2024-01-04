package handler

import (
	"encoding/json"
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

func RegNew(w http.ResponseWriter, r *http.Request) {
	log.Print("reg/new call")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Print(creds)
	w.Header().Set("Content-Type", "application/json")
	//ADD TO DB
	json.NewEncoder(w).Encode(map[string]string{"redirect": "/"})

}

func AuthTry(w http.ResponseWriter, r *http.Request) {
	log.Print("auth/try call")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Print(creds)
	if creds.Username == "1" && creds.Password == "1" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"redirect": "/"})
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
