package main

import (
	"log"

	"github.com/w1png/htmx/config"
	"github.com/w1png/htmx/storage"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = storage.InitStorage()
	if err != nil {
		log.Fatal(err)
		return
	}

	server := NewHttpServer()

	log.Fatal(server.Run())
}
