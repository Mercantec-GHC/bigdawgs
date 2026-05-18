package buildings

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type GetBuildingResponse struct {
	Message  string          `json:"message"`
	Building models.Building `json:"building"`
}

func GetSpecificBuilding(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		buildingKey := r.PathValue("building")
		if !models.IsValidBuildingKey(buildingKey) {
			http.Error(w, "invalid building key", http.StatusBadRequest)
			return
		}
		building := models.Building{
			UserID: userID,
			Key:    buildingKey,
		}

		result := db.Where(models.Building{
			UserID: building.UserID,
			Key:    building.Key,
		}).First(&building)

		if result.Error != nil {
			http.Error(w, "building not found", http.StatusNotFound)
			return
		}

		var doghouse models.Building
		if err := db.Where("user_id = ? AND key = ?", userID, string(models.Doghouse)).First(&doghouse).Error; err == nil {
			building.DoghouseLevel = doghouse.Level
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(GetBuildingResponse{
			Message:  "resource bag loaded",
			Building: building,
		})
	})
}
