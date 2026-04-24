package routes

import (
	"bigdawgs/handlers"
	"bigdawgs/handlers/buildings"
	"bigdawgs/handlers/resources"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Register(mux *http.ServeMux, db *gorm.DB) {
	auth := handlers.RequireAuth

	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	/* Buildings routes */
	mux.Handle("GET /buildings", auth(buildings.ListBuildingsHandler(db)))
	mux.Handle("POST /buildings/create", auth(buildings.CreateDefaultBuilding(db)))
	mux.Handle("POST /buildings/{building}/upgrade", auth(buildings.UpgradeBuilding(db)))

	/* Resources routes */
	mux.Handle("POST /resources/create", auth(resources.CreateDefaultResourceBag(db)))
	mux.Handle("GET /resources/getBag", auth(resources.GetResourceBag(db)))
	mux.Handle("DELETE /resources/deleteBag", auth(resources.DeleteResourceBag(db)))
}

func ListenAndServe(port string, db *gorm.DB) error {
	mux := http.NewServeMux()
	Register(mux, db)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	return http.ListenAndServe(addr, mux)
}
