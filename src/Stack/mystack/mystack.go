// Package
package mystack

// MyStack stack definition
type MyStack struct {
	length int
	top *StackNode
}
// StackNode Nodes
type StackNode struct {
	value interface {}
	previous *StackNode
}
// Push Push into stack
func (st *MyStack) Push(value interface {}) {
	// Properties
	previusValue := st.top
	newValue := StackNode{value, previusValue}
	// Push into stack
	st.top = &newValue
	st.length++
}
// Peek View top value
func (st *MyStack) Peek() interface{} {
	// Comprobe
	if (st.top == nil) {
		return nil
	}
	// Return value
	return st.top.value
}
// Pop extract value
func (st *MyStack) Pop() interface{} {
	// Comprobe
	if (st.top == nil) {
		return nil
	}
	// Get value
	myValue := st.top.value
	previous := st.top.previous
	// Change top
	st.top = previous
	// Return
	return myValue
}
// CreateStack Create new instance with a value
func CreateStack () *MyStack {
	// Result
	result := MyStack{}
	// Properties
	result.length = 0
	result.top = nil
	// Response
	return &result
}

// Main
func main() {

}