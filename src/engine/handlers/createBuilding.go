package handlers

import (
	"bigdawgs/models"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func CreateDefaultBuildingHandler(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rawUserID := r.PathValue("userID")
		parsedUserID, err := strconv.ParseUint(rawUserID, 10, 64)
		if err != nil || parsedUserID == 0 {
			http.Error(w, "invalid userID", http.StatusBadRequest)
			return
		}

		building := models.Building{
			UserID: uint(parsedUserID),
			Key:    string(models.BuildingDoghouse),
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
