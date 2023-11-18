package repository

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	visitor_tracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"gorm.io/gorm"
	"log"
)

type VisitorTrackingRepository interface {
	SaveVisitingUser(visitor models.VisitorTracking) error
	FetchAdblockRateForCampaign(campaignId string, startDate string, endDate string) (visitor_tracking.AdBlockRateResponse, error)
	FetchHistoricalDailyAdBlockRate(campaignId string, startDate string, endDate string) ([]visitor_tracking.HistoricalEntry, error)
	FetchHistoricalWeeklyAdBlockRate(campaignId string, startDate string, endDate string) ([]visitor_tracking.HistoricalEntry, error)
	FetchHistoricalMonthlyAdBlockRate(campaignId string, startDate string, endDate string) ([]visitor_tracking.HistoricalEntry, error)
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
	query := `INSERT INTO visitor_tracking (timestamp, adblock_user, campaign_id) VALUES (?, ?, ?)`

	return r.db.Exec(query, visitor.Timestamp, visitor.AdblockUser, visitor.CampaignID).Error
}

func (r *visitorTrackingRepository) FetchAdblockRateForCampaign(campaignId string, startDate string, endDate string) (visitor_tracking.AdBlockRateResponse, error) {
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

func (r *visitorTrackingRepository) FetchHistoricalDailyAdBlockRate(campaignId string, startDate string, endDate string) ([]visitor_tracking.HistoricalEntry, error) {
	var rateHistory []visitor_tracking.HistoricalEntry

	query := `
		SELECT
			DATE(timestamp) AS date,
			ROUND(SUM(CASE WHEN adblock_user = TRUE THEN 1 ELSE 0 END) * 100.0 / COUNT(*), 2) AS adblock_rate
		FROM
			visitor_tracking
		WHERE
			timestamp >= ? AND timestamp <= ?
			AND campaign_id = ?
		GROUP BY
			DATE(timestamp)
		ORDER BY
			date;`

	err := r.db.Raw(query, startDate, endDate, campaignId).Scan(&rateHistory).Error
	return rateHistory, err
}

func (r *visitorTrackingRepository) FetchHistoricalWeeklyAdBlockRate(campaignId string, startDate string, endDate string) ([]visitor_tracking.HistoricalEntry, error) {
	var rateHistory []visitor_tracking.HistoricalEntry

	query := `
		SELECT
			DATE_TRUNC('week', timestamp) AS date,
			ROUND(SUM(CASE WHEN adblock_user = TRUE THEN 1 ELSE 0 END) * 100.0 / COUNT(*), 2) AS adblock_rate
		FROM
			visitor_tracking
		WHERE
			timestamp >= ? AND timestamp <= ?
			AND campaign_id = ?
		GROUP BY
			date
		ORDER BY
			date;`

	err := r.db.Raw(query, startDate, endDate, campaignId).Scan(&rateHistory).Error
	return rateHistory, err
}

func (r *visitorTrackingRepository) FetchHistoricalMonthlyAdBlockRate(campaignId string, startDate string, endDate string) ([]visitor_tracking.HistoricalEntry, error) {
	var rateHistory []visitor_tracking.HistoricalEntry

	query := `
		SELECT
			DATE_TRUNC('month', timestamp) AS date,
			ROUND(SUM(CASE WHEN adblock_user = TRUE THEN 1 ELSE 0 END) * 100.0 / COUNT(*), 2) AS adblock_rate
		FROM
			visitor_tracking
		WHERE
			timestamp >= ? AND timestamp <= ?
			AND campaign_id = ?
		GROUP BY
			date
		ORDER BY
			date;`

	err := r.db.Raw(query, startDate, endDate, campaignId).Scan(&rateHistory).Error
	return rateHistory, err
}
