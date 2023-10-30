package handlers

import (
	"github.com/DarioKnezovic/analytics-service/internal/analytics"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnalyticsHandler struct {
	AnalyticsService analytics.AnalyticsService
}

func (h *AnalyticsHandler) TestHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, map[string]string{
		"test": "testing",
	})
}
