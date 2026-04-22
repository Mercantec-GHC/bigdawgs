package handlers

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

func RunTickLoop(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Println("Tilføj 100 dog points til alle brugere")
	}
}
