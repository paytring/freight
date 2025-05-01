# Freight Rate Calculator

[![Go Reference](https://pkg.go.dev/badge/paytring/freight.svg)](https://pkg.go.dev/paytring/freight)

A Go library for calculating freight and shipping rates from various providers.

## Description

This library provides a common interface and base implementation for calculating shipping costs. It allows integrating different shipping providers (like DHL, FedEx, etc.) through specific provider implementations.

## Features

*   Basic rate calculation based on weight and dimensions.
*   Provider-specific rate calculation (currently supports DHL via its API).
*   Configurable settings for API keys, secrets, and provider-specific parameters.
*   Structured logging using `zerolog`.

## Installation

To use this library in your Go project:

```bash
go get paytring/freight
```

## Usage

```go
package main

import (
	"fmt"
	"paytring/freight"

	"github.com/rs/zerolog/log"
)

func main() {
	// --- Basic Calculation Example ---
	basicRate := freight.NewRate("basic", "", "") // Provider name can be anything for basic
	basicRate.SetLogger(&log.Logger)

	details := freight.DeliveryDetails{
		Weight: 10.0,
		Dimensions: freight.DeliveryDimensions{
			Length: 15.0,
			Width:  10.0,
			Height: 5.0,
		},
			OriginCountryCode:    "US", // Updated field
		OriginCityName:       "New York", // New field
		DestinationCountryCode: "US", // Updated field
		DestinationCityName:  "Los Angeles", // New field
	}

	calculatedRate, err := basicRate.Calculate(details)
	if err != nil {
		log.Fatal().Err(err).Msg("Error calculating basic rate")
	}
	fmt.Printf("Basic Calculated Rate: %.2f\n", calculatedRate)

	// --- DHL Provider Example ---
	// Replace with your actual API Key and Secret
	dhlAPIKey := "YOUR_DHL_API_KEY"
	dhlAPISecret := "YOUR_DHL_API_SECRET"

	dhlProvider := freight.NewDHLProvider(dhlAPIKey, dhlAPISecret)
	dhlProvider.SetLogger(&log.Logger)

	// Set necessary configuration for DHL
	dhlConfig := map[string]string{
		"accountNumber": "YOUR_DHL_ACCOUNT_NUMBER",
		// Add other DHL specific config if needed
	}
	if err := dhlProvider.SetConfig(dhlConfig); err != nil {
		log.Fatal().Err(err).Msg("Error setting DHL config")
	}

	// Use the same delivery details (already updated structure)
	dhlRate, err := dhlProvider.Calculate(details) // DHL provider uses the updated details struct
	if err != nil {
		log.Fatal().Err(err).Msg("Error calculating DHL rate")
	}
	fmt.Printf("DHL Calculated Rate: %.2f\n", dhlRate)

}
```

**Note:** The DHL provider requires specific configuration (like `accountNumber`) and needs the API response parsing logic to be fully implemented in `dhl.go`.

## Configuration

*   **API Key & Secret:** Required for providers that use API authentication.
*   **Provider Config:** Use the `SetConfig` method to pass provider-specific parameters (e.g., account numbers, service types).

## Testing

Run the tests using the standard Go tool:

```bash
go test ./
```

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues.

## License

(Specify your license here, e.g., MIT License)
