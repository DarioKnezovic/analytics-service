package api

import (
	"github.com/DarioKnezovic/analytics-service/api/handlers"
	visitor_tracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, visitorTrackingService visitor_tracking.VisitorTrackingService) {
	visitorTrackingHandler := &handlers.VisitorTrackingHandler{
		VisitorTrackingService: visitorTrackingService,
	}

	router.GET("/api/test", visitorTrackingHandler.TestEndpoint)
	router.POST("/api/analytics/visiting-users", visitorTrackingHandler.RegisterVisitingUser)
	router.GET("/api/analytics/adblock-rate", visitorTrackingHandler.FetchAdBlockRate)
	router.GET("/api/analytics/adblock-rate/history", visitorTrackingHandler.FetchAdblockRateHistory)
}
