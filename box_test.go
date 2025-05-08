package freight

import (
	"testing"
)

func TestCalcParentBox(t *testing.T) {
	boxes := []Box{
		{
			Name:       "box1",
			Weight:     10,
			WeightUnit: "kg",
			Dimensions: Dimensions{
				Length: 20,
				Width:  15,
				Height: 10,
			},
			DimensionsUnit: "cm",
		},
		{
			Name:       "box2",
			Weight:     5000,
			WeightUnit: "g",
			Dimensions: Dimensions{
				Length: 30,
				Width:  20,
				Height: 15,
			},
			DimensionsUnit: "cm",
		},
		{
			Name:       "box3",
			Weight:     2.2,
			WeightUnit: "lb",
			Dimensions: Dimensions{
				Length: 25,
				Width:  10,
				Height: 20,
			},
			DimensionsUnit: "cm",
		},
	}

	parentBox, err := CalcParentBox(boxes)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if parentBox.Weight < 15.0 || parentBox.Weight > 16 {
		t.Errorf("expected weight to be around 15.1 kg, got %v", parentBox.Weight)
	}

	if parentBox.Dimensions.Length != 75 {
		t.Errorf("expected total length to be 75 cm, got %v", parentBox.Dimensions.Length)
	}

	if parentBox.Dimensions.Width != 45 {
		t.Errorf("expected total width to be 45 cm, got %v", parentBox.Dimensions.Width)
	}

	if parentBox.Dimensions.Height != 45 {
		t.Errorf("expected total height to be 45 cm, got %v", parentBox.Dimensions.Height)
	}
}
