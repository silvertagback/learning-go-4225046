package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	go say("Hello from a goroutine!")

	go func(message string) {
		fmt.Println(message)
	}("Hello from an anonymous function!")
	
	say("Hello from main function!")
	time.Sleep(2 * time.Second)
	fmt.Println("All done!")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}