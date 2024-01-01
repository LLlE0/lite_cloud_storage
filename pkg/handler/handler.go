package handler

import (
	"github.com/LLlE0/lite_cloud_storage/pkg/service"
	"github.com/go-chi/chi"
	"net/http"
)

type Handler struct {
	Services *service.Service
	Server   *Server
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the auth page"))
	})

	r.Get("/success", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is the success page"))
	})

	return r
}
