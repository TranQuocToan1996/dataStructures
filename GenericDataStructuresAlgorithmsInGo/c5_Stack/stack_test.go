package stack

import "testing"

func TestStackAlgo(t *testing.T) {

	top := "Toan1"

	nameStack := Stack[string]{}
	nameStack.Push("Toan")
	nameStack.Push("Toan2")
	nameStack.Push(top)

	topVal := nameStack.Top()
	if topVal != top {
		t.Errorf("Top func wrong implementation with %v and expect %v", topVal, top)
	}

	pop := nameStack.Pop()
	if pop != top {
		t.Errorf("Top pop wrong implementation with %v and expect %v", pop, top)
	}

	if nameStack.IsEmpty() {
		t.Errorf("Top empty wrong implementation with %v and expect %v", nameStack.IsEmpty(), true)
	}
	nameStack.Pop()
	nameStack.Pop()

	if !nameStack.IsEmpty() {
		t.Errorf("Top empty wrong implementation with %v and expect %v", !nameStack.IsEmpty(), false)
	}
}
