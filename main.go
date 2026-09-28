package main

import "fmt"

func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	result := CelsiusToFahrenheit(100)
	fmt.Println(result)
}
