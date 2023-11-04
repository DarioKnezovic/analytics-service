package repository

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"gorm.io/gorm"
	"log"
)

type VisitorTrackingRepository interface {
	SaveVisitingUser(visitor models.VisitorTracking) error
}

type visitorTrackingRepository struct {
	db *gorm.DB
}

func NewVisitorTrackingRepository(db *gorm.DB) VisitorTrackingRepository {
	return &visitorTrackingRepository{
		db: db,
	}
}

func (r *visitorTrackingRepository) SaveVisitingUser(visitor models.VisitorTracking) error {
	log.Println("Visitor is ready for storing in the database")

	return r.db.Create(&visitor).Error
}
