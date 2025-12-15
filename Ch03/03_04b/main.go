package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is a slice, not an array
	// var colors = []string{"Red", "Green", "Blue"}
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	colors = append(colors, "Purple", "Fuchsia")
	fmt.Println(colors)

	colors = remove(colors, 1)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println("Sorted colors:", colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}