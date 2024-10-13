package middleoflinkedlist

import (
	"testing"

	"github.com/qcbit/scratchpad/go/linkedlist"
)

const ErrorString = `got %v want %v`

func TestMiddleOfLinkedList(t *testing.T) {
	tests := []struct {
		num  []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 3},
		{[]int{-1, -2, -3, -4, -5}, -3},
		{[]int{1, 24, 6, 7, 468, 9, 11, 25, 97, 101}, 9},
	}
	t.Run("Middle of the Linked List", func(t *testing.T) {
		for _, tt := range tests {
			ll := linkedlist.NewLinkedList()
			for _, num := range tt.num {
				ll.Insert(num)
			}
			got := getMiddleOfLinkedList(ll.Head)
			if got != tt.want {
				t.Errorf(ErrorString, got, tt.want)
			}
		}
	})
}
