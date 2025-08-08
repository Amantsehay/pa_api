package house

import (
	"pa_api/internal/db"
	"time"
)

// HouseService provides appraisal logic for House properties
type HouseService struct{}

func (s *HouseService) Appraise(house House) House {
    value := 50000.0 + float64(house.NumBedrooms)*10000 + float64(house.NumBathrooms)*5000
    if house.HasGarage {
        value += 20000
    }
    house.Value = value
    house.AppraisedAt = time.Now()
    return house
}

func (s *HouseService) Save(house House) error {
    return db.DB.Create(&house).Error
}