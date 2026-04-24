package handlers

import (
	"bigdawgs/models"
	"encoding/json"
	"net/http"
	"time"

	"gorm.io/gorm"
)

const upgradeDuration = time.Second * 30

func ListBuildingsHandler(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		var buildings []models.Building
		if err := db.Where("user_id = ?", userID).Order("created_at ASC").Find(&buildings).Error; err != nil {
			http.Error(w, "failed to load buildings", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(buildings)
	})
}

func UpgradeBuilding(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		buildingKey := r.PathValue("building")
		if !models.IsValidBuildingKey(buildingKey) {
			http.Error(w, "invalid building name", http.StatusBadRequest)
			return
		}

		var building models.Building
		if err := db.Where("user_id = ? AND key = ?", userID, buildingKey).First(&building).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "building not found", http.StatusNotFound)
				return
			}

			http.Error(w, "failed to load building", http.StatusInternalServerError)
			return
		}

		if building.IsConstructing {
			http.Error(w, "building is already upgrading", http.StatusConflict)
			return
		}

		now := time.Now().UTC()
		completesAt := now.Add(upgradeDuration)
		updates := map[string]any{
			"is_constructing": true,
			"started_at":      now,
			"completes_at":    completesAt,
		}

		if err := db.Model(&building).Updates(updates).Error; err != nil {
			http.Error(w, "failed to start building upgrade", http.StatusInternalServerError)
			return
		}

		building.IsConstructing = true
		building.StartedAt = &now
		building.CompletesAt = &completesAt

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(building)
	})
}
