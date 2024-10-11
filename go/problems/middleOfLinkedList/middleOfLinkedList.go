package middleoflinkedlist

import (
	"github.com/qcbit/scratchpad/go/linkedlist"
)



func getMiddleOfLinkedList(head *linkedlist.Node) int {
	if head == nil {
		return 0
	}

	fast, slow := head, head	
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	if fast.Next != nil {
		slow = slow.Next
	}

	return slow.Value
}