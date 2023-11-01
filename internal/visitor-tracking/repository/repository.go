package repository

import "gorm.io/gorm"

type VisitorTrackingRepository interface {
}

type visitorTrackingRepository struct {
	db *gorm.DB
}

func NewVisitorTrackingRepository(db *gorm.DB) VisitorTrackingRepository {
	return &visitorTrackingRepository{
		db: db,
	}
}
