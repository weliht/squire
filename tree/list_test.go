package tree

import (
	//"fmt"
	"testing"
)

type node struct {
	BaseDNode[node]

	val int
}

func TestDListAppend(t *testing.T) {
	list := &DList[*node]{}

	n1 := &node{}
	list.Append(n1)

	// Head and Tail must be the n1 pointer
	if list.Head != n1 {
		t.Fatalf("list's head isn't the n1, head=%v - append first node", list.Head)
	}

	if list.Tail != n1 {
		t.Fatalf("list's tail isn't the n1, tail=%v - append first node", list.Tail)
	}

	n2 := &node{}
	list.Append(n2)

	// Tail and head's next must be the n2
	// Tail's prev must be n1
	if list.Tail != n2 || list.Head.next != n2 || list.Tail.prev != n1 {
		t.Errorf(
			"mismatch: tail=%v head.next=%v tail.prev=%v - append second node",
			list.Tail, list.Head.next, list.Tail.prev,
		)
	}
}

func TestDListAppendAt(t *testing.T) {
	list := &DList[*node]{}

	n1 := &node{}
	n2 := &node{}
	n3 := &node{}
	list.Append(n1, n2, n3)

	n4 := &node{}

	// move n4 node ahead of n2
	list.AppendAt(n2, n4)

	if n2.next != n4 || n4.prev != n2 {
		t.Fatalf("mismatch: expect=n4 is next of the n2 actual=n2.next=%v n4.prev=%v - append n4 ahead of n2", n2.next, n4.prev)
	}

	list.AppendAt(n3, n1)

	if !(n2 == list.Head && n1 == list.Tail && n3.next == n1 && n1.prev == n3) {
		t.Errorf("mismatch: expect=n2 is head, n1 is tail and n3 is behind of n1")
	}
}

func TestDListRemove(t *testing.T) {
	list := &DList[*node]{}

	n1 := &node{val: 1}
	n2 := &node{val: 2}
	n3 := &node{val: 3}
	list.Append(n1)
	list.Append(n2)
	list.Append(n3)

	list.Remove(n2)

	if !(n1.next == n3 && n3.prev == n1) {
		t.Fatalf("mismatch: expect=n1 is behind of n3 actual=n1.next=%v n3.prev=%v", n1.next, n3.prev)
	}

	// return n2 back to the list
	list.AppendAt(n1, n2)

	// remove n1 to check for list's head is set correctly
	list.Remove(n1)
	if !(list.Head == n2) {
		t.Errorf("mismatch: expect=list's head is n2 actual=head=%v - head removing", list.Head)
	}
}




