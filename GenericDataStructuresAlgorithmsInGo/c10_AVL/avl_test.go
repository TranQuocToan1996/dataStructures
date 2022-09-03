package avl

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Integer int

func (num Integer) String() string {
	return fmt.Sprintf("%d", num)
}
func TestAVL_Insert(t *testing.T) {
	myTree := AVLTree[Integer]{nil, 0}
	myTree.Insert(10)
	myTree.Insert(15)
	myTree.Insert(5)
	myTree.Insert(3)
	myTree.Insert(6)
	myTree.Insert(13)
	myTree.Insert(20)
	myTree.Insert(2)
	myTree.Insert(4)
	myTree.Insert(8)
	myTree.Insert(12)
	myTree.Insert(1)
	// myTree.Delete(20)
	ShowTreeGraph(myTree)
}

func inorderOperator(val Float) {
	val *= val
	fmt.Println(val.String())
}

// Satisfies OrderedStringer because of ~float64
// Also satisfies OrderedStringer because of String() method below
type Float float64

func (num Float) String() string {
	return fmt.Sprintf("%0.1f", num)
}

func TestAVL_InorderOperator(t *testing.T) {

	size := 10_000

	rand.Seed(time.Now().UnixNano())
	// Generate a random search tree
	randomSearchTree := AVLTree[Float]{nil, 0}
	for i := 0; i < 30; i++ {
		rn := 1.0 + 99.0*rand.Float64()
		randomSearchTree.Insert(Float(rn))
	}
	time.Sleep(3 * time.Second)
	ShowTreeGraph(randomSearchTree)
	randomSearchTree.InOrderTraverse(inorderOperator)
	min := randomSearchTree.Min()
	max := randomSearchTree.Max()
	fmt.Printf("\nMinimum value in tree is %0.1f Maximum value in tree is %0.1f", *min, *max)

	start := time.Now()
	tree := AVLTree[Integer]{nil, 0}

	for val := 0; val < size; val++ {
		tree.Insert(Integer(val))
	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime to build AVL tree with %v nodes: %s. Height of tree: %d", size, elapsed, tree.Height())
	numbers := make([]int, size)
	for i := 0; i < size; i++ {
		numbers[i] = i
	}
	start = time.Now()
	sort.Ints(numbers)
	elapsed = time.Since(start)
	fmt.Printf("\nTime to sort %v ints: %s\n", size, elapsed)

}
