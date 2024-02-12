package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

//////////////////////////////////////////////////////////////////////

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	//Check whether the auth session is active
	session, _ := h.SessionsStore.Get(r, "auth-session")
	//If so, redirect to the main page
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(w, r, fmt.Sprintf("/%s", session.Values["username"]), http.StatusFound)
		return
	}

	t, err := template.ParseFiles("../frontend/static/auth.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}

//////////////////////////////////////////////////////////////////////

func (h *Handler) MainPage(w http.ResponseWriter, r *http.Request) {
	//Check whether the auth session is active
	session, _ := h.SessionsStore.Get(r, "auth-session")

	//If it isn't, redirect to the auth pag
	if auth, ok := session.Values["authenticated"].(bool); !ok && !auth {

		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}
	username := path.Base(r.URL.Path)

	//selecting the user with such login from the DB
	row := h.DBInstance.QueryRow("SELECT password FROM users WHERE username = ?", username)
	var storedHashedPwd string
	//Scan DB response
	err := row.Scan(&storedHashedPwd)
	if err != nil {
		//If there is not such a login in the database, err value will be sql.ErrNoRows
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username", http.StatusUnauthorized)
		} else {
			//typical error handling
			log.Printf("Unexpected error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	log.Print(storedHashedPwd)
	log.Print(session.Values["password"])
	//if passwords mismatch, send this error, else create session
	if password, ok := session.Values["password"].(string); ok && storedHashedPwd != password {
		http.Error(w, "Hacker detected!", http.StatusUnauthorized)
		return
	}
	t, err := template.ParseFiles("../frontend/static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, r.URL.Path[1:])
}

//////////////////////////////////////////////////////////////////////

func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {
	//Check whether the auth session is active
	session, _ := h.SessionsStore.Get(r, "auth-session")
	//If so, redirect to the main page
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	t, err := template.ParseFiles("../frontend/static/registration.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
}
