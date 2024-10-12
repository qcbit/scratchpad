package linkedlistcycle2

import (
	"testing"

	"github.com/qcbit/scratchpad/go/linkedlist"
)

func TestLinkedListCycle(t *testing.T) {
	tests := []struct {
		list     []int
		cyclePos int
		want     int
	}{
		{[]int{1, 2}, 0, 0},
		{[]int{1}, -1, -1},
		{[]int{3, 2, 0, -4}, 1, 1},
	}

	for _, tt := range tests {
		ll := linkedlist.NewLinkedList()
		for _, num := range tt.list {
			ll.Insert(num)
		}
		insertCycle(tt.cyclePos, ll.Head)
		got := getCycleLength(ll.Head)
		if got != tt.want {
			t.Errorf("got %v want %v for %v", got, tt.want, tt.list)
		}
	}
}

func insertCycle(pos int, head *linkedlist.Node) {
	if pos < 0 {
		return
	}
	curr := head
	tail := head
	var count int
	for tail.Next != nil {
		tail = tail.Next
		if count < pos {
			curr = curr.Next
			count++
		}
	}
	tail.Next = curr
}
