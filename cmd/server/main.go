package main

import (
	"pa_api/internal/config"
	"pa_api/internal/db"
	"pa_api/internal/property/house"

	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()
    db.Connect()
    db.DB.AutoMigrate(&house.House{})

    houseService := &house.HouseService{}
    houseHandler := &house.HouseHandler{Service: houseService}

    r := gin.Default()
    r.POST("/house", houseHandler.CreateHouseHandler)
    // Optionally, wrap AppraiseHandler for Gin if needed
    r.Run(":8080")
}
