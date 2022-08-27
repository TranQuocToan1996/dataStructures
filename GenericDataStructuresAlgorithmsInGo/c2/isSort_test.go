package c2

import (
	"fmt"
	"math/rand"
	"runtime"
	"sort"
	"testing"
	"time"
)

// Using concurency

func isSegmentSort(data []float64, left, right int, ch chan<- bool) {
	for i := left + 1; i < right; i++ {
		if data[i-1] > data[i] {
			ch <- false
		}
	}
	ch <- true

}

func isSorted2(data []float64) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func isSorted1(data []float64) bool {
	var data1 []float64
	data1 = make([]float64, len(data)) // Creates a slice of len(data)
	copy(data1, data)                  // Copies data into data1
	sort.Float64s(data1)
	// Compare data and data1
	for i := 0; i < size; i++ {
		if data[i] != data1[i] {
			return false
		}
	}
	return true
}

func isSorted3(data []float64) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))

	for i := 0; i < numSegments; i++ {
		go isSegmentSort(data, i*segmentSize, i*segmentSize+segmentSize, ch)
	}

	var done int

	for {
		select {
		case isSort := <-ch:
			if !isSort {
				return false
			}
			done++
			if done == numSegments {
				return true
			}
		}
	}
}

const size = 1_000_000

var data []float64

func TestCompareSort(t *testing.T) {
	data1 := make([]float64, size)
	data2 := make([]float64, size)

	for i := 0; i < size; i++ {
		data1[i] = 100.0 * rand.Float64()
		data2[i] = float64(2 * i)
	}
	start := time.Now()
	result := isSorted2(data)
	elapsed := time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using sorted2", elapsed)
	start = time.Now()
	result = isSorted2(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using sorted2:", elapsed)
	start = time.Now()
	result = isSorted3(data)
	elapsed = time.Since(start)
	fmt.Println("\nSorted: ", result)
	fmt.Println("elapsed using concurrent sorted3", elapsed)
	start = time.Now()
	result = isSorted3(data2)
	elapsed = time.Since(start)
	fmt.Println("Sorted: ", result)
	fmt.Println("elapsed using concurrent sorted3:", elapsed)

}

var (
	queueChan = make(chan int, 100)
)

func ListeningQueue() {
	go func() {
		fmt.Println("start listening")
		for range queueChan {
			fmt.Println(<-queueChan)
		}
		fmt.Println("end listening")
	}()
}

func TestQueue(t *testing.T) {
	ListeningQueue()
	for i := 0; i < 5; i++ {

		queueChan <- rand.Int()
	}
	fmt.Println("sleep queue")
	time.Sleep(time.Second*3)
	fmt.Println("end TestQueue")
}
