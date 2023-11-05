package main

import (
	"fmt"
	"github.com/DarioKnezovic/analytics-service/api"
	"github.com/DarioKnezovic/analytics-service/config"
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

	router := gin.Default()
	api.RegisterRoutes(router, visitorTrackingService)

	log.Printf("Server listening on port %s", cfg.APIPort)
	log.Fatal(router.Run(fmt.Sprintf(":%s", cfg.APIPort)))
}
