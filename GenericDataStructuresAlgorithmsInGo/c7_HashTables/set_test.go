package hash

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	set1 := Set[int]{}
	set1.Insert(3)
	set1.Insert(5)
	set1.Insert(7)
	set1.Insert(9)
	set2 := Set[int]{}
	set2.Insert(3)
	set2.Insert(6)
	set2.Insert(8)
	set2.Insert(9)
	set2.Insert(11)
	set2.Delete(11)
	fmt.Println("Items in set2: ", set2.Item())
	fmt.Println("5 in set1: ", set1.In(5))
	fmt.Println("5 in set2: ", set2.In(5))
	fmt.Println("Union of set1 and set2: ", set1.Union(&set2).Item())
	fmt.Println("Intersection of set1 and set2: ",
		set1.Intersection(&set2).Item())
	fmt.Println("Difference of set2 with respect to set1: ",
		set2.Difference(&set1).Item())
	fmt.Println("Size of this difference: ", set1.
		Intersection(&set2).Size())
}
