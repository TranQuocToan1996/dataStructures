package stack

import (
	"fmt"
	"testing"
	"time"
)

// Size 100_000 at MSI GE60 2PL
// Time for 10 million Push() operations on nodeStack:  32.5274ms
// Time for 10 million Pop() operations on nodeStack:  2.0007ms
// Time for 10 million Push() operations on sliceStack:  16.5253ms
// Time for 10 million Pop() operations on sliceStack:  1.5304ms

const size = 100_000

func TestBenchMark(t *testing.T) {
	nodeStack := StackNode[int]{}
	sliceStack := Stack[int]{}
	// Benchmark nodeStack
	start := time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Push(i)
	}
	elapsed := time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ",
		elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ",
		elapsed)
	// Benchmark sliceStack
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Push(i)
	}

	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)
}
