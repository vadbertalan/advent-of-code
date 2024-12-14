package coordinate

import (
	"testing"
)

func TestParseCoordStr(t *testing.T) {
	tests := []struct {
		input     string
		separator string
		expected  Coord
	}{
		{"3-4", "-", Coord{Row: 3, Col: 4}},
		{"10,20", ",", Coord{Row: 10, Col: 20}},
		{"-5|15", "|", Coord{Row: -5, Col: 15}},
		{"0:0", ":", Coord{Row: 0, Col: 0}},
	}

	for _, test := range tests {
		result := ParseCoordStr(test.input, test.separator)
		if result != test.expected {
			t.Errorf("ParseCoordStr(%q, %q) = %v; want %v", test.input, test.separator, result, test.expected)
		}
	}
}

func TestParseCoordStrInvalidInput(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ParseCoordStr did not panic on invalid input")
		}
	}()

	ParseCoordStr("invalid-input", "-")
}
