package main

import (
	"log"
	"os"

	"bigdawgs/db"
	"bigdawgs/handlers"
	"bigdawgs/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	database, err := db.Connect()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	go handlers.RunTickLoop(database)

	log.Printf("engine listening on 0.0.0.0:%s", port)

	if err := routes.ListenAndServe(port, database); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
