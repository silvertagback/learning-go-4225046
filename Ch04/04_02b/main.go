package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	switch dayNumber {
	case 0:
		result = "It's Sunday"
	case 1:
		result = "It's Monday"
	case 2:
		result = "It's Tuesday"
	case 3:
		result = "It's Wednesday"
	case 4:
		result = "It's Thursday"
	case 5:
		result = "It's Friday"
	case 6:
		result = "It's Saturday"
	default:
		result = "Unknown day"
	}
	fmt.Println(result)

	switch weekday {	
		case time.Saturday,
				 time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	x := -42
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0:
		fmt.Println("x is positive")
	}	
}
