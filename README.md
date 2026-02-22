# CurrencyApi Golang wrapper 

[![Go project version](https://badge.fury.io/go/github.com%2Fhouseofapis%2Fcurrencyapi-go.svg)](https://badge.fury.io/go/github.com%2Fhouseofapis%2Fcurrencyapi-go) [![Coverage Status](https://coveralls.io/repos/github/houseofapis/currencyapi-go/badge.svg?branch=main)](https://coveralls.io/github/houseofapis/currencyapi-go?branch=main) 

> **Note:** API v1 is deprecated and will be retired on **31st July 2026**, at which point all v1 traffic will be redirected to v2. This SDK (v2.0.0+) targets API v2. If you are on an older version of this SDK, please upgrade.

<a href="https://currencyapi.net" title="CurrencyApi">CurrencyApi.net</a> provides live currency rates via a REST API. A live currency feed for over 166 currencies, including physical (USD, GBP, EUR + more) and cryptos (Bitcoin, Litecoin, Ethereum + more). A JSON and XML currency api updated every 60 seconds.

Features:

- Live exchange rates (updated every 60 seconds).
- 166 world currencies.
- Popular cryptocurrencies included; Bitcoin, Litecoin etc.
- Convert currencies on the fly with the convert endpoint.
- Historical currency rates back to year 2000.
- OHLC (Open, High, Low, Close) candlestick data for technical analysis (Tier 3+).
- Easy to follow <a href="https://currencyapi.net/documentation" title="currency-api-documentation">documentation</a>

Signup for a free or paid account <a href="https://currencyapi.net/#pricing-sec" title="currency-api-pricing">here</a>.

## This package is a:

Golang wrapper for <a href="https://currencyapi.net" title="CurrencyApi">CurrencyApi.net</a> endpoints.

## Developer Guide

For an easy to following developer guide, check out our [Golang Developer Guide](https://currencyapi.net/sdk/golang).

Alternatively keep reading below.

#### Prerequisites

- Minimum Go 1.21
- Working on Go 1.26
- Free or Paid account with CurrencyApi.net

#### Test Coverage

- 100% coverage

## Installation

```bash
go get github.com/houseofapis/currencyapi-go
```
then import the package with:

```golang
import (
    ...
	"github.com/houseofapis/currencyapi-go"
	...
)
```

## Usage

### Instantiating

```golang
client := currencyapi.Client("API_KEY")
```

### Live rates:

```golang
params := map[string]string{
    "output": "JSON",
    "base":   "USD",
}

body, err := client.Rates(params)
```

**Available params for rates endpoint**

| Methods | Description |
| --- | --- |
| `base` | The base currency you wish you receive the currency conversions for. This will output all currency conversions for that currency. **Default: USD**. |
| `output` | Response output in either JSON or XML. **Default: JSON**. |

<br>

### List of available currencies:

```golang
params := map[string]string{
    "output": "XML"
}

body, err := client.Currencies(params)
```

**Available methods for currencies endpoint**

| Methods | Description |
| --- | --- |
| `output` | Response output in either JSON or XML. **Default: JSON**. |

<br>

### Convert:

```golang
params := map[string]string{
    "output": "JSON",
    "from": "USD",
    "to": "GBP",
    "amount": 15.99,
}

body, err := client.Convert(params)
```

**Available methods for convert endpoint**

| Methods | Description |
| --- | --- |
| `amount` | The value of the currency you want to convert from. This should be a number and can contain a decimal place. **Required**. |
| `from` | The currency you want to convert. This will be a three letter ISO 4217 currency code from one of the currencies we have rates for. **Required**. |
| `to` | The currency you want to convert the amount 'to'. Again this will be a three letter currency code from the ones we offer. **Required**. |
| `output` | Response output in either JSON or XML. **Default: JSON**. |

<br>

### Historical:

```golang
params := map[string]string{
    "output": "JSON",
    "base": "GBP",
    "date": "2019-01-01"
}

body, err := client.History(params)
```

**Available methods for historical endpoint**

| Methods | Description |
| --- | --- |
| `date` | The historical date you wish to receive the currency conversions for. This should be formatted as YYYY-MM-DD. **Required**. |
| `base` | The base currency you wish you receive the currency conversions for. This will output all currency conversions for that currency. **Default: USD**. |
| `output` | Response output in either JSON or XML. **Default: JSON**. |

<br>

### Timeframe:

```golang
params := map[string]string{
    "output": "JSON",
    "base": "GBP",
    "start_date": "2019-01-01"
    "end_date": "2019-01-05"
}

body, err := client.Timeframe(params)
```

**Available methods for timeframe endpoint**

| Methods | Description |
| --- | --- |
| `start_date` | The historical date you wish to receive the currency conversions from. This should be formatted as YYYY-MM-DD. **Required**. |
| `end_date` | The historical date you wish to receive the currency conversions until. This should be formatted as YYYY-MM-DD. **Required**. |
| `base` | The base currency you wish you receive the currency conversions for. This will output all currency conversions for that currency. **Default: USD**. |
| `output` | Response output in either JSON or XML. **Default: JSON**. |

<br>

### OHLC (Open, High, Low, Close):

```golang
params := map[string]string{
    "quote":    "EUR",
    "date":     "2023-12-25",
    "interval": "1h",
}

body, err := client.Ohlc(params)
```

**Available params for OHLC endpoint**

| Parameter | Description |
| --- | --- |
| `quote` | The quote currency code to retrieve OHLC data for (e.g. `EUR`, `GBP`, `BTC`). **Required**. |
| `date` | The date you wish to retrieve OHLC data for. This should be formatted as YYYY-MM-DD and must be in the past. **Required**. |
| `base` | The base currency for the pair. **Default: USD**. |
| `interval` | The time interval for each OHLC candle. Allowed values: `5m`, `15m`, `30m`, `1h`, `4h`, `12h`, `1d`. **Default: 1d**. |
| `output` | Response output in either JSON or XML. **Default: JSON**. |

> **Note:** The OHLC endpoint requires a Tier 3 (Professional) subscription.

