package handlers

import (
	modalCtrTracking "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking"
	"github.com/DarioKnezovic/analytics-service/internal/models"
	"github.com/DarioKnezovic/analytics-service/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ModalCtrTrackingHandler struct {
	ModalCtrTrackingService modalCtrTracking.ModalCtrTrackingService
}

// isValidInteractionType checks if the provided interaction_type is valid.
func isValidInteractionType(interactionType string) bool {
	// Define a list of allowed interaction types
	allowedTypes := []string{"click", "hover", "submit", "scroll", "close", "input", "view", "download", "share", "expand", "collapse", "play", "pause", "navigation"}

	// Check if the provided interaction_type is in the list of allowed types
	for _, allowedType := range allowedTypes {
		if interactionType == allowedType {
			return true
		}
	}

	// If not found in the list, it's invalid
	return false
}

func (h *ModalCtrTrackingHandler) RegisterModalCtrTracking(c *gin.Context) {
	var newModalCtrTracking models.ModalCTRTracking

	if err := c.BindJSON(&newModalCtrTracking); err != nil {
		util.SendJSONResponse(c, http.StatusBadRequest, nil)
		return
	}

	if !isValidInteractionType(newModalCtrTracking.InteractionType) {
		util.SendJSONResponse(c, http.StatusBadRequest, gin.H{"error": "Invalid interaction_type value"})
		return
	}

	err := h.ModalCtrTrackingService.RegisterNewModalCtrTracking(newModalCtrTracking)
	if err != nil {
		log.Printf("Error during saving new visitor: %e", err)
		util.SendJSONResponse(c, http.StatusInternalServerError, nil)
		return
	}

	util.SendJSONResponse(c, http.StatusOK, "")
}

func (h *ModalCtrTrackingHandler) FetchModalCtrTrackingStatistics(c *gin.Context) {
	campaignId, interactionType, startDate, endDate := c.Query("campaign_id"), c.Query("interaction_type"), c.Query("start_date"), c.Query("end_date")

	if campaignId == "" {
		util.SendJSONResponse(c, http.StatusBadRequest, gin.H{"error": "campaign_id is required"})
		return
	}

	params := modalCtrTracking.ModalCtrTrackingParams{
		CampaignId:      campaignId,
		InteractionType: interactionType,
		StartDate:       startDate,
		EndDate:         endDate,
	}

	response, err := h.ModalCtrTrackingService.GetModalCtrTrackingData(params)
	if err != nil {
		log.Printf("Something went wrong during fetching modal ctr tracking data: %e", err)
		util.SendJSONResponse(c, http.StatusInternalServerError, gin.H{"error": "Failed to fetch modal ctr tracking data"})
		return
	}

	util.SendJSONResponse(c, http.StatusOK, response)
}
