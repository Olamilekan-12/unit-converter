package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	kgToPounds  = 2.20462
	oneKmToMile = 0.621371
)

func KilogramsToPounds(kg float64) float64 {
	return kg * kgToPounds
}

func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func KilometersToMiles(km float64) float64 {
	return km * oneKmToMile
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: converter <value> <conversion>")
		os.Exit(1)
	}
	parseArg, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot convert value to float")
		os.Exit(1)
	}
	var result float64
	switch args[2] {
	case "c-f":
		result = CelsiusToFahrenheit(parseArg)
		fmt.Printf("CelsiusToFahrenheit(%v) =  %.4f\n", parseArg, result)
	case "f-c":
		result = FahrenheitToCelsius(parseArg)
		fmt.Printf("FahrenheitToCelsius(%v) =  %.4f\n", parseArg, result)
	case "km-mi":
		result = KilometersToMiles(parseArg)
		fmt.Printf("KilometersToMiles(%v) =  %.4f\n", parseArg, result)
	case "kg-lb":
		result = KilogramsToPounds(parseArg)
		fmt.Printf("KilogramsToPounds(%v) =  %.4f\n", parseArg, result)
	default:
		fmt.Fprintln(os.Stderr, "unknown conversion")
		os.Exit(1)
	}
}
