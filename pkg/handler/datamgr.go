package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	log.Print("GetFile")
}

func (h *Handler) GetFolder(w http.ResponseWriter, r *http.Request) {
	log.Print("GetFolder")
}

func (h *Handler) GetFolderData(w http.ResponseWriter, r *http.Request) {
	log.Print("GetFolderData")
	json.NewEncoder(w).Encode(map[string][]string{"str": {"1", "2"}})
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Print("UploadFile")
}

func (h *Handler) AddFolder(w http.ResponseWriter, r *http.Request) {
	log.Print("AddFolder")
}
