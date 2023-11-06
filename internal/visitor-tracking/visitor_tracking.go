package visitor_tracking

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
)

type AdBlockRateResponse struct {
	NumberOfUsersWithAdBlock    int     `gorm:"column:number_of_users_with_adblock" json:"number_of_users_with_adblock"`
	NumberOfUsersWithoutAdBlock int     `gorm:"column:number_of_users_without_adblock" json:"number_of_users_without_adblock"`
	AdBlockRate                 float64 `gorm:"column:adblock_rate" json:"adblock_rate"`
}

type HistoricalDataResponse struct {
	Period    string            `json:"period"`
	StartDate string            `json:"start_date"`
	EndDate   string            `json:"end_date"`
	Data      []HistoricalEntry `json:"data"`
}

type HistoricalEntry struct {
	Date        string  `gorm:"column:date" json:"date"`
	AdBlockRate float64 `gorm:"column:adblock_rate" json:"adblock_rate"`
}

type AdBlockRateHistoryParams struct {
	Period     string `json:"period"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	CampaignId string `json:"campaign_id"`
}

type VisitorTrackingService interface {
	RegisterVisitingUser(visitor models.VisitorTracking) error
	CalculateAdBlockRate(campaignId string, startDate string, endDate string) (AdBlockRateResponse, error)
	CalculateAdBlockRateHistory(params AdBlockRateHistoryParams) (HistoricalDataResponse, error)
}
