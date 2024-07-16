package main

import (
	"context"
	"log"

	"github.com/RafaZeero/brand_monitor/cmd/api"
	"github.com/RafaZeero/brand_monitor/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	client, err := db.NewDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	initStorage(client)

	server := api.NewApiServer("3333", client)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *mongo.Client) {
	if db.Ping(context.TODO(), nil) != nil {
		log.Fatal("failed to connect to database")
	}

	log.Println("connected to database")
}
