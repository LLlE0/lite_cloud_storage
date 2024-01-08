package handler

import (
	"html/template"
	"log"
	"net/http"
)

//////////////////////////////////////////////////////////////////////

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	//Check whether the auth session is active
	session, _ := h.SessionsStore.Get(r, "auth-session")
	//If so, redirect to the main page
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(w, r, "/", http.StatusFound)
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
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		http.Redirect(w, r, "/auth", http.StatusFound)
		return
	}

	t, err := template.ParseFiles("../frontend/static/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "")
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
