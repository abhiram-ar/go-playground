package main

import (
	"fmt"
	"log"
	"note-api/internal/config"
	"note-api/internal/db"
	"note-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config error")
	}

	client, database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("db error")
	}
	defer func() {
		if err := db.Disconnect(client); err != nil {
			log.Printf("Mongo disconnect error %v", err)
		}
	}()

	router := server.NewRouter(database)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server failed")
	}

}
