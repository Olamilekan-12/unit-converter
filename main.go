package main

import "fmt"

func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	result := CelsiusToFahrenheit(100)
	result2 := FahrenheitToCelsius(212)
	fmt.Println(result)
	fmt.Println(result2)
}
