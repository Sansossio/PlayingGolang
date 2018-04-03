// Package
package main

// Imports
import (
	"fmt"
	MyStack "Stack/mystack"
)
// Properties
const maxValues = 10
// Main
func main() {
	// Instance
	myStack := MyStack.CreateStack()
	// Push
	for i := maxValues; i > 0; i-- {
		myStack.Push(i)
	}
	// Peek
	// Value
	value := myStack.Peek()
	// Print
	fmt.Println(value)
	// Pop
	for i := maxValues; i > 0; i-- {
		// Value
		value := myStack.Pop()
		// Print
		fmt.Println(value)
	}
}