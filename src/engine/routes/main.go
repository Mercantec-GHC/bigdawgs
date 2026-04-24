package routes

import (
	"bigdawgs/handlers"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Register(mux *http.ServeMux, db *gorm.DB) {
	auth := handlers.RequireAuth

	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	mux.Handle("GET /buildings", auth(handlers.ListBuildingsHandler(db)))
	mux.Handle("POST /buildings/create", auth(handlers.CreateDefaultBuildingHandler(db)))
	mux.Handle("POST /buildings/{building}/upgrade", auth(handlers.UpgradeBuilding(db)))
}

func ListenAndServe(port string, db *gorm.DB) error {
	mux := http.NewServeMux()
	Register(mux, db)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	return http.ListenAndServe(addr, mux)
}
