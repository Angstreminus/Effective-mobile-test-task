package main

import (
	"log"

	"github.com/Angstreminus/Effective-mobile-test-task/config"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/server"
)

func main() {
	Config, err := config.MustLoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	Server := server.NewServer(Config)
	Server.MustRun()
}
