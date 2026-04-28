package buildings

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func ListBuildingsHandler(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
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
