package c6queue

import (
	"fmt"
	"testing"
)

func TestSLL(t *testing.T) {
	cars := SLLList[string]{}
	cars.Append("Honda")
	cars.InsertAt(0, "Nissan")
	cars.InsertAt(0, "Chevy")
	cars.InsertAt(1, "Ford")
	cars.InsertAt(1, "Tesla")
	cars.InsertAt(0, "Audi")
	cars.InsertAt(2, "Volkswagon")
	cars.Append("Volvo")
	fmt.Println(cars.Items())
	fmt.Println("Index of Tesla: ", cars.IndexOf("Tesla"))
	cars.RemoveAt(0)
	car, _ := cars.RemoveAt(3)
	fmt.Println("car removed is: ", car)
	fmt.Println(cars.Items())
	cars.RemoveAt(cars.Size() - 1)
	fmt.Println(cars.Items())
	cars.Append("Lexus")
	fmt.Println(cars.Items())
	fmt.Println("First car in the list is: ", cars.First().Item)
	fmt.Println("Last car in the list is: ", cars.Items()[cars.Size()-1])
}
