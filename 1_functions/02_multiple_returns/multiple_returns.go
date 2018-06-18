package main

import (
	"fmt"
	"math"
)

func main() {

	d, e := divrem(5, 3)

	fmt.Printf("Quotient: %d, Remainder: %d \n", d, e)

	//named return

	diam, circum := circle(100)

	fmt.Printf("Diameter: %.2f Circumference: %.2f \n", diam, circum)

}

func divrem(i, j int) (int, int) {

	quotient := i / j
	modulus := i % j

	return quotient, modulus

}

func circle(radius float32) (diameter, circumf float32) {

	diameter = radius * 2
	circumf = 2 * math.Pi * radius

	return

}
