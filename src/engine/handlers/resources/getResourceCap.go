package resources

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

func GetResourceCap(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		var doghouse models.Building
		if err := db.Where("user_id = ? AND key = ?", userID, string(models.Doghouse)).First(&doghouse).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				http.Error(w, "doghouse not found", http.StatusNotFound)
				return
			}
			http.Error(w, "failed to load doghouse", http.StatusInternalServerError)
			return
		}

		cap := models.ResourceCap(doghouse.Level)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(cap)
	})
}
