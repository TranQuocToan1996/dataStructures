package avl

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

const (
	Concurrent = 32
)

var (
	max [Concurrent]float64 // Array Holds the maximum value in each AVL tree

	concurrrentSet [Concurrent]AVLTree[Float] // Slice of AVL trees

	size = 1_000 // Test size

	dataSet []float64
)

func BuildConcurrentSet(dataSet []float64) {
	// Use concurrent processing to construct concurrent AVL trees
	var wg sync.WaitGroup
	sort.Float64s(dataSet)
	segment := len(dataSet) / Concurrent
	for treeNumber := 0; treeNumber < Concurrent; treeNumber++ {
		wg.Add(1)
		go func(treeNumber int) {
			defer wg.Done()
			startVal := segment * treeNumber
			for j := startVal; j < startVal+segment; j++ {
				concurrrentSet[treeNumber].Insert(Float(dataSet[j]))
			}
			max[treeNumber] = dataSet[startVal+segment-1]
		}(treeNumber)
	}

	wg.Wait()

	max[Concurrent-1] = 100
}

func IsPresent(val float64) bool {
	// Determine which AVL tree val is in
	treeNumber := 0
	for ; treeNumber < len(max); treeNumber++ {
		if val <= max[treeNumber] {
			break
		}
	}
	return concurrrentSet[treeNumber].Search(Float(val))
}

func TestCompare(t *testing.T) {
	defer func() {
		a := recover()
		if a != nil {
			fmt.Println(a)
		}
	}()
	mySet := NewSet()
	dataSet = make([]float64, size)
	for i := 0; i < size; i++ {
		dataSet[i] = 100.0 * rand.Float64()
	}
	// Time construction of Set
	start := time.Now()
	for i := 0; i < size; i++ {
		mySet.Add(dataSet[i])
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime to build Set with %d numbers: %s", size, elapsed)
	// Time to test the presence of all numbers in dataSet
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !mySet.IsPresent(dataSet[i]) {
			fmt.Printf("%f not present", dataSet[i])
		}
	}

	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in Set: %s", elapsed)
	avlSet := AVLTree[Float]{nil, 0}
	// Time construction of avlSet
	start = time.Now()
	for i := 0; i < size; i++ {
		avlSet.Insert(Float(dataSet[i]))
	}
	elapsed = time.Since(start)
	fmt.Printf("\n\nTime to build avlSet with %d numbers: %s", size, elapsed)
	// Time to test the presence of all numbers in avlSet
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !mySet.IsPresent(dataSet[i]) {
			fmt.Printf("%f not present", dataSet[0])
		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in avlSet: %s", elapsed)
	// Use concurrent processing to construct concurrent avl trees
	start = time.Now()
	BuildConcurrentSet(dataSet)
	elapsed = time.Since(start)
	fmt.Printf("\n\nTime to build concurrent (%d) avlSet with %d numbers: %s", Concurrent, size, elapsed)
	// Test every number in dataSet against the concurrent set
	start = time.Now()
	for i := 0; i < len(dataSet); i++ {
		if !IsPresent(dataSet[i]) {
			fmt.Printf("%f not present", dataSet[i])

		}
	}
	elapsed = time.Since(start)
	fmt.Printf("\nTime to test the presence of all numbers in concurrent (%d) avlSet: %s", Concurrent, elapsed)
}
