package repository

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	visitor_tracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"gorm.io/gorm"
	"log"
)

type VisitorTrackingRepository interface {
	SaveVisitingUser(visitor models.VisitorTracking) error
	FetchAllVisitorsForCampaign(campaignId string, startDate string, endDate string) (visitor_tracking.AdBlockRateResponse, error)
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

func (r *visitorTrackingRepository) FetchAllVisitorsForCampaign(campaignId string, startDate string, endDate string) (visitor_tracking.AdBlockRateResponse, error) {
	var adBlockRate visitor_tracking.AdBlockRateResponse

	query := `
        SELECT
            COUNT(CASE WHEN adblock_user = true THEN 1 ELSE NULL END) AS number_of_users_with_adblock,
            COUNT(CASE WHEN adblock_user = false THEN 1 ELSE NULL END) AS number_of_users_without_adblock,
            (COUNT(CASE WHEN adblock_user = true THEN 1 ELSE NULL END)::DECIMAL / COUNT(*)) * 100 AS adblock_rate
        FROM visitor_tracking
        WHERE campaign_id = ?
    `

	args := []interface{}{campaignId}

	if startDate != "" {
		query += " AND timestamp >= ? "
		args = append(args, startDate)
	}

	if endDate != "" {
		query += " AND timestamp <= ?"
		args = append(args, endDate)
	}

	err := r.db.Raw(query, args...).Scan(&adBlockRate).Error
	return adBlockRate, err
}
