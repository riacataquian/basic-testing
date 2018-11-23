package basic

import "fmt"

// Greet greets the supplied name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Sum returns the sum of the supplied integers.
func Sum(nums ...int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}
