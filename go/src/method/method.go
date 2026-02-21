package main

import "fmt"

type retangle struct {
	width, height int
}

func (r retangle) area() int {
	return r.width * r.height
}

func (r retangle) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {
	r := retangle{width: 10, height: 5}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter:", r.perimeter())
}