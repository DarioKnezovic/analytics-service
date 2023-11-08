package main

import (
	"fmt"
	"github.com/DarioKnezovic/analytics-service/api"
	"github.com/DarioKnezovic/analytics-service/config"
	mctRepo "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking/repository"
	mctService "github.com/DarioKnezovic/analytics-service/internal/modal-ctr-tracking/service"
	"github.com/DarioKnezovic/analytics-service/internal/visitor-tracking/repository"
	"github.com/DarioKnezovic/analytics-service/internal/visitor-tracking/service"
	"github.com/DarioKnezovic/analytics-service/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.PerformAutoMigrations(db)
	if err != nil {
		log.Fatalf("Failed to perform auto migrations: %v", err)
	}

	visitorTrackingRepo := repository.NewVisitorTrackingRepository(db)
	visitorTrackingService := service.NewVisitorTrackingService(visitorTrackingRepo)

	modalCtrTrackingRepo := mctRepo.NewModalCtrTrackingRepository(db)
	modalCtrTrackingService := mctService.NewModalTrackingCtrService(modalCtrTrackingRepo)

	router := gin.Default()
	api.RegisterRoutes(router, visitorTrackingService, modalCtrTrackingService)

	log.Printf("Server listening on port %s", cfg.APIPort)
	log.Fatal(router.Run(fmt.Sprintf(":%s", cfg.APIPort)))
}
