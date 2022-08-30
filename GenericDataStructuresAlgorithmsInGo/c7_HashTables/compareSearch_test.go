package hash

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
)

var (
	mapCollection   map[string]string
	sliceCollection []string
)

func IsPresent(word string, sliceCollection []string) bool {
	for i := 0; i < len(sliceCollection); i++ {
		if sliceCollection[i] == word {
			return true
		}
	}
	return false
}

func IsPresentBinarySearch(word string, sliceCollection []string) bool {
	// The slice collection is sorted
	low := 0
	high := len(sliceCollection) - 1
	for low <= high {
		median := (low + high) / 2
		if sliceCollection[median] < word {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(sliceCollection) || sliceCollection[low] != word {
		return false
	}
	return true
}

func TestCompareMapAndSlice(t *testing.T) {
	file, err := os.Open("words.txt") // Change this if needed
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()
	// Fill mapCollection and sliceConnection with words
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	mapCollection = make(map[string]string)
	sliceCollection = make([]string, 1)
	var words []string
	for scanner.Scan() {
		word := scanner.Text()
		words = append(words, word)
		mapCollection[word] = word
		sliceCollection = append(sliceCollection, word)
	}
	// Benchmark time to test for presence of each word in mapCollection
	start := time.Now()
	for i := 0; i < len(words); i++ {
		_, present := mapCollection[words[i]]
		if !present {
			fmt.Println("Word not found in mapCollectio0n")
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Number of words in mapCollection: ", len(mapCollection))
	fmt.Println("\nTime to test words in mapCollection: ", elapsed)
	sort.Strings(sliceCollection)
	// Benchmark time to test for presence of each word in sliceCollection
	start = time.Now()
	for i := 0; i < len(sliceCollection); i++ {
		if !IsPresent(sliceCollection[i], sliceCollection) {
			fmt.Println("Word not found in mapCollection")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test words in sliceCollection: ", elapsed)
	// Benchmark time to test for presence of each word in sorted
	// sliceCollection
	start = time.Now()
	for i := 0; i < len(sliceCollection); i++ {
		if !IsPresentBinarySearch(sliceCollection[i], sliceCollection) {
			fmt.Println("Word not found in mapCollectio0n")
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test words in sorted sliceCollection: ", elapsed)
}

// Benchmark test begins to test words:  100000
// Time to test all words in myTable:  59.8845219s
// Time to test words in mapCollection:  10.9828ms
func TestCompareMapAndPureHashTables(t *testing.T) {
	size := 100_000
	myTable := NewTable()
	mapCollection := make(map[string]string)
	words := []string{}
	for i := 0; i < size; i++ {
		word := strconv.Itoa(i)
		words = append(words, word)
		myTable.Insert(word)
		mapCollection[word] = ""
	}
	fmt.Println("Benchmark test begins to test words: ", length)
	start := time.Now()
	for i := 0; i < length; i++ {
		if myTable.IsPresent(words[i]) == false {
			fmt.Println("Word not found in table: ", words[i])
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time to test all words in myTable: ", elapsed)
	start = time.Now()
	for i := 0; i < len(mapCollection); i++ {
		_, present := mapCollection[words[i]]
		if !present {
			fmt.Println("Word not found in mapCollection: ", words[i])
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Time to test words in mapCollection: ", elapsed)
}
