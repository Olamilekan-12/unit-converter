package main

import "testing"

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
			want := tt.want
			if got != want {
				t.Errorf("CelsiusToFahrenheit(%v) = %v, want=%v", tt.input, got, tt.want)
			}
		})

	}

}
