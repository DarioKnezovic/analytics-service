package service

import (
	modalCtrTrackingRepository "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking/repository"
	"github.com/DarioKnezovic/analytics-service/internal/models"
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
