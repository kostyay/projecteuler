package main

import "fmt"

const (
	// EndNumber: Upper limit of our search
	EndNumber = 1000
)

// The function will get a sum of all multiplies of n to the max value of max
// if ignoreMultiplier!=0 is given if the number is a multiplier of ignoreMultiplier it wont be considered in the result
func getSumOfMultiplies(n, max, ignoreMultiplier int) int {
	sum := 0
	for i := n; i < max; i += n {
		// if ingoreMultiplier is given and the current number is a multiplier of it, skip it for the sum
		if ignoreMultiplier != 0 && i%ignoreMultiplier == 0 {
			continue
		}

		sum += i
	}

	return sum
}

func main() {

	// Get sum of 3 multiplies up to EndNumber which are not multiplies of 5
	result := getSumOfMultiplies(3, EndNumber, 5) + getSumOfMultiplies(5, EndNumber, 0)
	fmt.Printf("Result is %d\n", result)
}
