package stack

type Node[T any] struct {
	value T
	next  *Node[T]
}

type StackNode[T any] struct {
	first *Node[T]
}

func (s *StackNode[T]) Push(item T) {
	s.first = &Node[T]{item, s.first}
}

func (s *StackNode[T]) Top() T {
	return s.first.value
}

func (s *StackNode[T]) Pop() T {
	pop := s.first.value

	s.first = s.first.next

	return pop
}
