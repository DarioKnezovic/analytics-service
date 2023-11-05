package handlers

import (
	"github.com/DarioKnezovic/analytics-service/internal/models"
	visitorTracking "github.com/DarioKnezovic/analytics-service/internal/visitor-tracking"
	"github.com/DarioKnezovic/analytics-service/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type VisitorTrackingHandler struct {
	VisitorTrackingService visitorTracking.VisitorTrackingService
}

func (h *VisitorTrackingHandler) TestEndpoint(c *gin.Context) {
	util.SendJSONResponse(c, http.StatusOK, map[string]string{"test": "testera"})
}

func (h *VisitorTrackingHandler) RegisterVisitingUser(c *gin.Context) {
	var newVisitor models.VisitorTracking

	log.Println("Received request for storing visiting user")

	if err := c.BindJSON(&newVisitor); err != nil {
		util.SendJSONResponse(c, http.StatusBadRequest, nil)
		return
	}

	err := h.VisitorTrackingService.RegisterVisitingUser(newVisitor)
	if err != nil {
		log.Printf("Error during saving new visitor: %e", err)
		util.SendJSONResponse(c, http.StatusInternalServerError, nil)
		return
	}
}

func (h *VisitorTrackingHandler) FetchAdBlockRate(c *gin.Context) {
	campaignId := c.Query("campaign_id")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if campaignId == "" {
		util.SendJSONResponse(c, http.StatusBadRequest, map[string]string{"error": "Parameter campaign_id is missing"})
		return
	}

	response, err := h.VisitorTrackingService.CalculateAdBlockRate(campaignId, startDate, endDate)
	if err != nil {
		log.Printf("Error during calculating adblock rate: %e", err)
		util.SendJSONResponse(c, http.StatusInternalServerError, nil)
		return
	}
	util.SendJSONResponse(c, http.StatusOK, response)
}
