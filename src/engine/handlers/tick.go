package handlers

import (
	"bigdawgs/models"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func RunTickLoop(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		if err := processTick(db); err != nil {
			fmt.Println("tick error:", err)
		}
	}
}

func processTick(db *gorm.DB) error {
	now := time.Now().UTC()

	finishedBuildings := []models.Building{}
	db.Where("is_constructing = ? AND completes_at <= ?", true, now).Find(&finishedBuildings)
	for _, b := range finishedBuildings {
		db.Model(&b).Updates(map[string]any{
			"level":           b.Level + 1,
			"is_constructing": false,
			"started_at":      nil,
			"completes_at":    nil,
		})
	}

	var userIDs []uint
	db.Model(&models.Building{}).Distinct("user_id").Pluck("user_id", &userIDs)

	for _, userID := range userIDs {
		var userBuildings []models.Building
		db.Where("user_id = ?", userID).Find(&userBuildings)

		var total models.Production
		for _, building := range userBuildings {
			production := building.ProductionPerTick()
			total.DogCoins += production.DogCoins
			total.DogBones += production.DogBones
			total.Dogs += production.Dogs
		}

		adjustments := map[models.ResourceKey]int64{
			models.ResourceDogCoin: total.DogCoins,
			models.ResourceDogBone: total.DogBones,
			models.ResourceDog:     total.Dogs,
		}
		for key, amount := range adjustments {
			if amount == 0 {
				continue
			}
			db.Model(&models.ResourceBag{}).
				Where("user_id = ? AND resource_key = ?", userID, string(key)).
				UpdateColumn("amount", gorm.Expr("amount + ?", amount))
		}
	}
	return nil
}
