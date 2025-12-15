package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	poodle := Dog{
		Breed:  "Poodle",
		Weight: 30,
	}
	fmt.Println("Dog:", poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
}

type Dog struct {
	Breed string
	Weight int
}