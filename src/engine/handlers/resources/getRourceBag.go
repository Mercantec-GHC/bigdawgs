package resources

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type GetBagResponse struct {
	Message      string                        `json:"message"`
	ResourceBags map[string]models.ResourceBag `json:"resourcesBag"`
}

func GetResourceBag(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		var resourceBags []models.ResourceBag
		if err := db.Where("user_id = ?", userID).Find(&resourceBags).Error; err != nil {
			http.Error(w, "failed to load resource bag", http.StatusInternalServerError)
			return
		}

		bagMap := make(map[string]models.ResourceBag, len(resourceBags))
		for _, r := range resourceBags {
			bagMap[r.ResourceKey] = r
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(GetBagResponse{
			Message:      "resource bag loaded",
			ResourceBags: bagMap,
		})
	})
}
