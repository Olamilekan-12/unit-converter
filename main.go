package main

import "fmt"

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
	result := CelsiusToFahrenheit(100)
	result2 := FahrenheitToCelsius(212)
	result3 := KilometersToMiles(1)
	result4 := KilometersToMiles(10)
	result5 := KilogramsToPounds(1)
	result6 := KilogramsToPounds(10)
	result7 := KilogramsToPounds(0)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)
}
