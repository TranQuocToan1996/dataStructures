package dynamic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCompareFib(t *testing.T) {

	nDynamicTop := 1_000
	nDynamicBot := 5_000
	nRecursive := 40

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		fib := FibonacciTopDown(nDynamicTop)
		elapsed := time.Since(start)
		fmt.Printf("Value of FibonacciTopDown(%v): %v\n", nDynamicTop, fib)
		fmt.Println("Computation time nano: ", int64(elapsed/time.Nanosecond))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		fib := FibonacciBottomUp(nDynamicBot)
		elapsed := time.Since(start)
		fmt.Printf("\nValue of FibonacciBottomUp(%v): %v\n", nDynamicBot, fib)
		fmt.Println("Computation time nano: ", int64(elapsed/time.Nanosecond))
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		start := time.Now()
		fib := FibRecursive(nRecursive)
		elapsed := time.Since(start)
		fmt.Printf("\nValue of FibRecursive(%v): %v\n", nRecursive, fib)
		fmt.Println("Computation time: ", elapsed)
	}()

	wg.Wait()
	fmt.Println("test done")

}
