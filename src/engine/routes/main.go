package routes

import (
	"bigdawgs/handlers"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Register(mux *http.ServeMux, db *gorm.DB) {
	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	mux.Handle("POST /users/{userID}/buildings/create", handlers.CreateDefaultBuildingHandler(db))
}

func ListenAndServe(port string, db *gorm.DB) error {
	mux := http.NewServeMux()
	Register(mux, db)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	return http.ListenAndServe(addr, mux)
}
