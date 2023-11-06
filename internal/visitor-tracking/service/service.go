package service

import (
	"fmt"
	"github.com/DarioKnezovic/analytics-service/internal/models"
	visitor_tracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"github.com/DarioKnezovic/analytics-service/internal/visitor-tracking/repository"
	"time"
)

type VisitorTrackingService struct {
	visitorTrackingRepository repository.VisitorTrackingRepository
}

const (
	DailyPeriod   = "daily"
	WeeklyPeriod  = "weekly"
	MonthlyPeriod = "monthly"
)

func NewVisitorTrackingService(visitorTrackingRepository repository.VisitorTrackingRepository) *VisitorTrackingService {
	return &VisitorTrackingService{
		visitorTrackingRepository: visitorTrackingRepository,
	}
}

func (s *VisitorTrackingService) RegisterVisitingUser(visitor models.VisitorTracking) error {
	visitor.Timestamp = time.Now()

	return s.visitorTrackingRepository.SaveVisitingUser(visitor)
}

func (s *VisitorTrackingService) CalculateAdBlockRate(campaignId string, startDate string, endDate string) (visitor_tracking.AdBlockRateResponse, error) {
	rate, err := s.visitorTrackingRepository.FetchAdblockRateForCampaign(campaignId, startDate, endDate)

	return rate, err
}

func (s *VisitorTrackingService) CalculateAdBlockRateHistory(params visitor_tracking.AdBlockRateHistoryParams) (visitor_tracking.HistoricalDataResponse, error) {
	// A function type that defines the signature of the repository methods
	type fetchFunc func(string, string, string) ([]visitor_tracking.HistoricalEntry, error)

	// A map to link periods to their respective repository functions
	periodToFetcher := map[string]fetchFunc{
		DailyPeriod:   s.visitorTrackingRepository.FetchHistoricalDailyAdBlockRate,
		WeeklyPeriod:  s.visitorTrackingRepository.FetchHistoricalWeeklyAdBlockRate,
		MonthlyPeriod: s.visitorTrackingRepository.FetchHistoricalMonthlyAdBlockRate,
	}

	// Obtain the appropriate fetcher function based on the period
	fetcher, ok := periodToFetcher[params.Period]
	if !ok {
		return visitor_tracking.HistoricalDataResponse{}, fmt.Errorf("invalid period: %v", params.Period)
	}

	// Use the fetcher to get the historical data
	response, err := fetcher(params.CampaignId, params.StartDate, params.EndDate)
	if err != nil {
		// Add more context to the error or handle it accordingly
		return visitor_tracking.HistoricalDataResponse{}, fmt.Errorf("failed to fetch historical data: %w", err)
	}

	// Depending on the actual logic, if printing is necessary, use a structured logger instead
	// log.Infof("Fetched response: %+v", response)

	// Prepare the response (Assuming it involves more than just wrapping the slice)
	return visitor_tracking.HistoricalDataResponse{Data: response}, nil
}
