package redblacktree

import "testing"

/*
//
Insertion time for red-black tree: 27.62615ms
Search time for red-black tree: 16.037945ms
Insertion time for AVL tree: 48.315163ms
Search time for AVL tree: 3.914522ms
*/
func TestRedBlackTree(t *testing.T) {
	myTree := NewRedBlackTree[Integer](10)
	myTree.Insert(20)
	myTree.Insert(4)
	myTree.Insert(15)
	myTree.Insert(17)
	myTree.Insert(40)
	myTree.Insert(50)
	myTree.Insert(60)
	myTree.Insert(70)
	myTree.Insert(35)
	myTree.Insert(38)
	myTree.Insert(18)
	myTree.Insert(19)
	myTree.Insert(45)
	myTree.Insert(30)
	myTree.Insert(25)
	ShowTreeGraph(*myTree)
}
