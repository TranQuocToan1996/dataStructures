package c9binarysearchtrees

import "algorithms/GenericDataStructuresAlgorithmsInGo/model"

type BinarySearchTree[T model.OrderedStringer] struct {
	Root     *Node[T]
	NumNodes int
}

type Node[T model.OrderedStringer] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

func (bst *BinarySearchTree[T]) Search(val T) bool {
	return search(bst.Root, val)
}

func (bst *BinarySearchTree[T]) Insert(val T) {
	if !bst.Search(val) {
		new := &Node[T]{val, nil, nil}
		if bst.Root == nil {
			bst.Root = new
		} else {
			insertNode(bst.Root, new)
		}
		bst.NumNodes++
	}
}

func (bst *BinarySearchTree[T]) Delete(val T) {
	if bst.Search(val) {
		deleteNode(bst.Root, val)
		bst.NumNodes--
	}
}

// op is operation to be performed when visiting each node
func (bst *BinarySearchTree[T]) InOrderTraverse(op func(T)) {
	inOrderTraverse(bst.Root, op)
}

func (bst *BinarySearchTree[T]) Min() *T {
	cur := bst.Root
	if cur == nil {
		return nil
	}

	for {
		if cur.Left == nil {
			return &cur.Value
		}
		cur = cur.Left
	}
}

func (bst *BinarySearchTree[T]) Max() (val *T, height int) {
	cur := bst.Root
	if cur == nil {
		return nil, 0
	}

	for {
		height++
		if cur.Right == nil {
			return &cur.Value, height
		}
		cur = cur.Right
	}

}

func insertNode[T model.OrderedStringer](node, newNode *Node[T]) {

	if newNode == nil {
		return
	}

	if node == nil {
		node = newNode
		return
	}

	// Left is less than right
	if newNode.Value < node.Value {
		if node.Left == nil {
			node.Left = newNode
		} else {
			// Continue go to the left side BST
			insertNode[T](node.Left, newNode)
		}
	} else {
		if node.Right == nil {
			node.Right = newNode
		} else {
			insertNode[T](node.Right, newNode)
		}
	}
}

func deleteNode[T model.OrderedStringer](node *Node[T], val T) *Node[T] {
	if node == nil {
		return nil
	}

	// Compare val
	if val < node.Value {
		node.Left = deleteNode(node.Left, val)
        return node
	}
	if val > node.Value {
		node.Right = deleteNode(node.Right, val)
        return node
	}

	// Already found/ meet boundary, now check child nodes
	if node.Left == nil && node.Right == nil {
		return nil
	}

	if node.Left == nil {
		return node.Right
	}
	if node.Right == nil {
		return node.Left
	}

	LeftmostRightside := node.Right
	for {
		//find smallest value on the Right side
		if LeftmostRightside != nil && LeftmostRightside.Left != nil {
			LeftmostRightside = LeftmostRightside.Left
		} else {
			break
		}
	}

	node.Value = LeftmostRightside.Value
	node.Right = deleteNode(node.Right, node.Value)
	return node

}

func search[T model.OrderedStringer](n *Node[T], value T) bool {
	if n == nil {
		return false
	}
	if value < n.Value {
		return search(n.Left, value)
	}
	if value > n.Value {
		return search(n.Right, value)
	}
	return true
}

func inOrderTraverse[T model.OrderedStringer](n *Node[T], op func(T)) {
	if n != nil {
		inOrderTraverse(n.Left, op)
		if op != nil {
			op(n.Value)
		}
		inOrderTraverse(n.Right, op)
	}
}
