// Package
package main

// Imports
import "fmt"

// Invert Inver values of two variables
func Invert(var1, var2 *int) {
	// Invert values
	*var1, *var2 = *var2, *var1
}

func main() {
	a := 2
	b := 4
	// Invert(&a, &b)
	fmt.Printf("a: %d, b: %d", a, b)

	var test uint = 4
	test <<= test

	ab := uint64(a)
	ab = ab<<test - 1
	fmt.Println()
	fmt.Printf("Test: %d, a: %d", test, ab)
}
