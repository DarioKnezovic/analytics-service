package api

import (
	"github.com/DarioKnezovic/analytics-service/api/handlers"
	"github.com/DarioKnezovic/analytics-service/internal/analytics"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, analyticsService analytics.AnalyticsService) {
	analyticsHandler := &handlers.AnalyticsHandler{
		AnalyticsService: analyticsService,
	}

	router.GET("/api/test", analyticsHandler.TestHandler)
}
