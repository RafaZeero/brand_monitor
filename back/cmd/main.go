package main

import (
	"log"

	"github.com/RafaZeero/brand_monitor/cmd/api"
)

func main() {
	server := api.NewApiServer("3333")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
