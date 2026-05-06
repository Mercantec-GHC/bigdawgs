package transactions

import (
	"errors"
	"fmt"

	"bigdawgs/models"

	"gorm.io/gorm"
)

var ErrInsufficientResources = errors.New("insufficient resources")

type TradeResult struct {
	Spent    models.ResourceBag
	Received models.ResourceBag
}

func Trade(db *gorm.DB, userID uint, spendKey string, spendAmount int64, receiveKey string, receiveAmount int64) (TradeResult, error) {
	var result TradeResult

	err := db.Transaction(func(tx *gorm.DB) error {
		var spendBag models.ResourceBag
		if err := tx.Where("user_id = ? AND resource_key = ?", userID, spendKey).First(&spendBag).Error; err != nil {
			return fmt.Errorf("spend resource bag not found: %w", err)
		}

		if spendBag.Amount < spendAmount {
			return fmt.Errorf("%w: have %d %s, need %d", ErrInsufficientResources, spendBag.Amount, spendKey, spendAmount)
		}

		spendBag.Amount -= spendAmount
		if err := tx.Save(&spendBag).Error; err != nil {
			return fmt.Errorf("failed to deduct spend resource: %w", err)
		}

		var receiveBag models.ResourceBag
		if err := tx.Where("user_id = ? AND resource_key = ?", userID, receiveKey).First(&receiveBag).Error; err != nil {
			return fmt.Errorf("receive resource bag not found: %w", err)
		}

		receiveBag.Amount += receiveAmount
		if err := tx.Save(&receiveBag).Error; err != nil {
			return fmt.Errorf("failed to credit receive resource: %w", err)
		}

		result = TradeResult{Spent: spendBag, Received: receiveBag}
		return nil
	})

	return result, err
}
