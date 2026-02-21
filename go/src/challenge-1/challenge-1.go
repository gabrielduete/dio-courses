package main

import "fmt"

func main() {
	kelvin := 100.20

	celsius := kelvin - 273.15

	fmt.Printf("%.2fK in Celsius is %.2fºC\n", kelvin, celsius)
}