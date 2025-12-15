package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go launched at", t)

	now := time.Now()
	fmt.Println("Current time:", now)

	duration := now.Sub(t)
	fmt.Println("Time since Go launched:", duration)

	fmt.Printf("Go has been around for %.0f hours\n", duration.Hours())

	fmt.Printf(now.Format(time.ANSIC) + "\n")

	tomorrow := now.Add(24 * time.Hour)
	fmt.Println("Same time tomorrow:", tomorrow.Format(time.ANSIC))
}
