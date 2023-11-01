package handlers

import (
	visitorTracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"github.com/DarioKnezovic/analytics-service/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VisitorTrackingHandler struct {
	VisitorTrackingService visitorTracking.VisitorTrackingService
}

func (h *VisitorTrackingHandler) TestEndpoint(c *gin.Context) {
	util.SendJSONResponse(c, http.StatusOK, map[string]string{"test": "testera"})
}
