package modal_ctr_tracking

import "github.com/DarioKnezovic/analytics-service/internal/models"

type ModalCtrTrackingService interface {
	RegisterNewModalCtrTracking(modalCtrTrackingData models.ModalCTRTracking) error
}
