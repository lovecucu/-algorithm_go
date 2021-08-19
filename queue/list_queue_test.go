package queue

import "testing"

func TestNewListQueue(t *testing.T) {
	queue := NewListQueue()
	if queue.length != 0 || queue.tail != nil || queue.head != nil {
		t.Error(`TestNewListQueue failed`)
	}
}

func TestListQueuePush(t *testing.T) {
	queue := NewListQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	if queue.length != 3 || queue.head.Value != 1 || queue.head.Next.Value != 2 || queue.head.Next.Next.Value != 3 || queue.tail.Value != 3 {
		t.Error(`TestListQueuePush failed`)
	}
}

func TestListQueuePop(t *testing.T) {
	queue := NewListQueue()
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
		t.Error(`TestListQueuePop failed`)
	}
}

func TestListQueuePeek(t *testing.T) {
	queue := NewListQueue()
	queue.Push(1)
	queue.Push(2)
	e1 := queue.Peek()
	if queue.length != 2 || e1 != 1 {
		t.Error(`TestListQueuePeek failed`)
	}
	queue.Pop()
	e2 := queue.Peek()
	if queue.length != 1 || e2 != 2 {
		t.Error(`TestListQueuePeek failed`)
	}
	queue.Pop()
	e3 := queue.Peek()
	if e3 != nil {
		t.Error(`TestListQueuePeek failed`)
	}
}

func TestListQueueEmpty(t *testing.T) {
	queue := NewListQueue()
	if !queue.Empty() {
		t.Error(`TestListQueueEmpty failed`)
	}

	queue.Push(1)
	queue.Push(2)
	queue.Pop()
	if queue.Empty() {
		t.Error(`TestListQueueEmpty failed`)
	}
	queue.Pop()
	if !queue.Empty() {
		t.Error(`TestListQueueEmpty failed`)
	}
}

func TestListQueueSearch(t *testing.T) {
	queue := NewListQueue()
	queue.Push(1)
	queue.Push(2)
	if queue.Search(2) != 2 || queue.Search(1) != 1 || queue.Search(3) != -1 {
		t.Error(`TestListQueueSearch failed`)
	}
	queue.Pop()
	queue.Pop()
	if queue.Search(1) != -1 {
		t.Error(`TestListQueueSearch failed`)
	}
}
