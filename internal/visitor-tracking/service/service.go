package service

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
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
