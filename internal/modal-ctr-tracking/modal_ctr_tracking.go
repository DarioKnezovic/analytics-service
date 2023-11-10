package modal_ctr_tracking

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"time"
)

type ModalCtrTrackingService interface {
	RegisterNewModalCtrTracking(modalCtrTrackingData models.ModalCTRTracking) error
	GetModalCtrTrackingData(params ModalCtrTrackingParams) ([]models.ModalCTRTracking, error)
}

type ModalCtrTrackingParams struct {
	InteractionType string `json:"interaction_type"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	CampaignId      string `json:"campaign_id"`
}

type FormattedModalCtrTrackingParams struct {
	InteractionType string     `json:"interaction_type"`
	StartDate       *time.Time `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
	CampaignId      string     `json:"campaign_id"`
}
