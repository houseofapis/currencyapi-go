package main

import (
	"fmt"
	"encoding/json"
	"github.com/houseofapis/currencyapi-go"
)


type JsonResponse struct {
    Valid     bool    `json:"valid"`
    StartDate string  `json:"start_date"`
    EndDate   string  `json:"end_date"`
    TimeframeRates     map[string]map[string]float64   `json:"rates"`
    Error     *struct {
        Code    int    `json:"code"`
        Message string `json:"message"`
    } `json:"error"`
}

func main() {
    client := currencyapi.Client("YOUR_API_KEY")

    timeframeParams := map[string]string{
        "output": "JSON",
        "base":   "GBP",
        "start_date": "2019-01-01",
        "end_date": "2019-01-02",
    }

    body, err := client.Timeframe(timeframeParams)
    var response JsonResponse
    if nil != json.Unmarshal(body, &response) {
        fmt.Println("Error parsing JSON: ", err)
        return
    }

    if response.Valid {
        for date, rates := range response.TimeframeRates {
            fmt.Printf("Date: %s\n", date)
            for currency, rate := range rates {
                fmt.Printf("Currency: %s, Rate: %f\n", currency, rate)
            }
        }
    } else {
        fmt.Printf("API returned error: Code: %d, Message: %s\n", response.Error.Code, response.Error.Message)
    }
}