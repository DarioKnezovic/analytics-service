package api

import (
	"github.com/DarioKnezovic/analytics-service/api/handlers"
	modal_ctr_tracking "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking"
	visitor_tracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, visitorTrackingService visitor_tracking.VisitorTrackingService, modalCtrTrackingService modal_ctr_tracking.ModalCtrTrackingService) {
	visitorTrackingHandler := &handlers.VisitorTrackingHandler{
		VisitorTrackingService: visitorTrackingService,
	}

	modalCtrTrackingHandler := &handlers.ModalCtrTrackingHandler{
		ModalCtrTrackingService: modalCtrTrackingService,
	}

	router.GET("/api/test", visitorTrackingHandler.TestEndpoint)
	router.POST("/api/analytics/visiting-users", visitorTrackingHandler.RegisterVisitingUser)
	router.GET("/api/analytics/adblock-rate", visitorTrackingHandler.FetchAdBlockRate)
	router.GET("/api/analytics/adblock-rate/history", visitorTrackingHandler.FetchAdblockRateHistory)

	router.POST("/api/analytics/modal-ctr-tracking", modalCtrTrackingHandler.RegisterModalCtrTracking)
	router.GET("/api/analytics/modal-ctr-tracking", modalCtrTrackingHandler.FetchModalCtrTrackingStatistics)
}
