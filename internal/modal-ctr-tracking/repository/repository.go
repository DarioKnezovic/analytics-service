package repository

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"gorm.io/gorm"
)

type ModalCtrTrackingRepository interface {
	StoreModalCtrTracking(data models.ModalCTRTracking) error
}

type modalCtrTrackingRepository struct {
	db *gorm.DB
}

func NewModalCtrTrackingRepository(db *gorm.DB) ModalCtrTrackingRepository {
	return &modalCtrTrackingRepository{
		db: db,
	}
}

func (r *modalCtrTrackingRepository) StoreModalCtrTracking(data models.ModalCTRTracking) error {
	return r.db.Create(&data).Error
}
