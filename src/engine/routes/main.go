package routes

import (
	"bigdawgs/handlers"
	"bigdawgs/handlers/buildings"
	"bigdawgs/handlers/resources"
	"bigdawgs/handlers/transactions"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Register(mux *http.ServeMux, db *gorm.DB) {
	auth := handlers.RequireAuth

	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	/* Buildings routes */
	mux.Handle("GET /buildings", auth(buildings.ListBuildingsHandler(db)))
	mux.Handle("GET /buildings/{building}", auth(buildings.GetSpecificBuilding(db)))
	mux.Handle("GET /buildings/{building}/{uid}", auth(buildings.GetUserSpecificBuilding(db)))
	mux.Handle("POST /buildings/create", auth(buildings.CreateDefaultBuilding(db)))
	mux.Handle("POST /buildings/{building}/upgrade", auth(buildings.UpgradeBuilding(db)))
	mux.Handle("POST /buildings/{building}/create", auth(buildings.CreateBuilding(db)))
	mux.Handle("DELETE /buildings/deleteAll", auth(buildings.DeleteBuildings(db)))

	/* Resources routes */
	mux.Handle("POST /resources/create", auth(resources.CreateDefaultResourceBag(db)))
	mux.Handle("GET /resources/getBag", auth(resources.GetResourceBag(db)))
	mux.Handle("GET /resources/getCap", auth(resources.GetResourceCap(db)))
	mux.Handle("DELETE /resources/deleteBag", auth(resources.DeleteResourceBag(db)))

	/* Transaction routes */
	mux.Handle("POST /transaction/trade", auth(transactions.TradeResources(db)))
}

func ListenAndServe(port string, db *gorm.DB) error {
	mux := http.NewServeMux()
	Register(mux, db)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	return http.ListenAndServe(addr, mux)
}
