package models

import (
	"time"

	"gorm.io/gorm"
)

type BuildingKey string

const (
	MeatFactory BuildingKey = "meat_factory"
	DogCoinDen  BuildingKey = "dog_coin_den"
	Doghouse    BuildingKey = "the_doghouse"
	DogKennel   BuildingKey = "the_dog_kennel"
	Market      BuildingKey = "market"
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
	UpgradeCost    Production
}

var BuildingDefinitions = map[BuildingKey]BuildingDefinition{
	MeatFactory: {
		Key:         MeatFactory,
		DisplayName: "Meat Factory",
		BaseProduction: Production{
			DogBones: 10,
		},
	},
	DogCoinDen: {
		Key:         DogCoinDen,
		DisplayName: "Dog Coin Den",
		BaseProduction: Production{
			DogCoins: 5,
		},
	},
	Doghouse: {
		Key:         Doghouse,
		DisplayName: "The Doghouse",
		BaseProduction: Production{
			Dogs: 2,
		},
	},
	DogKennel: {
		Key:         DogKennel,
		DisplayName: "The DogKennel",
	},
	Market: {
		Key:         Market,
		DisplayName: "The Market",
	},
}

type Building struct {
	gorm.Model
	UserID              uint   `gorm:"uniqueIndex:idx_user_building_key;not null;index"`
	Key                 string `gorm:"uniqueIndex:idx_user_building_key;not null"`
	Level               int    `gorm:"not null;default:0"`
	UpgradeCostDogCoins int64  `gorm:"not null;default:0"`
	UpgradeCostDogBones int64  `gorm:"not null;default:0"`
	UpgradeCostDogs     int64  `gorm:"not null;default:0"`
	IsConstructing      bool   `gorm:"not null;default:false"`
	StartedAt           *time.Time
	CompletesAt         *time.Time
}

func IsValidBuildingKey(key string) bool {
	_, ok := BuildingDefinitions[BuildingKey(key)]
	return ok
}

func (b Building) NormalizedLevel() int64 {
	if b.Level < 1 {
		return 1
	}

	return int64(b.Level)
}

func (b Building) ProductionPerTick() Production {
	definition, ok := BuildingDefinitions[BuildingKey(b.Key)]
	if !ok {
		return Production{}
	}

	multiplier := b.NormalizedLevel()

	return Production{
		DogCoins: definition.BaseProduction.DogCoins * multiplier,
		DogBones: definition.BaseProduction.DogBones * multiplier,
		Dogs:     definition.BaseProduction.Dogs * multiplier,
	}
}
