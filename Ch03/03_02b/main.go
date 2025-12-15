package main

import (
	"fmt"
)

func main() {

	anInt := 42
	var p *int = &anInt

	if p == nil {
		fmt.Println("p is nil, allocating memory")
	
	} else {
		fmt.Println("Value of p:", *p)
	}

	val1 := 42.13
	pointer1 := &val1
	*pointer1 = *pointer1 / 31
	fmt.Println("pointer1:", *pointer1)
	fmt.Println("val1:", *pointer1)
	
}
