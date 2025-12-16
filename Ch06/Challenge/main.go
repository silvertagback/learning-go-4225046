// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

type CartItem struct {
	Name  string
	Price float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []CartItem {
    var cart []CartItem
		decoder := json.NewDecoder(strings.NewReader(jsonString))
		_, err := decoder.Token()
		checkError(err)
		var item CartItem
		for decoder.More() {
			err := decoder.Decode(&item)
			checkError(err)
			cart = append(cart, item)
		}
	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// You can use this main() method to test your code.
	// However, the automated tests will ignore it.
}