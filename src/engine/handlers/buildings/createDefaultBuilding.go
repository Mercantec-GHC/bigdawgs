package buildings

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func CreateDefaultBuilding(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		buildings := make([]models.Building, 0, len(models.BuildingDefinitions))
		status := http.StatusCreated

		for _, building := range models.BuildingDefinitions {
			building := models.Building{
				UserID: userID,
				Key:    string(building.Key),
			}

			if building.Key == string(models.Doghouse) {
				building.Level = 1
			}

			result := db.Where(models.Building{
				UserID: building.UserID,
				Key:    building.Key,
			}).FirstOrCreate(&building)

			if result.Error != nil {
				http.Error(w, "failed to create building", http.StatusInternalServerError)
				return
			}

			if result.RowsAffected == 0 {
				status = http.StatusOK
			}

			buildings = append(buildings, building)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_ = json.NewEncoder(w).Encode(buildings)
	})
}
