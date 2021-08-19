package queue

import (
	"testing"
)

// MyCircularQueue circularQueue = new MyCircularQueue(3); // 设置长度为 3
// circularQueue.enQueue(1);  // 返回 true
// circularQueue.enQueue(2);  // 返回 true
// circularQueue.enQueue(3);  // 返回 true
// circularQueue.enQueue(4);  // 返回 false，队列已满
// circularQueue.Rear();  // 返回 3
// circularQueue.isFull();  // 返回 true
// circularQueue.deQueue();  // 返回 true
// circularQueue.enQueue(4);  // 返回 true
// circularQueue.Rear();  // 返回 4

func TestMyCircularQueue(t *testing.T) {
	circularQueue := Constructor(3)
	if circularQueue.cap != 3 || circularQueue.head != 0 || circularQueue.len != 0 {
		t.Error(`TestMyCircularQueue failed`)
	}
}

func TestMyCircularEnqueue(t *testing.T) {
	circularQueue := Constructor(3)
	if !circularQueue.EnQueue(1) || circularQueue.len != 1 || circularQueue.head != 0 {
		t.Error(`TestMyCircularEnqueue failed`)
	}

	if !circularQueue.EnQueue(2) || circularQueue.len != 2 || circularQueue.head != 0 {
		t.Error(`TestMyCircularEnqueue failed`)
	}

	if !circularQueue.EnQueue(3) || circularQueue.len != 3 || circularQueue.head != 0 {
		t.Error(`TestMyCircularEnqueue failed`)
	}

	if circularQueue.EnQueue(4) || circularQueue.len == 4 {
		t.Error(`TestMyCircularEnqueue failed`)
	}
}

func TestMyCircularDequeue(t *testing.T) {
	circularQueue := new(MyCircularQueue)
	circularQueue.cap = 3
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)
	circularQueue.EnQueue(3)

	if !circularQueue.DeQueue() || circularQueue.len != 2 || circularQueue.head != 1 {
		t.Error(`TestMyCircularDequeue failed`)
	}

	if !circularQueue.DeQueue() || circularQueue.len != 1 || circularQueue.head != 2 {
		t.Error(`TestMyCircularDequeue failed`)
	}

	if !circularQueue.DeQueue() || circularQueue.len != 0 || circularQueue.head != 0 {
		t.Error(`TestMyCircularDequeue failed`)
	}

	if circularQueue.DeQueue() {
		t.Error(`TestMyCircularDequeue failed`)
	}
}

func TestMyCircularFront(t *testing.T) {
	circularQueue := Constructor(3)
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)

	front := circularQueue.Front()
	if front != 1 {
		t.Error(`TestMyCircularFront failed`, circularQueue.head, front)
	}
	circularQueue.DeQueue()
	front = circularQueue.Front()
	if front != 2 {
		t.Error(`TestMyCircularFront failed`, circularQueue.head, front)
	}
	circularQueue.DeQueue()
	front = circularQueue.Front()
	if front != -1 {
		t.Error(`TestMyCircularFront failed`, circularQueue.head, front)
	}
}

func TestMyCircularRear(t *testing.T) {
	circularQueue := Constructor(3)
	circularQueue.EnQueue(1)
	circularQueue.EnQueue(2)

	rear := circularQueue.Rear()
	if rear != 2 {
		t.Error(`TestMyCircularFront failed`)
	}
	circularQueue.DeQueue()
	rear = circularQueue.Rear()
	if rear != 2 {
		t.Error(`TestMyCircularFront failed`, rear)
	}
	circularQueue.DeQueue()
	rear = circularQueue.Rear()
	if rear != -1 {
		t.Error(`TestMyCircularFront failed`)
	}
}
