package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func (h *Handler) GetFile(w http.ResponseWriter, r *http.Request) {
	fileName := "../_storage/" + strings.Split(r.URL.Path, "/")[1] + "/" + strings.Join((strings.Split(r.URL.Path, "/")[3:]), "/")
	log.Print("GetFile " + fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Print("File not found: ", fileName)
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Error sending file", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetFolder(w http.ResponseWriter, r *http.Request) {
	log.Print("GetFolder")
}

func (h *Handler) GetFolderData(w http.ResponseWriter, r *http.Request) {

	pathToCheck := viper.GetString("store") + r.URL.Path
	_, err := os.Stat(pathToCheck)
	log.Print(err)
	if err != nil {
		log.Print("Path " + pathToCheck + " does not exist!")
		json.NewEncoder(w).Encode(map[string][]string{"str": {}})
		return
	}
	log.Print("GetFolderData for " + r.URL.Path)
	dirContent, err := os.ReadDir(pathToCheck)
	if err == nil {
		log.Print(dirContent)
		var fs []string
		for _, val := range dirContent {
			fs = append(fs, val.Name())
		}
		json.NewEncoder(w).Encode(map[string][]string{"str": fs})
	} else {
		log.Print(err)
		json.NewEncoder(w).Encode(map[string][]string{"str": {}})
	}
}

func (h *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	log.Print("UploadFile")
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Print(err)
	}
	defer file.Close()
	namearr := strings.Split(header.Filename, "/")
	name := namearr[len(namearr)-1]
	out, err := os.Create(viper.GetString("store") + strings.Join(strings.Split(r.URL.Path, "/")[3:], "/") + "/" + name)
	if err != nil {
		log.Print(err)
	}
	defer out.Close()
	log.Print(strings.Join(strings.Split(r.URL.Path, "/")[1:], "/"))
	_, err = io.Copy(out, file)
	if err != nil {
		log.Print(err)
	}
	json.NewEncoder(w).Encode(string("File uploaded successfully"))
	log.Print("File uploaded successfully")
}

func (h *Handler) AddFolder(w http.ResponseWriter, r *http.Request) {
	log.Print("AddFolder")
}
