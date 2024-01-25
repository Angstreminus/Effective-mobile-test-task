package server

import (
	"github.com/Angstreminus/Effective-mobile-test-task/config"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/router"
)

// TODO initialize repos, services, etc

type Server struct {
	Config *config.Config
	Router *router.Router
}

func (s *Server) Run() {
}
