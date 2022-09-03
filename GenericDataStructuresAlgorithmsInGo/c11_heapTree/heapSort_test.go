package heap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestHeapSort(t *testing.T) {
	const size = 50_000
	slice := []float64{0.0, 2.7, -3.3, 9.6, -13.8, 26.0, 4.9, 2.6,
		5.1, 1.1}
	sorted := HeapSort[float64](slice, true)
	fmt.Println("After heapSort on slice: ", sorted)
	data := make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	start := time.Now()
	largeSorted := HeapSort[float64](data, true)
	elapsed := time.Since(start)
	fmt.Println("Time for heapSort of 50 million floats: ", elapsed)
	if !IsSorted[float64](largeSorted) {
		fmt.Println("largeSorted is not sorted.")
	}
}
