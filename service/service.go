package service

import "github.com/DarioKnezovic/analytics-service/repository"

type AnalyticsService struct {
	analyticsRepository repository.AnalyticsRepository
}

func NewAnalyticsService(analyticsRepository repository.AnalyticsRepository) *AnalyticsService {
	return &AnalyticsService{
		analyticsRepository: analyticsRepository,
	}
}
