package service

import "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking/repository"

type VisitorTrackingService struct {
	visitorTrackingRepository repository.VisitorTrackingRepository
}

func NewVisitorTrackingService(visitorTrackingRepository repository.VisitorTrackingRepository) *VisitorTrackingService {
	return &VisitorTrackingService{
		visitorTrackingRepository: visitorTrackingRepository,
	}
}
