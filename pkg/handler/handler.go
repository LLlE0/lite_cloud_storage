package handler

import (
	"database/sql"
	"github.com/LLlE0/lite_cloud_storage/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"path"
)

type Handler struct {
	Services   *service.Service
	Server     *service.Server
	DBInstance *sql.DB
}

func NewHandler(services *service.Service, server *service.Server, DB *sql.DB) *Handler {
	return &Handler{Services: services, Server: server, DBInstance: DB}
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
	r.Get("/auth", Auth)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Post("/auth/try", h.AuthTry)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/", MainPage)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/registration", Registration)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Post("/registration/new", h.RegNew)

	return r
}
