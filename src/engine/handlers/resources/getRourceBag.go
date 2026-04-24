package resources

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type GetBagResponse struct {
	Message     string               `json:"message"`
	ResourceBag []models.ResourceBag `json:"resourcesBag"`
}

func GetResourceBag(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		var resourceBag []models.ResourceBag
		if err := db.Where("user_id = ?", userID).Find(&resourceBag).Error; err != nil {
			http.Error(w, "failed to load buildings", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		_ = json.NewEncoder(w).Encode(GetBagResponse{
			Message:     "default resources created",
			ResourceBag: resourceBag,
		})
	})
}
