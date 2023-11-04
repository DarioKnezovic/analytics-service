package visitor_tracking

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
)

type VisitorTrackingService interface {
	RegisterVisitingUser(visitor models.VisitorTracking) error
}
