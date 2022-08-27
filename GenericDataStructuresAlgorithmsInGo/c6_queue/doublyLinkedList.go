package c6queue

import "algorithms/GenericDataStructuresAlgorithmsInGo/model"

type DLLNode[T model.Ordered] struct {
	Item T
	next *DLLNode[T]
	prev *DLLNode[T]
}

type DLLList[T model.Ordered] struct {
	first       *DLLNode[T]
	last        *DLLNode[T]
	numberItems int
}

func (list *DLLList[T]) Append(item T) {
	// Adds item to a new node at the end of the list
	newNode := DLLNode[T]{item, nil, nil}
	if list.first == nil {
		list.first = &newNode
		list.last = list.first
	} else {
		list.last.next = &newNode
		newNode.prev = list.last
		list.last = &newNode
	}
	list.numberItems += 1
}

func (list *DLLList[T]) InsertAt(index int, item T) error {
	// Adds item to a new node at position index in the list
	if index < 0 || index > list.numberItems {
		return ErrIndexOutBoundary
	}
	newNode := DLLNode[T]{item, nil, nil}
	if index == 0 {
		newNode.next = list.first
		if list.first != nil {
			list.first.prev = &newNode
		}
		list.first = &newNode
		list.numberItems += 1
		if list.numberItems == 1 {
			list.last = list.first
		}
		return nil // No error
	}
	node := list.first
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	newNode.next = node
	previous.next = &newNode
	node.prev = &newNode
	newNode.prev = previous
	list.numberItems += 1
	return nil // no error
}

func (list *DLLList[T]) RemoveAt(index int) (T, error) {
	if index < 0 || index > list.numberItems {
		var zero T
		return zero, ErrIndexOutBoundary
	}
	node := list.first
	if index == 0 {
		toRemove := node
		list.first = toRemove.next
		list.numberItems -= 1
		if list.numberItems <= 1 {
			list.last = list.first
		}
		return toRemove.Item, nil
	}
	count := 0
	previous := node
	for count < index {
		previous = node
		count++
		node = node.next
	}
	toRemove := node
	previous.next = toRemove.next
	toRemove.next.prev = previous
	list.numberItems -= 1
	if list.numberItems <= 1 {
		list.last = list.first
	}
	return toRemove.Item, nil
}

func (list *DLLList[T]) IndexOf(item T) int {
	node := list.first
	count := 0
	for {
		if node.Item == item {
			return count
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		count += 1
	}
}

func (list *DLLList[T]) ItemAfter(item T) T {
	// Scan list for the first occurence of item
	node := list.first
	for {
		if node == nil { // item not found
			var zero T
			return zero
		}
		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.next.Item
}
func (list *DLLList[T]) ItemBefore(item T) T {
	// Scan list for the first occurence of item
	node := list.first
	for {
		if node == nil { // item not found
			var zero T
			return zero
		}
		if node.Item == item {
			break
		}
		node = node.next
	}
	return node.prev.Item
}

func (list *DLLList[T]) Items() []T {
	result := []T{}
	node := list.first
	for i := 0; i < list.numberItems; i++ {
		result = append(result, node.Item)
		node = node.next
	}
	return result
}
func (list *DLLList[T]) ReverseItems() []T {
	result := []T{}
	node := list.last
	for {
		if node == nil {
			break
		}
		result = append(result, node.Item)
		node = node.prev
	}
	return result
}
func (list *DLLList[T]) First() *DLLNode[T] {
	return list.first
}
func (list *DLLList[T]) Last() *DLLNode[T] {
	return list.last
}
func (list *DLLList[T]) Size() int {
	return list.numberItems
}
