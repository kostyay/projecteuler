package main

import (
	"fmt"
	"math"
)

// is the given number prime
func isPrime(n int) bool {
	// skip odd numbers
	if n%2 == 0 {
		return false
	}

	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

// get next prime
func getNextPrime(n int) int {
	if n < 2 {
		return 2
	}

	for {
		n++
		if isPrime(n) {
			return n
		}
	}
}

func main() {
	targetNumber := 600851475143
	prime := 0
	for targetNumber != 1 {
		prime = getNextPrime(prime)
		if targetNumber%prime == 0 {
			targetNumber = targetNumber / prime
			fmt.Println(prime)
			prime = 2
		}
	}

}
