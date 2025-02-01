// LeetCode 141
package linkedlistcycle

import "github.com/qcbit/scratchpad/go/linkedlist"

func detectCycle(head *linkedlist.Node) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
