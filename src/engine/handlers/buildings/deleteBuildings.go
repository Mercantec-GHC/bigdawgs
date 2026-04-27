package buildings

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type DeleteBuildingsResponse struct {
	Message string `json:"message"`
}

func DeleteBuildings(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		if err := db.Unscoped().Where("user_id = ?", userID).Delete(&models.Building{}).Error; err != nil {
			http.Error(w, "failed to delete Dawg town", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		_ = json.NewEncoder(w).Encode(DeleteBuildingsResponse{
			Message: "Dawg town where deleted",
		})
	})
}
