package house

import "time"

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