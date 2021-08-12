package linked_list

import "testing"

func TestDoublePushFront(t *testing.T) {
	list := &List{}
	list.PushFront(1)
	if list.len != 1 || list.root.next.Value != 1 {
		t.Error(`TestDoublePushFront failed`)
	}

	list.PushFront(2)
	if list.len != 2 || list.root.next.Value != 2 {
		t.Error(`TestDoublePushFront failed`)
	}
}

func TestDoublePushFrontList(t *testing.T) {
	list := New()
	list.PushFront(2)
	list.PushFront(1)
	other := New()
	other.PushFront(4)
	other.PushFront(3)
	list.PushFrontList(other)
	firstEle := list.Front()
	if list.len != 4 || firstEle.Value != 3 || firstEle.next.Value != 4 {
		t.Error(`TestDoublePushFrontList failed`)
	}
}

func TestDoubleRemove(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	removeEle := list.Back()
	if list.Remove(removeEle) != 2 || list.Len() != 1 {
		t.Error(`TestDoubleRemove failed`)
	}
}

func TestDoublePushBackList(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	other := New()
	other.PushBack(3)
	other.PushBack(4)
	list.PushBackList(other)
	lastEle := list.Back()
	if list.len != 4 || lastEle.Value != 4 || lastEle.prev.Value != 3 {
		t.Error(`TestDoublePushBackList failed`, lastEle.Value, lastEle.prev.Value, lastEle.prev.prev.Value)
	}
}

func TestDoublePushBack(t *testing.T) {
	list := &List{}
	list.PushBack(1)
	if list.len != 1 || list.root.next.Value != 1 {
		t.Error(`TestDoublePushBack failed`)
	}

	list.PushBack(2)
	if list.len != 2 || list.root.next.Value != 1 {
		t.Error(`TestDoublePushBack failed`)
	}
}

func TestDoubleBack(t *testing.T) {
	list := New()
	if list.Back() != nil {
		t.Error(`TestDoubleBack failed`)
	}
	list.PushBack(1)
	list.PushBack(2)
	if list.len != 2 || list.Back().Value != 2 {
		t.Error(`TestDoubleBack failed`)
	}
}

func TestDoubleFront(t *testing.T) {
	list := New()
	if list.Front() != nil {
		t.Error(`TestDoubleBack failed`)
	}
	list.PushBack(1)
	list.PushBack(2)
	if list.len != 2 || list.Front().Value != 1 {
		t.Error(`TestDoubleFront failed`)
	}
}

func TestDoubleInsertAfter(t *testing.T) {
	list := &List{}
	list.PushBack(1)
	cur := list.InsertAfter(2, &Element{Value: 1})
	if cur != nil || list.len != 1 {
		t.Error(`TestDoubleInsertAfter failed`)
	}
	mark := list.Front()
	cur = list.InsertAfter(2, mark)
	if cur.Value != 2 || cur.prev.Value != 1 || list.len != 2 {
		t.Error(`TestDoubleInsertAfter failed`)
	}
}

func TestDoubleInsertBefore(t *testing.T) {
	list := &List{}
	list.PushBack(1)
	cur := list.InsertBefore(2, &Element{Value: 1})
	if cur != nil || list.len != 1 {
		t.Error(`TestDoubleInsertBefore failed`)
	}
	mark := list.Front()
	cur = list.InsertBefore(2, mark)
	if cur.Value != 2 || cur.next.Value != 1 || list.len != 2 {
		t.Error(`TestDoubleInsertBefore failed`)
	}
}

func TestDoubleLen(t *testing.T) {
	list := New()
	list.PushBack(1)
	if list.Len() != 1 {
		t.Error(`TestDoubleLen failed`)
	}
	list.PushBack(2)
	if list.Len() != 2 {
		t.Error(`TestDoubleLen failed`)
	}
}

func TestDoubleMoveAfter(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.MoveAfter(list.Front(), &Element{Value: 2})
	if list.Len() != 2 || list.Front().Value != 1 || list.Front().next.Value != 2 {
		t.Error(`TestDoubleMoveAfter failed`)
	}
	list.MoveAfter(list.Front(), list.Back())
	if list.Len() != 2 || list.Front().Value != 2 || list.Front().next.Value != 1 {
		t.Error(`TestDoubleMoveAfter failed`)
	}
}

func TestDoubleMoveBefore(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.MoveBefore(list.Back(), &Element{Value: 2})
	if list.Len() != 2 || list.Front().Value != 1 || list.Front().next.Value != 2 {
		t.Error(`TestDoubleMoveBefore failed`)
	}
	list.MoveBefore(list.Front(), list.Back())
	if list.Len() != 2 || list.Front().Value != 1 || list.Front().next.Value != 2 {
		t.Error(`TestDoubleMoveBefore failed`)
	}
	list.MoveBefore(list.Back(), list.Front())
	if list.Len() != 2 || list.Front().Value != 2 || list.Front().next.Value != 1 {
		t.Error(`TestDoubleMoveBefore failed`)
	}
}

func TestDoubleMoveToBack(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.MoveToBack(&Element{Value: 3})
	if list.Back().Value != 2 {
		t.Error(`TestDoubleMoveToBack failed`)
	}
	firstEle := list.Front()
	list.MoveToBack(firstEle)
	if list.Back() != firstEle {
		t.Error(`TestDoubleMoveToBack failed`)
	}
}

func TestDoubleMoveToFront(t *testing.T) {
	list := New()
	list.PushBack(1)
	list.PushBack(2)
	list.MoveToFront(&Element{Value: 3})
	if list.Front().Value != 1 {
		t.Error(`TestDoubleMoveToFront failed`)
	}
	lastEle := list.Back()
	list.MoveToFront(lastEle)
	if list.Front() != lastEle {
		t.Error(`TestDoubleMoveToFront failed`)
	}
}
