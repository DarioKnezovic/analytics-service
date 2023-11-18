package repository

import (
	modal_ctr_tracking "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking"
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"gorm.io/gorm"
)

type ModalCtrTrackingRepository interface {
	StoreModalCtrTracking(data models.ModalCTRTracking) error
	GetModalCtrTracking(params modal_ctr_tracking.FormattedModalCtrTrackingParams) ([]models.ModalCTRTracking, error)
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
	query := `INSERT INTO modal_ctr_tracking (session, additional_data, campaign_id, interaction_type, object_id) VALUES (?, ?, ?, ?, ?)`

	return r.db.Exec(query, data.Session, data.AdditionalData, data.CampaignID, data.InteractionType, data.ObjectID).Error
}

func (r *modalCtrTrackingRepository) GetModalCtrTracking(params modal_ctr_tracking.FormattedModalCtrTrackingParams) ([]models.ModalCTRTracking, error) {
	var trackings []models.ModalCTRTracking

	query := r.db.Where("campaign_id", params.CampaignId)

	if params.StartDate != nil {
		query = query.Where("timestamp >= ?", *params.StartDate)
	}
	if params.EndDate != nil {
		query = query.Where("timestamp <= ?", *params.EndDate)
	}
	if params.InteractionType != "" {
		query = query.Where("interaction_type = ?", params.InteractionType)
	}

	err := query.Find(&trackings).Error
	if err != nil {
		return nil, err
	}

	return trackings, nil
}
