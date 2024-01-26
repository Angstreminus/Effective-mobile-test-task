package main

import (
	"log"

	"github.com/Angstreminus/Effective-mobile-test-task/config"
	"github.com/Angstreminus/Effective-mobile-test-task/internal/server"
)

func main() {
	config, err := config.MustLoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(config)
	server.MustRun()
}
