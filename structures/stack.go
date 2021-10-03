package main

import "fmt"

/*
Slice implementation
Push: O(1)
Pop: O(1)
*/

type Stack struct {
	values []int
	top    int
}

// Constructor of the stack
func createStack() *Stack {
	return &Stack{values: make([]int, 0), top: -1}
}

// Add to the stack
func (s *Stack) push(value int) {
	s.values = append(s.values, value)
	s.top += 1
}

// Remove from the stack
func (s *Stack) pop() {
	if s.top == -1 {
		return
	}
	s.values = s.values[:len(s.values)-1]
	s.top -= 1
}

// Utility to print a stack
func printStack(s *Stack) {
	fmt.Println(s.values)
}

// Main entry point
func main() {
	stack := createStack()
	fmt.Println("Stack append")
	stack.push(10)
	stack.push(15)
	stack.push(20)
	fmt.Println(stack.top)
	printStack(stack)

	fmt.Println("Stack removal")
	stack.pop()
	stack.pop()
	stack.pop()
	stack.pop()
	fmt.Println(stack.top)
	printStack(stack)
}
