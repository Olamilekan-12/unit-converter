package main

import "testing"

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		input float64
		want  float64
	}{
		{100, 212},
		{0, 32},
		{-40, -40},
	}
	for _, tt := range tests {
		got := CelsiusToFahrenheit(tt.input)
		want := tt.want

		if got != want {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want=%v", tt.input, got, tt.want)
		}
	}

}
