package main

import (
	"fmt"
	"encoding/json"
	"github.com/houseofapis/currencyapi-go"
)	

type JsonResponse struct {
	Valid bool `json:"valid"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
	Result map[string]interface{} `json:"rates"`
}

func main() {
	client := currencyapi.Client("5dhVka7ajXAzvOHAir5jojH5vPZCiMFxzQCb")

	ratesParams := map[string]string{
		"output": "JSON",
		"base":   "USD",
	}

	body, err := client.Rates(ratesParams)
	var response JsonResponse
	if nil != json.Unmarshal(body, &response) {
		fmt.Println("Error parsing JSON: ", err)
		return
	}

	if response.Valid {
		fmt.Println(response.Result["AED"])
	} else {
		fmt.Printf("API returned error: Code: %d, Message: %s\n", response.Error.Code, response.Error.Message)
	}
}