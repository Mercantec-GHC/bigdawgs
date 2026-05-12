package models

import (
	"math"
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

const (
	ProductionMultiplier = 1.15
	CostMultiplier       = 1.25
)

type Production struct {
	DogCoins int64 `json:"dog_coins"`
	DogBones int64 `json:"dog_bones"`
	Dogs     int64 `json:"dogs"`
}

type BuildingDefinition struct {
	Key            BuildingKey
	DisplayName    string
	BaseProduction Production
	BaseCost       Production
}

var BuildingDefinitions = map[BuildingKey]BuildingDefinition{
	MeatFactory: {
		Key:            MeatFactory,
		DisplayName:    "Meat Factory",
		BaseProduction: Production{DogBones: 10},
		BaseCost:       Production{DogBones: 100, DogCoins: 50, Dogs: 5},
	},
	DogCoinDen: {
		Key:            DogCoinDen,
		DisplayName:    "Dog Coin Den",
		BaseProduction: Production{DogCoins: 5},
		BaseCost:       Production{DogBones: 200, DogCoins: 100, Dogs: 10},
	},
	Doghouse: {
		Key:            Doghouse,
		DisplayName:    "The Doghouse",
		BaseProduction: Production{Dogs: 2},
		BaseCost:       Production{DogBones: 300, DogCoins: 150, Dogs: 20},
	},
	DogKennel: {
		Key:         DogKennel,
		DisplayName: "The DogKennel",
		BaseCost:    Production{DogBones: 150, DogCoins: 75},
	},
	Market: {
		Key:         Market,
		DisplayName: "The Market",
		BaseCost:    Production{DogBones: 500, DogCoins: 250, Dogs: 25},
	},
}

type Building struct {
	gorm.Model
	UserID         uint   `gorm:"uniqueIndex:idx_user_building_key;not null;index"`
	Key            string `gorm:"uniqueIndex:idx_user_building_key;not null"`
	Level          int    `gorm:"not null;default:0"`
	IsConstructing bool   `gorm:"not null;default:false"`
	StartedAt      *time.Time
	CompletesAt    *time.Time
}

func IsValidBuildingKey(key string) bool {
	_, ok := BuildingDefinitions[BuildingKey(key)]
	return ok
}

func (b Building) ProductionPerTick() Production {
	definition, ok := BuildingDefinitions[BuildingKey(b.Key)]
	if !ok || b.Level < 1 {
		return Production{}
	}

	multiplier := math.Pow(ProductionMultiplier, float64(b.Level-1))
	return Production{
		DogCoins: int64(math.Round(float64(definition.BaseProduction.DogCoins) * multiplier)),
		DogBones: int64(math.Round(float64(definition.BaseProduction.DogBones) * multiplier)),
		Dogs:     int64(math.Round(float64(definition.BaseProduction.Dogs) * multiplier)),
	}
}

func (b Building) UpgradeCost() Production {
	definition, ok := BuildingDefinitions[BuildingKey(b.Key)]
	if !ok {
		return Production{}
	}

	multiplier := math.Pow(CostMultiplier, float64(b.Level))
	return Production{
		DogCoins: int64(math.Ceil(float64(definition.BaseCost.DogCoins) * multiplier)),
		DogBones: int64(math.Ceil(float64(definition.BaseCost.DogBones) * multiplier)),
		Dogs:     int64(math.Ceil(float64(definition.BaseCost.Dogs) * multiplier)),
	}
}
