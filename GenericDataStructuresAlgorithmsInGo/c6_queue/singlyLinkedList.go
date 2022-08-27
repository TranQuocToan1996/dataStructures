package c6queue

import (
	"algorithms/GenericDataStructuresAlgorithmsInGo/model"
	"fmt"
)

var (
	ErrIndexOutBoundary = fmt.Errorf("index out of the bound")
)

type SLLNode[T model.Ordered] struct {
	Item T
	next *SLLNode[T]
}

type SLLList[T model.Ordered] struct {
	first       *SLLNode[T]
	numberItems int
}

func (list *SLLList[T]) Append(item T) {
	new := &SLLNode[T]{item, nil}
	if list.first == nil {
		list.first = new
	} else {
		last := list.first
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = new
	}

	list.numberItems++

}

func (list *SLLList[T]) InsertAt(index int, item T) error {
	if index < 0 || index > list.numberItems {
		return ErrIndexOutBoundary
	}

	new := &SLLNode[T]{item, nil}

	// Insert at beginning of SLL
	if index == 0 {
		new.next = list.first
		list.first = new
		list.numberItems++
		return nil
	}

	// Move along SLL to the index
	count := 0
	previous := list.first
	cur := list.first

	for count < index {
		previous = cur
		cur = cur.next
		count++
	}

	// At the index
	new.next = cur
	previous.next = new
	list.numberItems++
	return nil

}

func (list *SLLList[T]) RemoveAt(index int) (T, error) {
	var res T
	if index < 0 || index > list.numberItems {
		return res, ErrIndexOutBoundary
	}

	if index == 0 {
		remove := list.first
		list.first = list.first.next
		list.numberItems--
		return remove.Item, nil
	}

	count := 0
	cur := list.first
	prev := list.first
	for count < index {
		prev = cur
		cur = cur.next
		count++
	}

	prev.next = cur.next
	list.numberItems--
	return cur.Item, nil

}

func (list *SLLList[T]) IndexOf(item T) int {
	cur := list.first

	count := 0
	for i := 0; i < list.numberItems; i++ {
		if cur.Item == item {
			return count
		}
		if cur.next == nil {
			return -1
		}
		cur = cur.next
		count++
	}

	return -1
}

func (list *SLLList[T]) ItemAfter(item T) (res T) {
	cur := list.first
	for i := 0; i < list.numberItems; i++ {
		if cur == nil {
			return
		}
		if cur.Item == item {
			res = cur.next.Item
			break
		}
		cur = cur.next
	}

	return

}

func (list *SLLList[T]) Items() []T {
	res := []T{}
	cur := list.first
	for i := 0; i < list.numberItems; i++ {
		res = append(res, cur.Item)
		cur = cur.next
	}
	return res
}

func (list *SLLList[T]) First() *SLLNode[T] {
	return list.first
}

func (list *SLLList[T]) Size() int {
	return list.numberItems
}
