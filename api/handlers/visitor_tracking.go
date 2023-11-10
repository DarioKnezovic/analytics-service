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

func (h *VisitorTrackingHandler) FetchAdblockRateHistory(c *gin.Context) {
	campaignId, period, startDate, endDate := c.Query("campaign_id"), c.Query("period"), c.Query("start_date"), c.Query("end_date")

	// Validate required parameters.
	if campaignId == "" {
		util.SendJSONResponse(c, http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}
	if period == "" {
		util.SendJSONResponse(c, http.StatusBadRequest, gin.H{"error": "period is required"})
		return
	}
	if startDate == "" {
		util.SendJSONResponse(c, http.StatusBadRequest, gin.H{"error": "start date is required"})
		return
	}
	if endDate == "" {
		util.SendJSONResponse(c, http.StatusBadRequest, gin.H{"error": "end date is required"})
		return
	}

	params := visitorTracking.AdBlockRateHistoryParams{
		CampaignId: campaignId,
		StartDate:  startDate,
		EndDate:    endDate,
		Period:     period,
	}

	response, err := h.VisitorTrackingService.CalculateAdBlockRateHistory(params)
	if err != nil {
		log.Printf("Something went wrong during fetching history: %e", err)
		util.SendJSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch adblock rate history"})
		return
	}

	util.SendJSONResponse(c, http.StatusOK, response)
}
