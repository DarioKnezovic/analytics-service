package service

import (
	modal_ctr_tracking "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking"
	modalCtrTrackingRepository "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking/repository"
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"time"
)

type ModalCtrTrackingService struct {
	modalCtrTrackingRepository modalCtrTrackingRepository.ModalCtrTrackingRepository
}

func NewModalTrackingCtrService(modalTrackingRepository modalCtrTrackingRepository.ModalCtrTrackingRepository) *ModalCtrTrackingService {
	return &ModalCtrTrackingService{
		modalCtrTrackingRepository: modalTrackingRepository,
	}
}

func (s *ModalCtrTrackingService) RegisterNewModalCtrTracking(modalCtrTrackingData models.ModalCTRTracking) error {
	return s.modalCtrTrackingRepository.StoreModalCtrTracking(modalCtrTrackingData)
}

func (s *ModalCtrTrackingService) GetModalCtrTrackingData(params modal_ctr_tracking.ModalCtrTrackingParams) ([]models.ModalCTRTracking, error) {
	var (
		startOfDay time.Time
		endOfDay   time.Time
	)
	if params.StartDate != "" {
		startDate, err := time.Parse("2006-01-02", params.StartDate)
		if err != nil {
			return nil, err
		}
		startOfDay = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 00, 00, 00, 00, startDate.Location())
	}
	if params.EndDate != "" {
		endDate, err := time.Parse("2006-01-02", params.EndDate)
		if err != nil {
			return nil, err
		}
		endOfDay = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 23, 59, 59, 0, endDate.Location())
	}

	formattedParams := modal_ctr_tracking.FormattedModalCtrTrackingParams{
		CampaignId:      params.CampaignId,
		StartDate:       &startOfDay,
		EndDate:         &endOfDay,
		InteractionType: params.InteractionType,
	}

	return s.modalCtrTrackingRepository.GetModalCtrTracking(formattedParams)
}
