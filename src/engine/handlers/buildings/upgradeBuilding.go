package buildings

import (
	"bigdawgs/handlers"
	"bigdawgs/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

const upgradeDuration = time.Second * 30

var ErrInsufficientResources = errors.New("insufficient resources")

type UpgradeBuildingResponse struct {
	Message  string            `json:"message"`
	Building models.Building   `json:"building"`
	CostPaid models.Production `json:"cost_paid"`
}

func UpgradeBuilding(db *gorm.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID, err := handlers.UserID(r)
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

		if models.BuildingKey(buildingKey) != models.Doghouse {
			var doghouse models.Building
			if err := db.Where("user_id = ? AND key = ?", userID, string(models.Doghouse)).First(&doghouse).Error; err != nil {
				http.Error(w, "failed to load doghouse", http.StatusInternalServerError)
				return
			}
			if building.Level+1 > doghouse.Level {
				http.Error(w, "doghouse needs to be a higher level before upgrading this building", http.StatusUnprocessableEntity)
				return
			}
		}

		if building.IsConstructing {
			http.Error(w, "building is already upgrading", http.StatusConflict)
			return
		}

		cost := building.UpgradeCost()

		txErr := db.Transaction(func(tx *gorm.DB) error {
			type costEntry struct {
				key    string
				amount int64
			}
			costs := []costEntry{
				{string(models.ResourceDogCoin), cost.DogCoins},
				{string(models.ResourceDogBone), cost.DogBones},
				{string(models.ResourceDog), cost.Dogs},
			}

			for _, c := range costs {
				if c.amount <= 0 {
					continue
				}
				res := tx.Model(&models.ResourceBag{}).
					Where("user_id = ? AND resource_key = ? AND amount >= ?", userID, c.key, c.amount).
					UpdateColumn("amount", gorm.Expr("amount - ?", c.amount))
				if res.Error != nil {
					return res.Error
				}
				if res.RowsAffected == 0 {
					return fmt.Errorf("%w: not enough %s", ErrInsufficientResources, c.key)
				}
			}

			now := time.Now().UTC()
			completesAt := now.Add(upgradeDuration)
			return tx.Model(&building).Updates(map[string]any{
				"is_constructing": true,
				"started_at":      now,
				"completes_at":    completesAt,
			}).Error
		})

		if txErr != nil {
			if errors.Is(txErr, ErrInsufficientResources) {
				http.Error(w, txErr.Error(), http.StatusUnprocessableEntity)
				return
			}
			http.Error(w, "failed to start building upgrade", http.StatusInternalServerError)
			return
		}

		building.IsConstructing = true

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(UpgradeBuildingResponse{
			Message:  "upgrade started",
			Building: building,
			CostPaid: cost,
		})
	})
}
