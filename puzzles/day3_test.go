package puzzles

import "testing"

func TestVoltage(t *testing.T) {
	input := 23222224242
	expected := 12

	result := checkFullPatternMatch(input)

	if result != expected {
		t.Error("Expected a value")
	}
}
