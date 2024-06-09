package server

import (
	"encoding/json"
	"log"
	"net/http"

	"movie-tracker/cmd/web"
	"movie-tracker/internal/controller"
	"movie-tracker/libs"

	"github.com/a-h/templ"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.Handle("/", libs.IsAuthenticated((templ.Handler(web.Home()))))
	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/web", templ.Handler(web.HelloForm()))
	mux.Handle("/login-view", libs.IsAuthenticated(templ.Handler(web.Login())))
	mux.HandleFunc("POST /login", controller.Auth)

	mux.HandleFunc("/hello", web.HelloWebHandler)
	mux.HandleFunc("/health", s.healthHandler)
	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
