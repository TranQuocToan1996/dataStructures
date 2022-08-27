package c6queue

import (
	"fmt"
	"testing"
)

func TestPriority(t *testing.T) {
	airlineQueue := NewPriorityQueue[Passenger](3)
	passengers := []Passenger{{"Erika", 3}, {"Robert", 3}, {"Danielle", 3},
		{"Madison", 1}, {"Frederik", 1}, {"James", 2},
		{"Dante", 2}, {"Shelley", 3}}
	fmt.Println("Passsengers: ", passengers)
	for i := 0; i < len(passengers); i++ {
		airlineQueue.Insert(passengers[i], passengers[i].priority)
	}
	fmt.Println("First passenger in line: ", airlineQueue.First())
	airlineQueue.Remove()
	airlineQueue.Remove()
	airlineQueue.Remove()
	fmt.Println("First passenger in line: ", airlineQueue.First())

}
