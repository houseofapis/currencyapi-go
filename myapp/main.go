package main

import (
	"fmt"
	"log"
	"github.com/houseofapis/currencyapi"
)

func main() {
	client := currencyapi.Client("PY9VM5DrVNptjWjQ2dXGp3sgQI61W7jmTiex")

	ratesParams := map[string]string{
		"output": "XML",
		"base":   "USD",
	}

	body, err := client.Rates(ratesParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body)) 
}