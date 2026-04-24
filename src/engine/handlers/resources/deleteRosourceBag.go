package resources

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type DeleteBagResponse struct {
	Message string `json:"message"`
}

func DeleteResourceBag(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
		if err != nil {
			http.Error(w, "missing authenticated user", http.StatusUnauthorized)
			return
		}

		if err := db.Unscoped().Where("user_id = ?", userID).Delete(&models.ResourceBag{}).Error; err != nil {
			fmt.Println(err)
			http.Error(w, "failed to delete resource bag", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		_ = json.NewEncoder(w).Encode(DeleteBagResponse{
			Message: "Dawgbags where deleted",
		})
	})
}
