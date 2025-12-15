package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
}

func doSomething() {
	fmt.Println("Doing something")
	value1 := 5
	value2 := 10
	value3 := 56
	result := addValues(value1, value2)
	total, count, average := addAllValues(value1, value2, value3, 100, 200)
	fmt.Printf("The sum of %d and %d is %d\n", value1, value2, result)
	fmt.Printf("The total sum is %d\n", total)
	fmt.Printf("The count of values is %d\n", count)
	fmt.Printf("The average is %.2f\n", average)
}

func addValues(a, b int) int {
	return a + b
}

func addAllValues(values ...int) (int, int, float64) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	count := len(values)
	average := float64(sum) / float64(count)
	return sum, count, average
}
