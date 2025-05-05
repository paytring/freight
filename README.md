# Freight Rate Calculator

A Go library for calculating freight and shipping rates from various providers (DHL, FedEx, and more) with a unified interface.

## Features
- Seamless provider switching: Just change the provider name in config, no code changes needed.
- Supports DHL and FedEx APIs (with full authentication and package details).
- Returns both rate and currency for each calculation.
- Extensible: Add more providers easily.

## Installation

```bash
go get github.com/paytring/freight
```

## Usage Example

```go
package main

import (
	"fmt"
	"github.com/paytring/freight"
	"github.com/rs/zerolog/log"
)

func main() {
	// Choose provider: "DHL" or "FEDEX"
	rate := freight.NewRate("FEDEX", "YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET")
	rate.SetLogger(&log.Logger)
	rate.SetConfig(map[string]string{
		"accountNumber":         "YOUR_FEDEX_ACCOUNT_NUMBER",
			"serviceType":           "FEDEX_INTERNATIONAL_PRIORITY",
		"pickupType":            "CONTACT_FEDEX_TO_SCHEDULE",
		"packagingType":         "FEDEX_25KG_BOX",
		"weightUnit":            "LB",
	})

	details := freight.DeliveryDetails{
		Weight: 22.0,
		Dimensions: freight.DeliveryDimensions{
			Length: 10.0,
			Width:  8.0,
			Height: 2.0,
		},
		OriginCountryCode:      "CA",
		OriginCityName:         "Toronto",
		OriginPostalCode:       "m1m1m1",
		DestinationCountryCode: "US",
		DestinationCityName:    "Memphis",
		DestinationPostalCode:  "38116",
	}

	rateValue, currency, err := rate.Calculate(details)
	if err != nil {
		log.Fatal().Err(err).Msg("Error calculating rate")
	}
	fmt.Printf("Calculated Rate: %.2f %s\n", rateValue, currency)
}
```

## FedEx API Details
See [docs/fedex.md](docs/fedex.md) for full request/response examples and authentication instructions.

## Testing

```bash
go test ./
```

## License
MIT or your preferred license.
