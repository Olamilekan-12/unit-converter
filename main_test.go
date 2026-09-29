package main

import (
	"math"
	"testing"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 1e-9
}

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{"boiling", 100, 212},
		{"freezing", 0, 32},
		{"minus-forty", -40, -40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.input)
			if !almostEqual(got, tt.want) {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want=%v", tt.input, got, tt.want)
			}
		})

	}

}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{"boiling", 212, 100},
		{"freezing", 32, 0},
		{"body-temperature", 98.6, 37},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FahrenheitToCelsius(tt.input)
			if !almostEqual(got, tt.want) {
				t.Errorf("FahrenheitToCelsius(%v)=%v, want=%v", tt.input, got, tt.want)
			}
		})

	}

}

func TestKilometersToMiles(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{"one kilometer", 1, 0.621371},
		{"ten kilometers", 10, 6.21371},
		{"hundred kilometers", 100, 62.1371},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KilometersToMiles(tt.input)
			if !almostEqual(got, tt.want) {
				t.Errorf("KilometersToMiles(%v)=%v, want=%v", tt.input, got, tt.want)
			}
		})
	}
}

func TestKilogramsToPounds(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{"one_kilogram_to_pounds", 1, 2.20462},
		{"ten_kilograms_to_pounds", 10, 22.0462},
		{"zero_kilograms_to_pounds", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KilogramsToPounds(tt.input)
			if !almostEqual(got, tt.want) {
				t.Errorf("KilogramsToPounds(%v)=%v, want=%v", tt.input, got, tt.want)
			}
		})
	}

}
