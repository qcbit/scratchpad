package circularqueue

type CircularQueue interface {
	MyCircularQueue()
	Front() int
	Rear() int
	enQueue(v int) bool
	deQueue() bool
	isEmpty() bool
	isFull() bool
}

type MyCircularQueue struct {
	q        []int
	capacity int
	size     int
	head     int
	tail     int
}

func (cq *MyCircularQueue) NewCircularQueue(k int) *MyCircularQueue {
	ary := make([]int, k)
	for i := range k {
		ary[i] = -1 // depends range of values
	}
	return &MyCircularQueue{
		q:        ary,
		capacity: k,
		head:     -1, // depends on constraints
		tail:     -1, // depends on constraints
	}
}

func (cq *MyCircularQueue) enQueue(v int) bool {
	if cq.size == 0 {
		cq.q[0] = v
		cq.head++
		cq.tail++
		cq.size++
	} else {
		if cq.size == cq.capacity {
			cq.Resize(2 * cq.capacity)
		}
		cq.tail = (cq.tail + 1) % cq.capacity
		cq.q[cq.tail] = v
		cq.size++
	}
	return true
}

func (cq *MyCircularQueue) deQueue() bool {
	if cq.size == 0 {
		return false
	} else if cq.size == 1 {
		cq.q[cq.head] = -1
		cq.head = -1
		cq.tail = -1
		cq.size = 0
	} else {
		cq.q[cq.head] = -1
		cq.head = (cq.head + 1) % cq.capacity
		cq.size--
	}
	return true
}

func (cq *MyCircularQueue) Front() int {
	if cq.size == 0 {
		return -1
	}
	return cq.q[cq.head]
}

func (cq *MyCircularQueue) Rear() int {
	if cq.size == 0 {
		return -1
	}
	return cq.q[cq.tail]
}

func (cq *MyCircularQueue) isEmpty() bool {
	return cq.size == 0
}

func (cq *MyCircularQueue) isFull() bool {
	return cq.size == cq.capacity
}

func (cq *MyCircularQueue) Resize(newCapacity int) {
	ary := make([]int, newCapacity)
	for i := range newCapacity {
		ary[i] = -1
	}
	var next int
	for cq.head != cq.tail {
		ary[next] = cq.q[cq.head]
		next++
		cq.head = (cq.head + 1) % cq.capacity
	}
	ary[next] = cq.q[cq.head] // copy tail
	cq.q = ary
	cq.head = 0
	cq.tail = next - 1
	cq.capacity = newCapacity
}
