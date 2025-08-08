package house

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type House struct {
    ID          string    `json:"id" gorm:"primaryKey"`
    Address     string    `json:"address"`
    Owner       string    `json:"owner"`
    AreaSqM     float64   `json:"area_sqm"`
    AppraisedAt time.Time `json:"appraised_at"`
    Value       float64   `json:"value"`
    NumBedrooms  int  `json:"num_bedrooms"`
    NumBathrooms int  `json:"num_bathrooms"`
    HasGarage    bool `json:"has_garage"`
    YearBuilt    int  `json:"year_built"`
}

type HouseService interface {
    Appraise(house House) House
    Save(house House) error
}

type HouseHandler struct {
    Service HouseService
}

func (h *HouseHandler) AppraiseHandler(w http.ResponseWriter, r *http.Request) {
    var house House
    if err := json.NewDecoder(r.Body).Decode(&house); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    appraised := h.Service.Appraise(house)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(appraised)
}

func (h *HouseHandler) CreateHouseHandler(c *gin.Context) {
    var house House
    if err := c.ShouldBindJSON(&house); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.Service.Save(house); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save house"})
        return
    }
    c.JSON(http.StatusCreated, house)
}