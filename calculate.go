package freight

type DeliveryDetails struct {
	Weight                 float64
	Dimensions             DeliveryDimensions
	OriginCountryCode      string
	OriginCityName         string
	OriginPostalCode       string // Added
	DestinationCountryCode string
	DestinationCityName    string
	DestinationPostalCode  string // Added
}

type DeliveryDimensions struct {
	Length float64
	Width  float64
	Height float64
}

// Rate now acts as a dynamic provider selector
func (r *Rate) Calculate(details DeliveryDetails) ([]map[string]interface{}, error) {
	var provider Provider
	switch r.Provider {
	case "FEDEX":
		provider = NewFedExProvider(r.ApiKey, r.ApiSecret)
	default:
		return r.calculateBase(details)
	}

	provider.SetLogger(&r.Logger)
	if err := provider.SetConfig(r.Config); err != nil {
		return nil, err
	}
	return provider.Calculate(details)

}

// calculateBase handles the base Rate calculation logic
func (r *Rate) calculateBase(details DeliveryDetails) ([]map[string]interface{}, error) {
	// Implement your base rate calculation logic here
	// This is a placeholder implementation
	result := []map[string]interface{}{
		{
			"price":    0.0,
			"provider": "DEFAULT",
		},
	}
	return result, nil
}
