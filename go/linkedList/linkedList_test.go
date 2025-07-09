package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("Insert 1 into an empty linked list", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(1)
		if ll.Head.Value != 1 {
			t.Errorf("got %v want 1", ll.Head.Value)
		}
	})

	t.Run("Delete an inserted value into empty linked list result in head being nil", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(1)
		ll.Delete(1)
		if ll.Head != nil {
			t.Errorf("Expected linked list's head to be nil, but see %v", ll.Head.Value)
		}
	})

	t.Run("Add 1,2,3 and delete the middle", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(1)
		ll.Insert(2)
		ll.Insert(3)
		ll.Delete(2)
		res := make([]int, 0, 3)
		curr := ll.Head
		for curr != nil {
			res = append(res, curr.Value)
			curr = curr.Next
		}
		want := []int{1, 3}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("got %v res want %v", res, want)
		}
	})

	t.Run("Add 1,2,3 and delete the beginning", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(1)
		ll.Insert(2)
		ll.Insert(3)
		ll.Delete(1)
		res := make([]int, 0, 3)
		curr := ll.Head
		for curr != nil {
			res = append(res, curr.Value)
			curr = curr.Next
		}
		want := []int{2, 3}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("got %v res want %v", res, want)
		}
	})

	t.Run("Add 1,2,3 and delete the end", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(1)
		ll.Insert(2)
		ll.Insert(3)
		ll.Delete(3)
		res := make([]int, 0, 3)
		curr := ll.Head
		for curr != nil {
			res = append(res, curr.Value)
			curr = curr.Next
		}
		want := []int{1, 2}
		if !reflect.DeepEqual(res, want) {
			t.Errorf("got %v res want %v", res, want)
		}
	})

	t.Run("Reverse a linked list", func(t *testing.T) {
		ll := NewLinkedList()
		ll.Insert(1)
		ll.Insert(2)
		ll.Insert(3)
		Reverse(ll)
		want := []int{3, 2, 1}
		for i, v := range want {
			if ll.Head == nil {
				t.Errorf("Expected linked list's head to not be nil, but it is")
				return
			}
			if ll.Head.Value != v {
				t.Errorf("got %v want %v at index %d", ll.Head.Value, v, i)
			}
			ll.Head = ll.Head.Next
		}
	})
}
