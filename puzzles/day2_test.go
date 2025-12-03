package puzzles

import "testing"

func TestFullPatterMatch(t *testing.T) {
	input := 123123123
	expected := 123123123

	result := checkFullPatternMatch(input)

	if result != expected {
		t.Error("Expected a value")
	}
}
