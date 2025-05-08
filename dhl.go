package freight

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const dhlAPIURL = "https://api-mock.dhl.com/mydhlapi/rates"

// DHLProvider implements the rate calculation logic for DHL.
type DHLProvider struct {
	Rate // Embed the base Rate struct
}

// NewDHLProvider creates a new instance of DHLProvider.
func NewDHLProvider(apiKey string, apiSecret string) *DHLProvider {
	return &DHLProvider{
		Rate: Rate{
			Provider:  DHL, // Assuming DHL constant is defined elsewhere
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
}

// Calculate fetches the rate from the DHL API.
// Note: This is a basic implementation based on the curl command.
// Error handling, response parsing, and configuration need refinement.
func (d *DHLProvider) Calculate(details DeliveryDetails) (float64, string, error) {
	d.Logger.Info().Msg("Calculating rate for delivery via DHL")

	client := resty.New()

	// Basic Auth
	auth := base64.StdEncoding.EncodeToString([]byte(d.ApiKey + ":" + d.ApiSecret))

	// Prepare Query Params
	queryParams := url.Values{}

	// Use the new explicit fields and add validation
	if details.OriginAddress.CountryCode == "" || details.OriginAddress.City == "" || details.DestinationAddress.CountryCode == "" || details.DestinationAddress.City == "" {
		return 0, "", fmt.Errorf("origin and destination country codes and city names are required")
	}

	queryParams.Set("accountNumber", d.Config["accountNumber"]) // Assuming accountNumber is in Config
	queryParams.Set("originCountryCode", details.OriginAddress.CountryCode)
	queryParams.Set("originCityName", details.OriginAddress.City)
	queryParams.Set("destinationCountryCode", details.DestinationAddress.CountryCode)
	queryParams.Set("destinationCityName", details.DestinationAddress.City)
	queryParams.Set("weight", strconv.FormatFloat(details.Weight, 'f', -1, 64))
	queryParams.Set("length", strconv.FormatFloat(details.Dimensions.Length, 'f', -1, 64))
	queryParams.Set("width", strconv.FormatFloat(details.Dimensions.Width, 'f', -1, 64))
	queryParams.Set("height", strconv.FormatFloat(details.Dimensions.Height, 'f', -1, 64))
	queryParams.Set("plannedShippingDate", time.Now().Format("2006-01-02")) // Example date format
	queryParams.Set("isCustomsDeclarable", "false")                         // Example value
	queryParams.Set("unitOfMeasurement", "metric")                          // Example value, might need config

	// Prepare Headers - Using placeholders for most values as per curl example
	headers := map[string]string{
		"Authorization":          "Basic " + auth,
		"Message-Reference":      "SOME_STRING_VALUE",             // Replace with dynamic or configured value
		"Message-Reference-Date": time.Now().Format(time.RFC3339), // Example format
		// Add other headers from curl example, potentially from config
		"Plugin-Name":                      "SOME_STRING_VALUE",
		"Plugin-Version":                   "SOME_STRING_VALUE",
		"Shipping-System-Platform-Name":    "SOME_STRING_VALUE",
		"Shipping-System-Platform-Version": "SOME_STRING_VALUE",
		"Webstore-Platform-Name":           "SOME_STRING_VALUE",
		"Webstore-Platform-Version":        "SOME_STRING_VALUE",
		"x-version":                        "SOME_STRING_VALUE",
	}

	resp, err := client.R().
		SetQueryParamsFromValues(queryParams).
		SetHeaders(headers).
		// EnableTrace(). // Uncomment for debugging
		Get(dhlAPIURL)

	if err != nil {
		d.Logger.Error().Err(err).Msg("DHL API request failed")
		return 0, "", fmt.Errorf("failed to call DHL API: %w", err)
	}

	if resp.IsError() {
		d.Logger.Error().Str("status", resp.Status()).Bytes("body", resp.Body()).Msg("DHL API returned error")
		return 0, "", fmt.Errorf("DHL API error: %s", resp.Status())
	}

	d.Logger.Info().Str("status", resp.Status()).Msg("DHL API request successful")
	d.Logger.Debug().Bytes("body", resp.Body()).Msg("DHL API Response Body") // Log raw response for inspection

	// Parse response for rate and currency
	var result struct {
		Products []struct {
			TotalPrice float64 `json:"totalPrice"`
			Currency   string  `json:"currency"`
		} `json:"products"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		d.Logger.Error().Err(err).Msg("Failed to parse DHL response")
		return 0, "", fmt.Errorf("failed to parse DHL response: %w", err)
	}
	if len(result.Products) == 0 {
		return 0, "", fmt.Errorf("no rate found in DHL response")
	}
	rate := result.Products[0].TotalPrice
	currency := result.Products[0].Currency
	d.Logger.Info().Float64("rate", rate).Str("currency", currency).Msg("DHL rate extracted")
	return rate, currency, nil
}
