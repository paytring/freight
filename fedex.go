package freight

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

const fedexAPIURL = "https://api.fedex.com/rate/v1/rates/quotes"

type FedExProvider struct {
	Rate
	bearerToken string
	tokenExpiry int64
}

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
		return f.bearerToken, nil
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
		SetHeader("accept", "application/json, text/plain, */*").
		SetHeader("accept-language", "en-GB,en-US;q=0.9,en;q=0.8").
		SetHeader("cache-control", "no-cache").
		SetHeader("content-type", "application/x-www-form-urlencoded;charset=UTF-8").
		SetHeader("origin", "https://www.fedex.com").
		SetHeader("referer", "https://www.fedex.com/secure-login/en-in/").
		SetFormData(map[string]string{
			"grant_type":    "client_credentials",
			"client_id":     clientID,
			"client_secret": clientSecret,
		}).
		Post("https://api.fedex.com/auth/oauth/v2/token")

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

func fetchFedExRates(clientID, clientSecret string, payload map[string]interface{}) (map[string]interface{}, error) {
	client := resty.New()

	// Fetch Bearer Token
	token, expiry, err := fetchFedExAuthTokenWithExpiry(clientID, clientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch FedEx auth token: %w", err)
	}
	if expiry <= time.Now().Unix() {
		return nil, fmt.Errorf("FedEx auth token expired")
	}

	// Make API request
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+token).
		SetHeader("origin", "https://www.fedex.com").
		SetHeader("referer", "https://www.fedex.com/en-in/online/rating.html").
		SetHeader("x-clientid", "MAGR").
		SetHeader("x-clientversion", "1.0").
		SetHeader("x-locale", "en_IN").
		SetHeader("Content-Type", "application/json").
		SetBody(payload).
		Post("https://api.fedex.com/rate/v2/rates/quotes")

	if err != nil {
		return nil, fmt.Errorf("failed to fetch FedEx rates: %w", err)
	}
	if resp.IsError() {
		return nil, fmt.Errorf("FedEx rates API error: %s", resp.Status())
	}

	// Parse response
	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, fmt.Errorf("failed to parse FedEx rates response: %w", err)
	}

	return result, nil
}

// Updated the FedExProvider to include available service options, their codes, descriptions, delivery dates, and pricing parameters.
func (f *FedExProvider) Calculate(details DeliveryDetails) ([]FrightRate, error) {
	f.Logger.Info().Msg("Calculating rate for delivery via FedEx")

	// Build request payload
	payload := map[string]interface{}{
		"accountNumber": map[string]interface{}{
			"value": f.Config["accountNumber"],
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
			"rateRequestType": []string{"ACCOUNT", "LIST"},
			"requestedPackageLineItems": []map[string]interface{}{
				{
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

	// Sample response for testing purposes

	sampleResponse := []FrightRate{}
	sampleResponse = append(sampleResponse, FrightRate{
		Title:       "FedEx Express Saver",
		Description: "Delivery via FedEx By 2023-10-15 17:00",
		Handle:      "FEDEX_EXPRESS_SAVER",
		Price:       100.0,
		Currency:    "USD",
	})

	sampleResponse = append(sampleResponse, FrightRate{
		Title:       "FedEx 2Day",
		Description: "Delivery via FedEx By 2023-10-15 17:00",
		Handle:      "FEDEX_2DAY",
		Price:       150.0,
		Currency:    "USD",
	})

	return sampleResponse, nil

	// Uncomment the following lines to use the actual FedEx API

	// Use fetchFedExRates to get rates
	result, err := fetchFedExRates(f.ApiKey, f.ApiSecret, payload)
	if err != nil {
		f.Logger.Error().Err(err).Msg("Failed to fetch FedEx rates")
		return nil, fmt.Errorf("failed to fetch FedEx rates: %w", err)
	}

	// Parse response to extract service options
	output, ok := result["output"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format from FedEx API")
	}

	rateReplyDetails, ok := output["rateReplyDetails"].([]interface{})
	if !ok || len(rateReplyDetails) == 0 {
		return nil, fmt.Errorf("no service options found in FedEx response")
	}

	// Extract service options
	serviceOptions := []FrightRate{}
	for _, detail := range rateReplyDetails {
		detailMap, ok := detail.(map[string]interface{})
		if !ok {
			continue
		}

		serviceType, _ := detailMap["serviceType"].(string)
		serviceName, _ := detailMap["serviceName"].(string)
		commit, _ := detailMap["commit"].(map[string]interface{})
		dateDetail, _ := commit["dateDetail"].(map[string]interface{})
		deliveryDate, _ := dateDetail["day"].(string)
		deliveryTime, _ := dateDetail["time"].(string)

		ratedShipmentDetails, _ := detailMap["ratedShipmentDetails"].([]interface{})
		if len(ratedShipmentDetails) == 0 {
			continue
		}
		firstRatedDetail, _ := ratedShipmentDetails[0].(map[string]interface{})
		price, _ := firstRatedDetail["totalNetCharge"].(float64)
		currency, _ := firstRatedDetail["currency"].(string)

		serviceOptions = append(serviceOptions, FrightRate{
			Title:       serviceName,
			Description: "Delivery via FedEx By " + deliveryDate + " " + deliveryTime,
			Handle:      serviceType,
			Price:       price,
			Currency:    currency,
		})
	}

	return serviceOptions, nil
}
