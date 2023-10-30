package main

import (
	"fmt"
	"github.com/DarioKnezovic/analytics-service/api"
	"github.com/DarioKnezovic/analytics-service/repository"
	"github.com/DarioKnezovic/analytics-service/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func main() {
	db := &gorm.DB{}

	analyticsRepo := repository.NewAnalyticsRepository(db)
	analyticsService := service.NewAnalyticsService(analyticsRepo)

	router := gin.Default()
	api.RegisterRoutes(router, analyticsService)

	log.Printf("Server listening on port 5555")
	log.Fatal(router.Run(fmt.Sprintf(":5555")))
}
