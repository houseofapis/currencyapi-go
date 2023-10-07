package currencyapi

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const baseURL = "https://currencyapi.net/api/v1"

type ClientS struct {
	APIKey string
}

func Client(apiKey string) *ClientS {
	return &ClientS{
		APIKey: apiKey,
	}
}

func (c *ClientS) get(endpoint string, params map[string]string) ([]byte, error) {
	if c.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	url := fmt.Sprintf("%s/%s?key=%s", baseURL, endpoint, c.APIKey)

	// Append other parameters
	for key, value := range params {
		url += fmt.Sprintf("&%s=%s", key, value)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set the custom header
	req.Header.Set("X-Sdk", "go")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *ClientS) Rates(params map[string]string) ([]byte, error) {
	return c.get("rates", params)
}

func (c *ClientS) Currencies(params map[string]string) ([]byte, error) {
	return c.get("currencies", params)
}

func (c *ClientS) Convert(params map[string]string) ([]byte, error) {
	return c.get("convert", params)
}

func (c *ClientS) History(params map[string]string) ([]byte, error) {
	return c.get("history", params)
}

func (c *ClientS) Timeframe(params map[string]string) ([]byte, error) {
	return c.get("timeframe", params)
}