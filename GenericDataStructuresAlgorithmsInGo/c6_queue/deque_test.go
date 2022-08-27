package c6queue

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	myDeque := Deque[int]{}
	myDeque.InsertFront(5)
	myDeque.InsertBack(10)
	myDeque.InsertFront(2)
	myDeque.InsertBack(12) // 2 5 10 12
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())
	myDeque.RemoveLast()
	myDeque.RemoveFirst()
	fmt.Println("myDeque.First() = ", myDeque.First())
	fmt.Println("myDeque.Last() = ", myDeque.Last())
}
