package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

// refactored the above test. This reads better
func TestCalculation(t *testing.T) {
	got := Calculate(2)
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestTableCalculate(t *testing.T) {
	// an array of type struct
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{0, 2},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, receieved: {}", test.input, test.expected, output)
		}
	}
}

func TestMultiply(t *testing.T) {
	if Multiply(3) != 9 {
		t.Error("Expected 3 * 3 to equal 9")
	}

}

// This test reads better
func TestMultiplication(t *testing.T) {
	got := Multiply(3)
	want := 9

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
