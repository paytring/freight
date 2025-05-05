package freight

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const fedexAPIURL = "https://apis-sandbox.fedex.com/rate/v1/rates/quotes"

// FedExProvider implements the rate calculation logic for FedEx.
type FedExProvider struct {
	Rate        // Embed the base Rate struct
	bearerToken string
	tokenExpiry int64 // Unix timestamp
}

// NewFedExProvider creates a new instance of FedExProvider.
func NewFedExProvider(apiKey, apiSecret string) *FedExProvider {
	return &FedExProvider{
		Rate: Rate{
			Provider:  "FEDEX",
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		},
	}
}

func (f *FedExProvider) getBearerToken() (string, error) {
	if f.bearerToken != "" && f.tokenExpiry > 0 && f.tokenExpiry > time.Now().Unix()+60 {
		return f.bearerToken, nil // Token is still valid
	}
	clientID := f.ApiKey
	clientSecret := f.ApiSecret
	token, expiry, err := fetchFedExAuthTokenWithExpiry(clientID, clientSecret)
	if err != nil {
		return "", err
	}
	f.bearerToken = token
	f.tokenExpiry = expiry
	return token, nil
}

func fetchFedExAuthTokenWithExpiry(clientID, clientSecret string) (string, int64, error) {
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     clientID,
			"client_secret": clientSecret,
		}).
		Post("https://apis-sandbox.fedex.com/oauth/token")

	if err != nil {
		return "", 0, fmt.Errorf("failed to get FedEx auth token: %w", err)
	}
	if resp.IsError() {
		return "", 0, fmt.Errorf("FedEx auth error: %s", resp.Status())
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int64  `json:"expires_in"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", 0, fmt.Errorf("failed to parse FedEx auth response: %w", err)
	}
	if result.AccessToken == "" {
		return "", 0, fmt.Errorf("FedEx auth token missing in response")
	}
	expiry := time.Now().Unix() + result.ExpiresIn
	return result.AccessToken, expiry, nil
}

// Calculate fetches the rate from the FedEx API.
func (f *FedExProvider) Calculate(details DeliveryDetails) (float64, string, error) {
	f.Logger.Info().Msg("Calculating rate for delivery via FedEx")

	client := resty.New()

	accountNumber := f.Config["accountNumber"]
	if accountNumber == "" {
		return 0, "", fmt.Errorf("FedEx accountNumber is required in config")
	}

	token, err := f.getBearerToken()
	if err != nil {
		return 0, "", fmt.Errorf("could not get FedEx bearer token: %w", err)
	}

	// Build request payload
	payload := map[string]interface{}{
		"accountNumber": map[string]interface{}{
			"value": accountNumber,
		},
		"requestedShipment": map[string]interface{}{
			"shipper": map[string]interface{}{
				"address": map[string]interface{}{
					"postalCode":  details.OriginPostalCode,
					"countryCode": details.OriginCountryCode,
				},
			},
			"recipient": map[string]interface{}{
				"address": map[string]interface{}{
					"postalCode":  details.DestinationPostalCode,
					"countryCode": details.DestinationCountryCode,
					"residential": true,
				},
			},
			"serviceType":     f.Config["serviceType"],
			"pickupType":      f.Config["pickupType"],
			"packagingType":   f.Config["packagingType"],
			"rateRequestType": []string{"ACCOUNT", "LIST"},
			"requestedPackageLineItems": []map[string]interface{}{
				{
					"subPackagingType": "CASE",
					// "declaredValue": map[string]interface{}{
					// 	"amount":   100,
					// 	"currency": "CAD",
					// },
					"weight": map[string]interface{}{
						"units": f.Config["weightUnit"],
						"value": details.Weight,
					},
					"dimensions": map[string]interface{}{
						"length": details.Dimensions.Length,
						"width":  details.Dimensions.Width,
						"height": details.Dimensions.Height,
					},
				},
			},
		},
	}

	resp, err := client.R().
		SetHeader("X-locale", "en_US").
		SetHeader("Content-Type", "application/json").
		SetHeader("authorization", "Bearer "+token).
		SetBody(payload).
		Post(fedexAPIURL)

	if err != nil {
		f.Logger.Error().Err(err).Msg("FedEx API request failed")
		return 0, "", fmt.Errorf("failed to call FedEx API: %w", err)
	}

	if resp.IsError() {
		f.Logger.Error().Str("status", resp.Status()).Bytes("body", resp.Body()).Msg("FedEx API returned error")
		return 0, "", fmt.Errorf("FedEx API error: %s", resp.Status())
	}

	f.Logger.Info().Str("status", resp.Status()).Msg("FedEx API request successful")
	f.Logger.Debug().Bytes("body", resp.Body()).Msg("FedEx API Response Body")

	// Parse response to extract the rate and currency
	var result struct {
		Output struct {
			RateReplyDetails []struct {
				RatedShipmentDetails []struct {
					TotalNetCharge float64 `json:"totalNetCharge"`
					Currency       string  `json:"currency"`
				} `json:"ratedShipmentDetails"`
			} `json:"rateReplyDetails"`
		} `json:"output"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		f.Logger.Error().Err(err).Msg("Failed to parse FedEx response")
		return 0, "", fmt.Errorf("failed to parse FedEx response: %w", err)
	}

	if len(result.Output.RateReplyDetails) == 0 || len(result.Output.RateReplyDetails[0].RatedShipmentDetails) == 0 {
		return 0, "", fmt.Errorf("no rate found in FedEx response")
	}

	rate := result.Output.RateReplyDetails[0].RatedShipmentDetails[0].TotalNetCharge
	currency := result.Output.RateReplyDetails[0].RatedShipmentDetails[0].Currency
	f.Logger.Info().Float64("rate", rate).Str("currency", currency).Msg("FedEx rate extracted")
	return rate, currency, nil
}
