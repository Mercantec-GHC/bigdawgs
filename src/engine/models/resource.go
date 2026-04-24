package models

import "gorm.io/gorm"

type ResourceKey string

const (
	ResourceDogCoin ResourceKey = "dogcoin"
	ResourceDogBone ResourceKey = "dogbones"
	ResourceDog     ResourceKey = "dogs"
)

type ResourceBag struct {
	gorm.Model
	User_id     uint   `gorm:"uniqueIndex:idx_user_resource_key;not null;index"`
	ResourceKey string `gorm:"uniqueIndex:idx_user_resource_key;not null"`
	Amount      int64  `gorm:"not null;default:0"`
}

func DefaultResourceBalances(userID uint) []ResourceBag {
	return []ResourceBag{
		{User_id: userID, ResourceKey: string(ResourceDogCoin), Amount: 0},
		{User_id: userID, ResourceKey: string(ResourceDogBone), Amount: 0},
		{User_id: userID, ResourceKey: string(ResourceDog), Amount: 0},
	}
}
