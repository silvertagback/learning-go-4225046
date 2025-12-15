package main

import (
	"fmt"
	"strconv"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum:", floatSum)

	fmt.Println("Math")

	total := float64(intSum) + floatSum
	fmt.Println("Total Sum:", total)

	product := float64(intSum) * floatSum
	fmt.Println("Product:", product)

	val1 := "42"
	val2 := "15.5"
	float1, _ := strconv.ParseFloat(val1, 64)
	float2, _ := strconv.ParseFloat(val2, 64)
	sum := float1 + float2
	fmt.Println("Sum of parsed values:", sum)

}
