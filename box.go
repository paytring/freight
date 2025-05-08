package freight

import (
	"errors"
)

type Box struct {
	Name           string
	Weight         float64
	WeightUnit     string
	Dimensions     Dimensions
	DimensionsUnit string
}

type Dimensions struct {
	Length float64
	Width  float64
	Height float64
}

// Convert weight to kilograms
func parseBoxWeight(box Box) (float64, error) {
	switch box.WeightUnit {
	case "kg":
		return box.Weight, nil
	case "g":
		return box.Weight / 1000, nil
	case "lb":
		return box.Weight * 0.45359237, nil // Use a more precise conversion factor for pounds to kilograms
	default:
		return 0, errors.New("unknown weight unit")
	}
}

// Convert dimensions to centimeters
func parseBoxDimensions(box Box) (Dimensions, error) {
	switch box.DimensionsUnit {
	case "cm":
		return box.Dimensions, nil
	case "m":
		return Dimensions{
			Length: box.Dimensions.Length * 100,
			Width:  box.Dimensions.Width * 100,
			Height: box.Dimensions.Height * 100,
		}, nil
	case "in":
		return Dimensions{
			Length: box.Dimensions.Length * 2.54,
			Width:  box.Dimensions.Width * 2.54,
			Height: box.Dimensions.Height * 2.54,
		}, nil
	default:
		return Dimensions{}, errors.New("unknown dimensions unit")
	}
}

// Validate units for all boxes
func validateUnits(boxes []Box) error {
	for _, box := range boxes {
		if box.WeightUnit != "kg" && box.WeightUnit != "g" && box.WeightUnit != "lb" {
			return errors.New("all weights must be in kg, g, or lb")
		}
		if box.DimensionsUnit != "cm" {
			return errors.New("all dimensions must be in cm")
		}
	}
	return nil
}

// Calculate the parent box
func CalcParentBox(boxes []Box) (Box, error) {
	if err := validateUnits(boxes); err != nil {
		return Box{}, err
	}

	totalWeight := 0.0
	totalLength := 0.0
	totalWidth := 0.0
	totalHeight := 0.0

	for _, box := range boxes {
		weight, err := parseBoxWeight(box)
		if err != nil {
			return Box{}, err
		}
		dimensions, err := parseBoxDimensions(box)
		if err != nil {
			return Box{}, err
		}

		totalWeight += weight
		totalLength += dimensions.Length
		totalWidth += dimensions.Width
		totalHeight += dimensions.Height
	}

	return Box{
		Name:       "parent_box",
		Weight:     totalWeight,
		WeightUnit: "kg",
		Dimensions: Dimensions{
			Length: totalLength,
			Width:  totalWidth,
			Height: totalHeight,
		},
		DimensionsUnit: "cm",
	}, nil
}
