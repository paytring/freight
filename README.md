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
		"accountNumber":        "YOUR_FEDEX_ACCOUNT_NUMBER",
		"serviceType":          "FEDEX_INTERNATIONAL_PRIORITY",
		"accountKey": 			"YOUR_FEDEX_ACCOUNT_KEY"
		"packagingType": 		"FEDEX_10KG_BOX",
		"weightUnit":    		"KG",
		"pickupType":    		"DROPOFF_AT_FEDEX_LOCATION",
	})

	details := freight.DeliveryDetails{
		Weight: 15.0,
		Dimensions: freight.DeliveryDimensions{
			Length: 10.0,
			Width:  10.0,
			Height: 10.0,
		},
		OriginAddress: freight.Address{
			City:                "Chennai",
			PostalCode:          "600003",
			CountryCode:         "IN",
			Residential:         false,
			StateOrProvinceCode: "TN",
		},
		DestinationAddress: freight.Address{
			City:                "Oklahoma City",
			PostalCode:          "73102",
			CountryCode:         "US",
			Residential:         false,
			StateOrProvinceCode: "OK",
		},
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
