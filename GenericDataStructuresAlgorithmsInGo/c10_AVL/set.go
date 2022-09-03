package avl

import (
	"fmt"
	"sync"
)

func NewSet() *Set {
	return &Set{
		container: make(map[float64]struct{}),
	}
}

type Set struct {
	container map[float64]struct{}
	mutex     sync.RWMutex
}

func (c *Set) IsPresent(key float64) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, present := c.container[key]
	return present
}
func (c *Set) Add(key float64) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.container[key] = struct{}{}
}

func (c *Set) Remove(key float64) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	_, present := c.container[key]
	if !present {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *Set) Size() int {
	return len(c.container)
}
