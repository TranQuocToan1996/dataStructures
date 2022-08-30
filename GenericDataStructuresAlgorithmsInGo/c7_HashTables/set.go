package hash

import (
	"algorithms/GenericDataStructuresAlgorithmsInGo/model"
	"sync"
)

type Set[T model.Ordered] struct {
	items map[T]bool
	mutex sync.Mutex
}

func (set *Set[T]) Insert(item T) {
	if set.items == nil {
		set.items = make(map[T]bool)
	}

	_, ok := set.items[item]
	if !ok {
		set.mutex.Lock()
		set.items[item] = true
		set.mutex.Unlock()
	}
}

func (set *Set[T]) Delete(item T) {
	if set.items == nil {
		return
	}

	_, ok := set.items[item]
	if !ok {
		set.mutex.Lock()
		delete(set.items, item)
		set.mutex.Unlock()
	}
}

func (set *Set[T]) In(item T) bool {
	_, present := set.items[item]
	return present
}

func (set *Set[T]) Item() []T {
	items := []T{}

	for key := range set.items {
		items = append(items, key)
	}

	return items
}

func (set *Set[T]) Size() int {
	return len(set.items)
}

func (set *Set[T]) Union(anotherSet *Set[T]) *Set[T] {
	res := &Set[T]{
		items: anotherSet.items,
	}

	for key := range set.items {
		_, present := res.items[key]
		if !present {
			res.mutex.Lock()
			res.items[key] = true
			res.mutex.Unlock()
		}
	}

	return res
}

func (set *Set[T]) Intersection(set2 *Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for i := range set2.items {
		_, present := set.items[i]
		if present {
			result.items[i] = true
		}
	}
	return &result
}

func (set *Set[T]) Difference(set2 *Set[T]) *Set[T] {
	result := Set[T]{}
	result.items = make(map[T]bool)
	for i := range set.items {
		_, present := set2.items[i]
		if !present {
			result.items[i] = true
		}
	}
	return &result
}

func (set *Set[T]) Subset(set2 *Set[T]) bool {
	for i := range set.items {
		_, present := set2.items[i]
		if !present {
			return false
		}
	}
	return true
}
