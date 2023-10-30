package repository

import "gorm.io/gorm"

type AnalyticsRepository interface {
	GetAnalytics()
}

type analyticsRepository struct {
	db *gorm.DB
}

func (r analyticsRepository) GetAnalytics() {
	//TODO implement me
	panic("implement me")
}

func NewAnalyticsRepository(db *gorm.DB) AnalyticsRepository {
	return &analyticsRepository{
		db: db,
	}
}
