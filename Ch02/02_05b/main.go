package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Println("Rounded sum:", sum)

	fmt.Println("The value of Pi is approximately", math.Pi)

	circleRadius := 15.5
	circumference := 2 * math.Pi * circleRadius
	fmt.Printf("Circumference of circle with radius %.2f is %.2f\n", circleRadius, circumference)
}
