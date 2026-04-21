package main

import (
	"log"
	"net/http"
	"os"

	"bigdawgs/db"
)

func main() {
	if _, err := db.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	addr := "0.0.0.0:" + port
	log.Printf("engine listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
