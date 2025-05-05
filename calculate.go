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
func (r *Rate) Calculate(details DeliveryDetails) (float64, string, error) {
	var provider Provider
	switch r.Provider {
	case "DHL":
		provider = NewDHLProvider(r.ApiKey, r.ApiSecret)
	case "FEDEX":
		provider = NewFedExProvider(r.ApiKey, r.ApiSecret)
	default:
		provider = r // fallback to base Rate logic
	}
	provider.SetLogger(&r.Logger)
	if err := provider.SetConfig(r.Config); err != nil {
		return 0, "", err
	}
	return provider.Calculate(details)
}
