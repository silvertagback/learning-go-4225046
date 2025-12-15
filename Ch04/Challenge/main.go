// Write your answer here, and then test your code.

package main


func main() {
	// You can use this main() function to test your code.
	cart := []CartItem{
		{"Apple", 2.99, 3},
		{"Bread", 1.50, 8},
		{"Milk", 0.49, 12},
	}
	total := calculateTotal(cart)
	println("Total cart value:", total)

	if showExpectedResult {
		println("Expected total: 26.85")
	}
	calculateTotal(cart)
}
// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = false;
const showHints = false;

type CartItem struct{
    Name string
    Price float64
    Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
    // Your code goes here.
		var total float64 = 0.0
		for _, item := range cart {
			total += item.Price * float64(item.Quantity)
		}
		return total
}