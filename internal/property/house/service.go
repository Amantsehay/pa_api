package house

import (
    "time"
)

type House struct {
    ID          string
    Address     string
    Owner       string
    AreaSqM     float64
    AppraisedAt time.Time
    Value       float64
    NumBedrooms  int
    NumBathrooms int
    HasGarage    bool
    YearBuilt    int
}

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