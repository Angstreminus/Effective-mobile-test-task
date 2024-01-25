package server

import (
	"log"
	"net/http"

	"github.com/Angstreminus/Effective-mobile-test-task/config"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/handler"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/postgres"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/repository"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/service"
)

type Server struct {
	Config *config.Config
	Router *http.ServeMux
}

func NewServer(config *config.Config) *Server {
	return &Server{
		Config: config,
	}
}

func (s *Server) MustRun() {
	dbHandler, err := postgres.NewDatabaseHandler(s.Config)
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewUserRepository(dbHandler)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)
	s.Router.HandleFunc("/users/create", handler.CreateUser)
	s.Router.HandleFunc("/users", handler.GetUsersHandler)
	if err = http.ListenAndServe(s.Config.ServerAddr, nil); err != nil {
		log.Fatal(err)
	}
}
