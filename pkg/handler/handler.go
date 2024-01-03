package handler

import (
	"github.com/LLlE0/lite_cloud_storage/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"html/template"
	"log"
	"net/http"
	"path"
)

type Handler struct {
	Services *service.Service
	Server   *Server
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {

	/////////////////////////////////////////////////////////////////////////////////////////////
	r := chi.NewRouter()
	r.Use(middleware.RedirectSlashes)
	fs := http.FileServer(http.Dir("../frontend/js/"))
	r.Handle("/*", fs)
	r.Get("/js/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Clean("../frontend"+r.URL.Path))
	})

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("../frontend/static/auth.html")
		if err != nil {
			log.Fatal(err)
		}

		t.Execute(w, "")

	})

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Post("/auth/try", func(w http.ResponseWriter, r *http.Request) {

		log.Print(r.Body)
		log.Print("ТОЧНО ЭТОТ ВЫЗВАЛСЯ")
		http.Redirect(w, r, "/index", http.StatusSeeOther)

	})

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/index", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("../frontend/static/index.html")
		if err != nil {
			log.Fatal(err)
		}

		t.Execute(w, "")
	})

	return r
}
