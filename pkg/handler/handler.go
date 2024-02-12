package handler

import (
	"database/sql"
	"github.com/LLlE0/lite_cloud_storage/pkg/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/sessions"
	"net/http"
	"path"
)

type Handler struct {
	//service - running an application
	Services *service.Service
	//server - config the http server
	Server *service.Server
	//dbinstance - operations with database
	DBInstance *sql.DB
	//sessionsstore - store and work with cookies
	SessionsStore *sessions.CookieStore
}

// Constructor of a handler
func NewHandler(services *service.Service, server *service.Server, DB *sql.DB) *Handler {
	return &Handler{Services: services, Server: server, DBInstance: DB, SessionsStore: NewSessionStorage()}
}

// API-Handler itself
func (h *Handler) InitRoutes() *chi.Mux {

	/////////////////////////////////////////////////////////////////////////////////////////////
	//init new router
	r := chi.NewRouter()
	// redirect /auth/ to /auth
	r.Use(middleware.RedirectSlashes)
	//seek for js in the 'js' folder
	fs := http.FileServer(http.Dir("../frontend/js/"))
	//seek for files all around the /frontend/ folder
	r.Handle("/*", fs)
	//Opera GX does not treat js files like js files, idk, but making a getter for any .js file solved this problem
	r.Get("/js/*", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, path.Clean("../frontend"+r.URL.Path))
	})

	//serve all the api-routes

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/auth", h.Auth)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Post("/auth/try", h.AuthTry)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Route("/{user}", func(r chi.Router) {
		r.Get("/", h.MainPage)
		r.Post("/getfile", h.GetFile)
		r.Post("/getfolder", h.GetFolder)

		r.Post("/addfile", h.UploadFile)
		r.Post("/addfolder", h.UploadFolder)

		r.Put("/logout", h.Logout)
	})
	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Get("/registration", h.Registration)

	/////////////////////////////////////////////////////////////////////////////////////////////
	r.Post("/registration/new", h.RegNew)

	return r
}
