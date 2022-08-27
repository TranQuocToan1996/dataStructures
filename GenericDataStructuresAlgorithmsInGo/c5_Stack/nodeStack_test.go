package stack

import (
	"fmt"
	"testing"
)

func TestNodeStack(t *testing.T) {
	nameStack := Stack[string]{}
	nameStack.Push("Zachary")
	nameStack.Push("Adolf")
	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}

	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s",
			poppedFromStack)
	}
	// Create a stack of integers
	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0)
	if !intStack.IsEmpty() {
		top := intStack.Top()
		fmt.Printf("\nValue on top of intStack is %d", top)
	}
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
	if !intStack.IsEmpty() {
		popFromStack := intStack.Pop()
		fmt.Printf("\nValue popped from intStack is %d", popFromStack)
	}
}
