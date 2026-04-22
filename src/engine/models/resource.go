package models

import "gorm.io/gorm"

type ResourceKey string

const (
	ResourceDogCoin ResourceKey = "dogcoin"
	ResourceDogBone ResourceKey = "dogbones"
	ResourceDog     ResourceKey = "dogs"
)

type ResourceBalance struct {
	gorm.Model
	UserID      uint   `gorm:"index;not null"`
	ResourceKey string `gorm:"uniqueIndex:idx_user_resource_key;not null"`
	Amount      int64  `gorm:"not null;default:0"`
	Capacity    int64  `gorm:"not null;default:0"`
}

func DefaultResourceBalances(userID uint) []ResourceBalance {
	return []ResourceBalance{
		{UserID: userID, ResourceKey: string(ResourceDogCoin), Amount: 0, Capacity: 0},
		{UserID: userID, ResourceKey: string(ResourceDogBone), Amount: 0, Capacity: 0},
		{UserID: userID, ResourceKey: string(ResourceDog), Amount: 0, Capacity: 0},
	}
}
