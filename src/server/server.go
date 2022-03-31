package server

import (
	"github.com/gorilla/mux"
	"github.com/reeves122/micro-airlines-api-go/handler"
	"github.com/reeves122/micro-airlines-api-go/repository"
	"log"
	"net/http"
)

type server struct {
	router *mux.Router
	repo   repository.IRepository
}

func NewServer(repo repository.IRepository) *server {
	router := mux.NewRouter()
	return &server{
		router: router,
		repo:   repo,
	}
}

func (s *server) Run(host string) {
	s.setRouters()
	log.Fatal(http.ListenAndServe(host, s.router))
}

func (s *server) setRouters() {
	s.Get("/api/healthcheck", s.GetHealthCheck)
	s.Get("/api/players", s.GetAllPlayers)
	s.Post("/api/players", s.CreatePlayer)
	//s.Get("/api/players/{username}", s.GetPlayer)
	//s.Put("/api/players/{username}", s.UpdatePlayer)
	//s.Delete("/api/players/{username}", s.DeletePlayer)
	//s.Put("/api/players/{username}/disable", s.DisablePlayer)
	//s.Put("/api/players/{username}/enable", s.EnablePlayer)
}

func (s *server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("GET")
}

func (s *server) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	s.router.HandleFunc(path, f).Methods("POST")
}

//func (s *server) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	s.router.HandleFunc(path, f).Methods("PUT")
//}
//
//func (s *server) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	s.router.HandleFunc(path, f).Methods("DELETE")
//}

func (s *server) GetHealthCheck(w http.ResponseWriter, r *http.Request) {
	handler.HealthCheck(s.repo, w, r)
}

func (s *server) GetAllPlayers(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPlayers(s.repo, w, r)
}

func (s *server) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	handler.CreatePlayer(s.repo, w, r)
}

//func (s *server) GetPlayer(w http.ResponseWriter, r *http.Request) {
//	handler.GetPlayer(s.repo, w, r)
//}
//
//func (s *server) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
//	handler.UpdatePlayer(s.repo, w, r)
//}
//
//func (s *server) DeletePlayer(w http.ResponseWriter, r *http.Request) {
//	handler.DeletePlayer(s.repo, w, r)
//}
//
//func (s *server) DisablePlayer(w http.ResponseWriter, r *http.Request) {
//	handler.DisablePlayer(s.repo, w, r)
//}
//
//func (s *server) EnablePlayer(w http.ResponseWriter, r *http.Request) {
//	handler.EnablePlayer(s.repo, w, r)
//}
