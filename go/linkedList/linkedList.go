package linkedlist

import "fmt"

// Node is a node that has a value
// and a pointer to the next node
type Node struct {
	Value int
	Next  *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type LinkedList interface {
	Insert(value int) bool
	Delete(value int) bool
	DeleteAtBeginning(value int) bool
	Search(value int) bool
	Traverse()
}

type MyLinkedList struct {
	Head *Node
}

func NewLinkedList() *MyLinkedList {
	return &MyLinkedList{Head: nil}
}

// Insert inserts a new Node into the linked list
func (ll *MyLinkedList) Insert(value int) bool {
	if ll.Head == nil {
		ll.Head = NewNode(value)
		return true
	}

	curr := ll.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = NewNode(value)
	return true
}

// Search searches if value is in the linked list and returns a boolean
func (ll *MyLinkedList) Search(value int) bool {
	if ll.Head == nil {
		return false
	}
	curr := ll.Head
	for curr.Next != nil {
		if curr.Value == value {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (ll *MyLinkedList) Traverse() {
	if ll.Head == nil {
		return
	}
	curr := ll.Head
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
}

func (ll *MyLinkedList) Delete(value int) bool {
	if ll.Head == nil {
		return false
	}

	if ll.Head.Value == value {
		ll.Head = ll.Head.Next
		return true
	}

	prev := ll.Head
	curr := ll.Head.Next
	for curr != nil {
		if curr.Value == value {
			prev.Next = curr.Next
			return true
		}
		prev = curr
		curr = curr.Next
	}

	return false
}

// reverse reverses a linked list in place
func Reverse(list *MyLinkedList) {
	if list.Head == nil {
		return
	}

	var prev *Node
	curr := list.Head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	list.Head = prev
}
