package models

import (
	"time"

	"gorm.io/gorm"
)

type BuildingKey string

const (
	BuildingMeatFactory BuildingKey = "meat_factory"
	BuildingDogCoinDen  BuildingKey = "dog_coin_den"
	BuildingDoghouse    BuildingKey = "the_doghouse"
)

type Production struct {
	DogCoins int64
	DogBones int64
	Dogs     int64
}

type BuildingDefinition struct {
	Key            BuildingKey
	DisplayName    string
	BaseProduction Production
}

var BuildingDefinitions = map[BuildingKey]BuildingDefinition{
	BuildingMeatFactory: {
		Key:         BuildingMeatFactory,
		DisplayName: "Meat Factory",
		BaseProduction: Production{
			DogBones: 10,
		},
	},
	BuildingDogCoinDen: {
		Key:         BuildingDogCoinDen,
		DisplayName: "Dog Coin Den",
		BaseProduction: Production{
			DogCoins: 5,
		},
	},
	BuildingDoghouse: {
		Key:         BuildingDoghouse,
		DisplayName: "The Doghouse",
		BaseProduction: Production{
			Dogs: 2,
		},
	},
}

type Building struct {
	gorm.Model
	UserID         uint   `gorm:"uniqueIndex:idx_user_building_key;not null"`
	Key            string `gorm:"uniqueIndex:idx_user_building_key;not null"`
	Level          int    `gorm:"not null;default:1"`
	Count          int    `gorm:"not null;default:1"`
	IsConstructing bool   `gorm:"not null;default:false"`
	StartedAt      *time.Time
	CompletesAt    *time.Time
}

func (b Building) NormalizedLevel() int64 {
	if b.Level < 1 {
		return 1
	}

	return int64(b.Level)
}

func (b Building) NormalizedCount() int64 {
	if b.Count < 1 {
		return 1
	}

	return int64(b.Count)
}

func (b Building) ProductionPerTick() Production {
	definition, ok := BuildingDefinitions[BuildingKey(b.Key)]
	if !ok {
		return Production{}
	}

	multiplier := b.NormalizedLevel() * b.NormalizedCount()

	return Production{
		DogCoins: definition.BaseProduction.DogCoins * multiplier,
		DogBones: definition.BaseProduction.DogBones * multiplier,
		Dogs:     definition.BaseProduction.Dogs * multiplier,
	}
}
