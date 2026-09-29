package main

import "fmt"

func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func KilometersToMiles(km float64) float64 {
	oneKmToMile := 0.621371
	return km * oneKmToMile
}

func main() {
	result := CelsiusToFahrenheit(100)
	result2 := FahrenheitToCelsius(212)
	result3 := KilometersToMiles(1)
	result4 := KilometersToMiles(10)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
