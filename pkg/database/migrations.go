package database

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"gorm.io/gorm"
	"log"
)

func PerformAutoMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.VisitorTracking{})

	if err != nil {
		return err
	}

	log.Print("Auto migrations have been successfully finished")

	return nil
}
