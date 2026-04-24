package resources

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type CreateDefaultResponse struct {
	Message   string               `json:"message"`
	Resources []models.ResourceBag `json:"resources"`
}

func CreateDefaultResourceBag(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		var existing []models.ResourceBag
		result := db.Where("user_id = ?", userID).Find(&existing)
		if result.Error != nil {
			http.Error(w, "failed to check existing resources", http.StatusInternalServerError)
			return
		}
		if len(existing) > 0 {
			http.Error(w, "user already has a resource bag", http.StatusConflict)
			return
		} else {
			defaults := models.DefaultResourceBalances(userID)
			created := make([]models.ResourceBag, 0, len(defaults))

			for _, row := range defaults {
				if err := db.Create(&row).Error; err != nil {
					http.Error(w, "failed to create default resources", http.StatusInternalServerError)
					return
				}

				created = append(created, row)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			_ = json.NewEncoder(w).Encode(CreateDefaultResponse{
				Message:   "default resources created",
				Resources: created,
			})
		}
	})
}
