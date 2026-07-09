/*
Double Linked List implementation. It use generic for node type.
Generic T parameter exepect struct pointer type (&Dlist[*MyStruct]{}).

This implementation doesn't define internal Node. Instead it,
we use type that passing into DList T generic param as list Node.
Use BaseDNode struct to implement IDNode methods.

NOTE: We can see such expression as "var nilval T" in following functions.
This is used to determine if node is nil.
"mismatched types ..." error will occur if we will trying to compare some one
filed with nil value.
Type system doesn't figure that value of type T are pointer even
if we use *MyStruct as T parameter.

Ref to similar problems: https://github.com/golang/go/issues/53656
*/
package tree

type IDNode[T any] interface {
	comparable

	Next() T
	Prev() T
	SetNext(T)
	SetPrev(T)
}


type DList[T IDNode[T]] struct {
	Head T
	Tail T
}

// Push the node to tail.
func (l *DList[T]) Append(node ...T) {
	var nilval T
	for _, _node := range node {
		if l.Head == nilval {
			l.Head = _node
			l.Tail = _node
		} else {
			l.Tail.SetNext(_node)
			_node.SetPrev(l.Tail)
			l.Tail = _node
		}
	}
}

func (l *DList[T]) AppendAt(at T, node T) {
	if at.Next() == node {
		// the node is already next of the at
		return 
	}

	var nilval T
	if node.Next() == at {
		// the node is prev of the at
		// just swap them
		at.SetPrev(node.Prev())
		if node.Prev() != nilval {
			node.Prev().SetNext(at)
		}

		node.SetNext(at.Next())
		if at.Next() != nilval {
			at.Next().SetPrev(node)
		}

		if node == l.Head {
			l.Head = at
		}

		if at == l.Tail {
			l.Tail = node
		}

		return
	}

	// the node and the at are far from each other
	// remove the node from the list
	// and append it ahead of the at
	l.Remove(node)

	at.SetNext(node)
	node.SetPrev(at)
	node.SetNext(at.Next())

	if at == l.Tail {
		l.Tail = node
	}
}

/*
Remove node from the list.
Algorithm:
1. Change refs of the next and prev node. Set the next node
as next of prev node and the prev node as prev of the next.

prev <-> current <-> next
prev <-> next

2. Check for the node is list's Head or Tail.
Is Head: set next as new Head
Is Tail: set prev as new Tail

*/
func (l *DList[T]) Remove(node T) {
	var nilval T
	if node.Prev() != nilval { node.Prev().SetNext(node.Next()) }
	if node.Next() != nilval { node.Next().SetPrev(node.Prev()) }

	if node == l.Head {
		l.Head = node.Next()
	}

	if node == l.Tail {
		l.Tail = node.Prev()
	}
}



func (l *DList[T]) Find(cb func(T) bool) T {
	var nilval T
	node := l.Head
	for node != nilval {
		if cb(node) { return node }

		node = node.Next()
	}

	return nilval
}

type BaseDNode[T any] struct {
	next *T
	prev *T
}

//go:inline
func (n *BaseDNode[T]) Next() *T {
	return n.next
}
//go:inline
func (n *BaseDNode[T]) Prev() *T {
	return n.prev
}
//go:inline
func (n *BaseDNode[T]) SetNext(node *T) {
	n.next = node
}
//go:inline
func (n *BaseDNode[T]) SetPrev(node *T) {
	n.prev = node
}
