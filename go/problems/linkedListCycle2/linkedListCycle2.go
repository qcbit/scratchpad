// LeetCode 142
package linkedlistcycle2

import "github.com/qcbit/scratchpad/go/linkedlist"

func getCycleLength(head *linkedlist.Node) int {
	hare, tortoise := head, head
	for hare.Next != nil && hare.Next.Next != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next
		if hare == tortoise {
			third := head
			var count int
			for tortoise != third {
				tortoise = tortoise.Next
				third = third.Next
				count++
			}
			return count
		}
	}
	return -1
}
