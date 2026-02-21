package currencyapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func TestGet(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/some_endpoint"
	jsonResponse := `{"valid": true, "data": "some_data"}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"output": "JSON",
		"param":  "value",
	}

	body, err := client.get("some_endpoint", params)
	if err != nil {
		t.Errorf("Error fetching data: %v", err)
	}

	expectedResponseBody := `{"valid": true, "data": "some_data"}`
	if string(body) != expectedResponseBody {
		t.Errorf("Unexpected response body. Got: %s, Expected: %s", body, expectedResponseBody)
	}

	var responseMap map[string]interface{}
	if err := json.Unmarshal(body, &responseMap); err != nil {
		t.Errorf("Error unmarshalling response body: %v", err)
	}

	if validValue, ok := responseMap["valid"].(bool); !ok || !validValue {
		t.Errorf("Unexpected value for 'valid' field")
	}
}

func TestRates(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/rates"
	jsonResponse := `{"valid": true, "rates": {"USD": 1.0, "EUR": 0.85}}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"output": "JSON",
		"base":   "USD",
	}

	body, err := client.Rates(params)
	if err != nil {
		t.Errorf("Error fetching rates: %v", err)
	}

	expectedResponseBody := `{"valid": true, "rates": {"USD": 1.0, "EUR": 0.85}}`
	assert.Equal(t, expectedResponseBody, string(body), "Unexpected response body")

	var responseMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &responseMap))

	if rates, ok := responseMap["rates"].(map[string]interface{}); !ok || len(rates) != 2 {
		t.Errorf("Unexpected value for 'rates' field")
	}
}

func TestCurrencies(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/currencies"
	jsonResponse := `{"valid": true, "currencies": ["USD", "EUR", "GBP"]}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"output": "JSON",
	}

	body, err := client.Currencies(params)
	if err != nil {
		t.Errorf("Error fetching currencies: %v", err)
	}

	expectedResponseBody := `{"valid": true, "currencies": ["USD", "EUR", "GBP"]}`
	assert.Equal(t, expectedResponseBody, string(body), "Unexpected response body")

	var responseMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &responseMap))

	if currencies, ok := responseMap["currencies"].([]interface{}); !ok || len(currencies) != 3 {
		t.Errorf("Unexpected value for 'currencies' field")
	}
}

func TestConvert(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/convert"
	jsonResponse := `{"valid": true, "result": 85.0}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"output": "JSON",
		"to":     "EUR",
		"from":   "USD",
		"amount": "100",
	}

	body, err := client.Convert(params)
	if err != nil {
		t.Errorf("Error converting currency: %v", err)
	}

	expectedResponseBody := `{"valid": true, "result": 85.0}`
	assert.Equal(t, expectedResponseBody, string(body), "Unexpected response body")

	var responseMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &responseMap))

	if result, ok := responseMap["result"].(float64); !ok || result != 85.0 {
		t.Errorf("Unexpected value for 'result' field")
	}
}

func TestHistory(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/history"
	jsonResponse := `{"valid": true, "rates": {"2023-10-01": {"USD": 1.0, "EUR": 0.85}, "2023-10-02": {"USD": 1.02, "EUR": 0.84}}}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"output": "JSON",
		"base":   "USD",
		"date":   "2023-10-01",
	}

	body, err := client.History(params)
	if err != nil {
		t.Errorf("Error fetching history: %v", err)
	}

	expectedResponseBody := `{"valid": true, "rates": {"2023-10-01": {"USD": 1.0, "EUR": 0.85}, "2023-10-02": {"USD": 1.02, "EUR": 0.84}}}`
	assert.Equal(t, expectedResponseBody, string(body), "Unexpected response body")

	var responseMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &responseMap))

	if rates, ok := responseMap["rates"].(map[string]interface{}); !ok || len(rates) != 2 {
		t.Errorf("Unexpected value for 'rates' field")
	}
}

func TestTimeframe(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/timeframe"
	jsonResponse := `{"valid": true, "rates": {"2023-10-01": {"USD": 1.0, "EUR": 0.85}, "2023-10-02": {"USD": 1.02, "EUR": 0.84}}}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"output":     "JSON",
		"base":       "USD",
		"start_date": "2023-10-01",
		"end_date":   "2023-10-02",
	}

	body, err := client.Timeframe(params)
	if err != nil {
		t.Errorf("Error fetching timeframe: %v", err)
	}

	expectedResponseBody := `{"valid": true, "rates": {"2023-10-01": {"USD": 1.0, "EUR": 0.85}, "2023-10-02": {"USD": 1.02, "EUR": 0.84}}}`
	assert.Equal(t, expectedResponseBody, string(body), "Unexpected response body")

	var responseMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &responseMap))

	if rates, ok := responseMap["rates"].(map[string]interface{}); !ok || len(rates) != 2 {
		t.Errorf("Unexpected value for 'rates' field")
	}
}


func TestOhlc(t *testing.T) {
	defer gock.Off()

	client := Client("testAPIKey")

	url := "https://currencyapi.net/api/v2/ohlc"
	jsonResponse := `{"valid": true, "base": "USD", "quote": "EUR", "date": "2023-12-25", "interval": "1d", "ohlc": [{"start": "2023-12-25T00:00:00Z", "open": 0.92, "high": 0.925, "low": 0.918, "close": 0.922}]}`

	gock.New(url).
		MatchParam("key", "testAPIKey").
		Reply(200).
		JSON([]byte(jsonResponse))

	params := map[string]string{
		"currency": "EUR",
		"date":     "2023-12-25",
		"interval": "1d",
	}

	body, err := client.Ohlc(params)
	if err != nil {
		t.Errorf("Error fetching OHLC data: %v", err)
	}

	var responseMap map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &responseMap))

	assert.Equal(t, true, responseMap["valid"], "Expected valid to be true")
	assert.Equal(t, "USD", responseMap["base"], "Expected base to be USD")
	assert.Equal(t, "EUR", responseMap["quote"], "Expected quote to be EUR")
	assert.Equal(t, "2023-12-25", responseMap["date"], "Expected date to match")
	assert.Equal(t, "1d", responseMap["interval"], "Expected interval to be 1d")

	ohlc, ok := responseMap["ohlc"].([]interface{})
	if !ok || len(ohlc) != 1 {
		t.Errorf("Unexpected value for 'ohlc' field")
	}
}

func TestClient(t *testing.T) {
	apiKey := "testAPIKey"
	client := Client(apiKey)

	assert.Equal(t, apiKey, client.APIKey, "API key mismatch")
}

func stringPtr(s string) *string {
	return &s
}