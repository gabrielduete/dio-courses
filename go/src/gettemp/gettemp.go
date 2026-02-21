package main

import "fmt"

const ebullitionF float64 = 212.0

func main() {
	// var tempF float64 = ebullitionF
	tempF := ebullitionF

	// var tempC float64 = (tempF - 32) * 5 / 9
	tempC := (tempF - 32) * 5 / 9

	fmt.Println("Temp ebullition water in Fahrenheit", tempF)
	fmt.Println("Temp ebullition water in Celsius", tempC)
}