package freight

import (
	"testing"
)

func TestDefaultPriceCalculation(t *testing.T) {
	// Create a new Rate instance
	rate := NewRate("test_provider", "test_api_key", "test_api_secret")

	rate.SetLogger(nil)

	// Set up the delivery details with new fields
	details := DeliveryDetails{
		Weight: 10.0,
		Dimensions: DeliveryDimensions{
			Length: 22.91,
			Width:  5.0,
			Height: 5.0,
		},
		OriginCountryCode:      "US",          // New field
		OriginCityName:         "New York",    // New field
		DestinationCountryCode: "US",          // New field
		DestinationCityName:    "Los Angeles", // New field
	}

	expectedRate := 25.0 // Expected rate based on the weight and dimensions

	calculatedRate, err := rate.Calculate(details)

	if err != nil {
		t.Fatalf("Error calculating rate: %v", err)
	}
	if calculatedRate != expectedRate {
		t.Errorf("Expected rate %v, got %v", expectedRate, calculatedRate)
	}

	t.Logf("Expected rate %v, got %v", expectedRate, calculatedRate)
}
