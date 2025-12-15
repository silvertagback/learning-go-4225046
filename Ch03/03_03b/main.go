package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("Colors array:", colors)
	fmt.Println(colors[1])

	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println("Numbers array:", numbers)

	fmt.Println("Number of colors", len(colors))
	fmt.Println("Number of numbers", len(numbers))
}
