package freight

type DeliveryDetails struct {
	Weight                 float64
	Dimensions             DeliveryDimensions
	OriginCountryCode      string // Changed from Origin
	OriginCityName         string // New field
	DestinationCountryCode string // Changed from Destination
	DestinationCityName    string // New field
}

type DeliveryDimensions struct {
	Length float64
	Width  float64
	Height float64
}

func (r *Rate) Calculate(details DeliveryDetails) (float64, error) {

	r.Logger.Info().Msg("Calculating rate for delivery")

	finalRate := details.Weight * 1.5

	if details.Dimensions.Length > 20 || details.Dimensions.Width > 20 || details.Dimensions.Height > 20 {
		finalRate += 10.0
	}

	r.Logger.Info().Msgf("Final rate calculated: %v", finalRate)

	return finalRate, nil

}
