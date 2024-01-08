package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ////////////////////////////////////////////////////////////////////////////
func (h *Handler) RegNew(w http.ResponseWriter, r *http.Request) {
	//new variable to store login and password
	var creds Credentials

	//decode the body of the request into the variable (error if wrong json structure)
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//log message when new user is being created
	log.Print("New user: " + creds.Username + " " + hashPwd(creds.Password))
	//set the header, so that the client will know how to deal with the response
	w.Header().Set("Content-Type", "application/json")
	//inserting the user into the database (error if something goes wrong)
	_, err = h.DBInstance.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, creds.Username, hashPwd(creds.Password))
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	//start a session
	session, _ := h.SessionsStore.Get(r, "auth-session")
	//field of a session which represents the fact of users authentification
	session.Values["authenticated"] = true
	//save and send the cookie
	session.Save(r, w)
	json.NewEncoder(w).Encode(map[string]string{"redirect": "/"})
}

// ////////////////////////////////////////////////////////////////////////////
func (h *Handler) AuthTry(w http.ResponseWriter, r *http.Request) {
	//new variable to store login and password
	var creds Credentials
	//decode the body of the request into the variable (error if wrong json structure)
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	//selecting the user with such login from the DB
	row := h.DBInstance.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username)
	var storedHashedPwd string
	//Scan DB response
	err = row.Scan(&storedHashedPwd)
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
	//if passwords mismatch, send this error, else create session
	if storedHashedPwd != hashPwd(creds.Password) {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	//create session
	session, _ := h.SessionsStore.Get(r, "auth-session")
	session.Values["authenticated"] = true
	session.Save(r, w)
	//log message when the user logged in
	log.Print("New session: " + creds.Username + "for " + fmt.Sprint(session.Options.MaxAge))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"redirect": "/"})

}

//////////////////////////////////////////////////////////////////////////////

// small function to hash passwords
func hashPwd(a string) string {
	hasher := sha256.New()
	hasher.Write([]byte(a))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
