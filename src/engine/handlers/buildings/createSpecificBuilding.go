package buildings

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateBuilding(db *gorm.DB) http.Handler {
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
		}).FirstOrCreate(&building)

		if result.Error != nil {
			http.Error(w, "failed to create building", http.StatusInternalServerError)
			return
		}

		status := http.StatusCreated
		if result.RowsAffected == 0 {
			status = http.StatusOK
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(building)
	})
}
