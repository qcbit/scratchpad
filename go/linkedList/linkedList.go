package linkedlist

import "fmt"

// Node is a node that has a value
// and a pointer to the next node
type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value: value}
}

type LinkedList interface {
	Insert(value int) bool
	Delete(value int) bool
	DeleteAtBeginning(value int) bool
	Search(value int) bool
	Traverse()
}

type MyLinkedList struct {
	head *Node
}

func NewLinkedList() *MyLinkedList {
	return &MyLinkedList{}
}

// Insert inserts a new Node into the linked list
func (ll *MyLinkedList) Insert(value int) bool {
	if ll.head == nil {
		ll.head = NewNode(value)
		return true
	}

	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = NewNode(value)
	return true
}

// Search searches if value is in the linked list and returns a boolean
func (ll *MyLinkedList) Search(value int) bool {
	if ll.head == nil {
		return false
	}
	curr := ll.head
	for curr.next != nil {
		if curr.value == value {
			return true
		}
		curr = curr.next
	}
	return false
}

func (ll *MyLinkedList) Traverse() {
	if ll.head == nil {
		return
	}
	curr := ll.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func (ll *MyLinkedList) Delete(value int) bool {
	if ll.head == nil {
		return false
	}

	if ll.head.value == value {
		ll.head = ll.head.next
		return true
	}

	prev := ll.head
	curr := ll.head.next
	for curr != nil {
		if curr.value == value {
			prev.next = curr.next
			return true
		}
		prev = curr
		curr = curr.next
	}

	return false
}
