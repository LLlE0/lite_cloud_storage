package handler

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

// ////////////////////////////////////////////////////////////////////////////
func (h *Handler) RegNew(w http.ResponseWriter, r *http.Request) {
	log.Print("reg/new call")

	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Print(creds.Username, hashPwd(creds.Password))
	w.Header().Set("Content-Type", "application/json")
	_, err = h.DBInstance.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`, creds.Username, hashPwd(creds.Password))
	if err != nil {
		log.Fatalf("Failed to execute query: %v", err)
	}
	log.Print(creds.Username, hashPwd(creds.Password))
	json.NewEncoder(w).Encode(map[string]string{"redirect": "/"})
}

// ////////////////////////////////////////////////////////////////////////////
func (h *Handler) AuthTry(w http.ResponseWriter, r *http.Request) {
	log.Print("auth/try call")
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Print(creds)
	row := h.DBInstance.QueryRow("SELECT password FROM users WHERE username = ?", creds.Username)
	var storedHashedPwd string
	err = row.Scan(&storedHashedPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			log.Printf("Unexpected error: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	if storedHashedPwd != hashPwd(creds.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"redirect": "/"})
}

//////////////////////////////////////////////////////////////////////////////

func hashPwd(a string) string {
	hasher := sha256.New()
	hasher.Write([]byte(a))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
