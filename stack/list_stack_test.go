package stack

import "testing"

func TestNewListStack(t *testing.T) {
	stack := NewListStack()
	if stack.length != 0 || stack.tail != nil {
		t.Error(`TestNewListStack failed`)
	}
}

func TestListStackPush(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.length != 3 || stack.tail.Value != 3 || stack.tail.Prev.Value != 2 || stack.tail.Prev.Prev.Value != 1 {
		t.Error(`TestArrayStackPush failed`)
	}
}

func TestListStackPop(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	e1 := stack.Pop()
	e2 := stack.Pop()
	e3 := stack.Pop()
	e4 := stack.Pop()
	e5 := stack.Pop()
	if e1 != 4 || e2 != 3 || e3 != 2 || e4 != 1 || e5 != nil {
		t.Error(`TestListStackPop failed`)
	}
}

func TestListStackPeek(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	e1 := stack.Peek()
	if stack.length != 2 || e1 != 2 {
		t.Error(`TestListStackPeek failed`)
	}
	stack.Pop()
	e2 := stack.Peek()
	if stack.length != 1 || e2 != 1 {
		t.Error(`TestListStackPeek failed`)
	}
	stack.Pop()
	e3 := stack.Peek()
	if e3 != nil {
		t.Error(`TestListStackPeek failed`)
	}
}

func TestListStackEmpty(t *testing.T) {
	stack := NewListStack()
	if !stack.Empty() {
		t.Error(`TestListStackEmpty failed`)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	if stack.Empty() {
		t.Error(`TestListStackEmpty failed`)
	}
	stack.Pop()
	if !stack.Empty() {
		t.Error(`TestListStackEmpty failed`)
	}
}

func TestListStackSearch(t *testing.T) {
	stack := NewListStack()
	stack.Push(1)
	stack.Push(2)
	if stack.Search(2) != 1 || stack.Search(1) != 2 || stack.Search(3) != -1 {
		t.Error(`TestListStackSearch failed`)
	}
	stack.Pop()
	stack.Pop()
	if stack.Search(1) != -1 {
		t.Error(`TestListStackSearch failed`)
	}
}
