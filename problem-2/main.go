package main

import "fmt"

const (
	MaxNumber = 4000000
)

// Calculate next fib value given a and b
func getNextFib(a, b int) int {
	return a + b
}

func main() {
	a, b := 1, 2
	sum := 0

	for b < MaxNumber {
		// If b is even add it to the sum
		if b%2 == 0 {
			sum += b
		}

		// Save old b (it will be a in next sequence)
		oldB := b
		// Calculate new B
		b = getNextFib(a, oldB)
		// Update A
		a = oldB
	}

	fmt.Printf("Result: %d", sum)
}
