package service

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	visitor_tracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"github.com/DarioKnezovic/analytics-service/internal/visitor-tracking/repository"
	"time"
)

type VisitorTrackingService struct {
	visitorTrackingRepository repository.VisitorTrackingRepository
}

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
	rate, err := s.visitorTrackingRepository.FetchAllVisitorsForCampaign(campaignId, startDate, endDate)

	return rate, err
}
