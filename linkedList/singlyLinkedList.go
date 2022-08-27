package main

import "fmt"

// singlyNode of linked list
type singlyNode struct {
	data interface{}
	next *singlyNode
}

type singlyLinkList struct {
	head   *singlyNode // The first node
	length int         // How long the linked list
}

// prepend add the node to the head of linked list
// Pointer receiver to make sure the changing happen in the source value, not in the copy value
func (l *singlyLinkList) prepend(newNode *singlyNode) {
	// Set the head to temp variable
	second := l.head
	// Now assign the head to newNode
	l.head = newNode
	// point to the previous head
	l.head.next = second
	// Add a node increases length
	l.length++
}

// printListData prints the datas of the linked List
func (l singlyLinkList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		// We can minus l.length because the receiver is just the copy one
		l.length--
	}
	fmt.Print("\n")
}

// deleteNodeWithValue find the node that need to delete, actually we just unlinked it.
func (l *singlyLinkList) deleteNodeWithValue(val int) {

	// Handlle runtime error when delete a linked list with no node
	if l.length == 0 {
		fmt.Printf("Empty linked list!")
		return
	}

	// Handle runtime error when deletes head
	if l.head.data == val {
		// the second node will become head
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	// In the single linked list, we only have the next node addr, so that we only compare the value of next node.
	// Loop through the linked list unstill find the next node value == deleted-value
	for previousToDelete.next.data != val {
		// Handle case that point to nil
		if previousToDelete.next.next == nil {
			fmt.Println("The needed-delete-value don't exist in the linked list")
			return
		}
		previousToDelete = previousToDelete.next

	}
	// eg: 1 -> 2 ->3 ->4 . Need to delete is 3, the previousToDelete is 2.
	// so we just assign the node 4 to node 3, so that the node 3 is unlink
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := singlyLinkList{}
	node1 := &singlyNode{data: 1}
	node2 := &singlyNode{data: 2}
	node3 := &singlyNode{data: 3}
	node4 := &singlyNode{data: 4}
	node5 := &singlyNode{data: 5}
	node6 := &singlyNode{data: 6}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printListData()
	mylist.deleteNodeWithValue(3)
	mylist.printListData()

	mylist.deleteNodeWithValue(100)
	mylist.printListData()
	mylist.deleteNodeWithValue(6)
	mylist.printListData()

	emptyLinkedList := singlyLinkList{}
	emptyLinkedList.deleteNodeWithValue(1)
	emptyLinkedList.printListData()

}
