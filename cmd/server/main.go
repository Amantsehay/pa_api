package main

import (
    "net/http"
    "pa_api/internal/property/house"
)

func main() {
    houseService := &house.HouseService{}
    houseHandler := &house.HouseHandler{Service: houseService}

    http.HandleFunc("/appraise/house", houseHandler.AppraiseHandler)
    http.ListenAndServe(":8080", nil)
}
