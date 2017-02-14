package server

import "testing"

type expected struct {
	indicator *Indicator
	variation string
}

type testcases struct {
	input    *Restaurant
	expected *expected
}

// Test variation modulo
func TestNewIndicatorVariation(t *testing.T) {

	var testCases = []*testcases{
		{
			&Restaurant{Variation: "1"},
			&expected{variation: "1"},
		},
		{
			&Restaurant{Variation: "0"},
			&expected{variation: "0"},
		},
		{
			&Restaurant{Variation: "-1"},
			&expected{variation: "1"},
		},
	}

	for _, test := range testCases {
		_ = test.input.newIndicator()

		if test.input.Variation != test.expected.variation {
			t.Errorf("For variation: %s. Expected: %s. Received: %s",
				test.input.Variation, test.expected.variation, test.input.Variation)
		}
	}
}

// Test indicator direction for variation values
func TestNewIndicatorDirection(t *testing.T) {

	var testCases = []*testcases{
		{
			&Restaurant{Variation: "1"},
			&expected{indicator: &Indicator{Direction: "⇧"}},
		},
		{
			&Restaurant{Variation: "0"},
			&expected{indicator: &Indicator{Direction: "▭"}},
		},
		{
			&Restaurant{Variation: "-1"},
			&expected{indicator: &Indicator{Direction: "⇩"}},
		},
	}

	for _, test := range testCases {
		observed := test.input.newIndicator()

		if observed.Direction != test.expected.indicator.Direction {
			t.Errorf("For variation: %s. Expected: %s. Received: %s",
				test.input.Variation, test.expected.indicator.Direction, observed.Direction)
		}
	}
}

// Test indicator colour for variation values
func TestNewIndicatorColour(t *testing.T) {

	var testCases = []*testcases{
		{
			&Restaurant{Variation: "1"},
			&expected{indicator: &Indicator{Colour: "green"}},
		},
		{
			&Restaurant{Variation: "0"},
			&expected{indicator: &Indicator{}},
		},
		{
			&Restaurant{Variation: "-1"},
			&expected{indicator: &Indicator{Colour: "red"}},
		},
	}

	for _, test := range testCases {
		observed := test.input.newIndicator()

		if observed.Colour != test.expected.indicator.Colour {
			t.Errorf("For colour: %s. Expected: %s. Received: %s",
				test.input.Indicator.Colour, test.input.Indicator.Colour, observed.Colour)
		}
	}
}
