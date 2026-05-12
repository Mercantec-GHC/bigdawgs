package buildings

import (
	"bigdawgs/models"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type GetUserBuildingResponse struct {
	Message  string          `json:"message"`
	Building models.Building `json:"building"`
}

func GetUserSpecificBuilding(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		uidStr := r.PathValue("uid")
		if uidStr == "" {
			http.Error(w, "uid not provided", http.StatusBadRequest)
			return
		}
		userIDToUse, err := strconv.Atoi(uidStr)
		if err != nil {
			http.Error(w, "invalid uid", http.StatusBadRequest)
			return
		}

		buildingKey := r.PathValue("building")
		if !models.IsValidBuildingKey(buildingKey) {
			http.Error(w, "invalid building key", http.StatusBadRequest)
			return
		}
		building := models.Building{
			UserID: uint(userIDToUse),
			Key:    buildingKey,
		}

		result := db.Where(models.Building{
			UserID: building.UserID,
			Key:    building.Key,
		}).First(&building)

		if result.Error != nil {
			http.Error(w, "building not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(GetBuildingResponse{
			Message:  "resource bag loaded",
			Building: building,
		})
	})
}
