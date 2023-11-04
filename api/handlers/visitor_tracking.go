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
