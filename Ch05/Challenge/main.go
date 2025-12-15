// Write your answer here, and then test your code.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
    // Your code goes here.
	value1 := convertInputToValue(input1)
	value2 := convertInputToValue(input2)

	switch operation {
	case "add", "+":
		return addValues(value1, value2)
	case "subtract", "-":
		return subtractValues(value1, value2)
	case "multiply", "*":
		return multiplyValues(value1, value2)
	case "divide", "/":
		return divideValues(value1, value2)
	default:
		fmt.Println("Unsupported operation:", operation)
		return 0
	}
}

func convertInputToValue(input string) float64 {
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Values must be numbers", err)
	}
	return value
}

func addValues(value1, value2 float64) float64 {
	sum := value1 + value2
	return sum
}

func subtractValues(value1, value2 float64) float64 {
	diff := value1 - value2
	return diff
}

func multiplyValues(value1, value2 float64) float64 {
	product := value1 * value2
	return product
}

func divideValues(value1, value2 float64) float64 {
	quotient := value1 / value2
	return quotient
}

func main() {
	input1 := "10"
	input2 := "5.5"
	operation := "/" // Change to "subtract", "multiply", or "divide" to test other operations.
	result := calculate(input1, input2, operation)
	fmt.Printf("The result of %s %s %s is: %.2f\n", input1, operation, input2, result)
}