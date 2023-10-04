package currencyapinet

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const baseURL = "https://currencyapi.net/api/v1"

type ClientS struct {
	APIKey string
}

type APIError struct {
	Status int
	Err    error
	Msg    *string
}

func (e *APIError) Error() string {
	if e.Msg != nil {
		return fmt.Sprintf("Error fetching rates. Status code: %d. Error: %s. Message: %s", e.Status, e.Err, *e.Msg)
	}
	return fmt.Sprintf("Error fetching rates. Status code: %d. Error: %s", e.Status, e.Err)
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
		return nil, &APIError{Status: 0, Err: err}
	}

	// Set the custom header
	req.Header.Set("X-Sdk", "go")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, &APIError{Status: 0, Err: err}
	}
	defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        str := string(body)
        return nil, &APIError{
            Status: resp.StatusCode,
            Err:    fmt.Errorf("Non-OK status code: %d", resp.StatusCode),
            Msg:    &str,
        }
    }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, &APIError{Status: resp.StatusCode, Err: err}
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