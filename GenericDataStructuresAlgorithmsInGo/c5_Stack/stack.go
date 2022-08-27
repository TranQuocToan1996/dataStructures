package stack

type Ordered interface {
	~float64 | ~int | ~string
}

type Stack[T Ordered] struct {
	items []T
}

func Zero[T Ordered]() T {
	var result T
	return result
}

func (s *Stack[T]) Push(item T) {
	if Zero[T]() != item {
		s.items = append(s.items, item)
	}
}

func (s *Stack[T]) Pop() T {
	leng := len(s.items)
	if leng > 0 {
		pop := s.items[leng-1]
		s.items = s.items[:leng-1]
		return pop
	} else {
		return Zero[T]()
	}
}

func (s *Stack[T]) Top() T {
	leng := len(s.items)
	if leng > 0 {
		return s.items[leng-1]
	} else {
		return Zero[T]()
	}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
