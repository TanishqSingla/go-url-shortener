package base62

import (
	"testing"
)

func TestToBase62(t *testing.T) {
	input := 123456789089898
	expected := "z3wBXxG2"

	result := ToBase62(input)

	if result != expected {
		t.Fatalf("expected: %s for input: %d but got %s", expected, input, result)
	}
}

func TestToBase10(t *testing.T) {
	input := "z3wBXxG2"
	expected := 123456789089898

	result := ToBase10(input)

	if result != expected {
		t.Fatalf("expected: %d for input: %s but got %d", expected, input, result)
	}
}
