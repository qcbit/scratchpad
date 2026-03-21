package circularqueue

import "testing"

func TestCircularQueue(t *testing.T) {

	t.Run("Test Empty Queue", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		if !cq.isEmpty() {
			t.Fatal("Expected empty queue")
		}
	})

	t.Run("Test Dequeue Empty Queue", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		if cq.deQueue() {
			t.Fatal("Expected deQueue to fail on empty queue")
		}
	})

	t.Run("Test Front and Rear on Empty Queue", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		if cq.Front() != -1 {
			t.Fatal("Expected Front to return -1 on empty queue")
		}
		if cq.Rear() != -1 {
			t.Fatal("Expected Rear to return -1 on empty queue")
		}
	})

	t.Run("Test Enqueue", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		if !cq.enQueue(1) {
			t.Fatal("enQueue 1 failed")
		}
	})

	t.Run("Test Dequeue", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		cq.enQueue(1)
		if !cq.deQueue() {
			t.Fatal("deQueue failed")
		}
		if !cq.isEmpty() {
			t.Fatal("Expected queue to be empty after deQueue")
		}
	})

	t.Run("Test Capacity 3 Expansion", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		if !cq.enQueue(1) {
			t.Fatal("enQueue 1 failed")
		}
		if !cq.enQueue(2) {
			t.Fatal("enQueue 2 failed")
		}
		if !cq.enQueue(3) {
			t.Fatal("enQueue 3 failed")
		}
		if !cq.enQueue(4) {
			t.Fatal("enQueue 4 failed")
		}
	})

	t.Run("Test Front and Rear", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		cq.enQueue(1)
		cq.enQueue(2)
		cq.enQueue(3)
		if cq.Front() != 1 {
			t.Fatal("Expected Front to return 1")
		}
		if cq.Rear() != 3 {
			t.Fatal("Expected Rear to return 3")
		}
	})

	t.Run("Test isFull", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		cq.enQueue(1)
		cq.enQueue(2)
		cq.enQueue(3)
		if !cq.isFull() {
			t.Fatal("Expected queue to be full")
		}
	})

	t.Run("Test Resize", func(t *testing.T) {
		cq := (&MyCircularQueue{}).NewCircularQueue(3)
		cq.enQueue(1)
		cq.enQueue(2)
		cq.enQueue(3)
		cq.enQueue(4) // This should trigger a resize
		if cq.capacity != 6 {
			t.Fatal("Expected capacity to be 6 after resize")
		}
		if cq.Front() != 1 {
			t.Fatal("Expected Front to return 1 after resize")
		}
		if cq.Rear() != 4 {
			t.Fatal("Expected Rear to return 4 after resize")
		}
	})

}