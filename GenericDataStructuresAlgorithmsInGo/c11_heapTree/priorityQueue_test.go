package heap

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	myQueue := PriorityQueue[string]{}
	myQueue.Push("Helen")
	myQueue.Push("Apollo")
	myQueue.Push("Richard")
	myQueue.Push("Barbara")
	fmt.Println(myQueue)
	myQueue.Pop()
	fmt.Println(myQueue)
	myQueue.Push("Arlene")
	fmt.Println(myQueue)
	myQueue.Pop()
	myQueue.Pop()
	fmt.Println(myQueue)
}
