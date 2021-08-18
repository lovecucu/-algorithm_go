package queue

import "testing"

func TestNewArrayQueue(t *testing.T) {
	if NewArrayQueue(0) != nil {
		t.Error(`TestNewArrayQueue failed`)
	}
	queue := NewArrayQueue(5)
	if queue == nil || queue.length != 0 || cap(queue.data) != 5 {
		t.Error(`TestNewArrayQueue failed`)
	}
}

func TestArrayQueuePush(t *testing.T) {
	queue := NewArrayQueue(2)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.length != 3 || queue.data[0] != 1 || queue.data[1] != 2 || queue.data[2] != 3 || cap(queue.data) != 4 {
		t.Error(`TestArrayQueuePush failed`)
	}
}

func TestArrayQueuePop(t *testing.T) {
	queue := NewArrayQueue(2)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	e1 := queue.Pop()
	e2 := queue.Pop()
	e3 := queue.Pop()
	e4 := queue.Pop()
	e5 := queue.Pop()
	if e1 != 1 || e2 != 2 || e3 != 3 || e4 != 4 || e5 != nil {
		t.Error(`TestArrayQueuePop failed`)
	}
}

func TestArrayQueuePeek(t *testing.T) {
	queue := NewArrayQueue(2)
	queue.Push(1)
	queue.Push(2)
	e1 := queue.Peek()
	if queue.length != 2 || e1 != 1 {
		t.Error(`TestArrayQueuePeek failed`)
	}
	queue.Pop()
	e2 := queue.Peek()
	if queue.length != 1 || e2 != 2 {
		t.Error(`TestArrayQueuePeek failed`)
	}
	queue.Pop()
	e3 := queue.Peek()
	if e3 != nil {
		t.Error(`TestArrayQueuePeek failed`)
	}
}

func TestArrayQueueEmpty(t *testing.T) {
	queue := NewArrayQueue(2)
	if !queue.Empty() {
		t.Error(`TestArrayQueueEmpty failed`)
	}

	queue.Push(1)
	queue.Push(2)
	queue.Pop()
	if queue.Empty() {
		t.Error(`TestArrayQueueEmpty failed`)
	}
	queue.Pop()
	if !queue.Empty() {
		t.Error(`TestListQueueEmpty failed`)
	}
}

func TestArrayQueueSearch(t *testing.T) {
	queue := NewArrayQueue(2)
	queue.Push(1)
	queue.Push(2)
	if queue.Search(2) != 2 || queue.Search(1) != 1 || queue.Search(3) != -1 {
		t.Error(`TestArrayQueueSearch failed`)
	}
	queue.Pop()
	queue.Pop()
	if queue.Search(1) != -1 {
		t.Error(`TestArrayQueueSearch failed`)
	}
}
