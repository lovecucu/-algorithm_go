package stack

import "testing"

func TestNewArrayStack(t *testing.T) {
	if NewArrayStack(0) != nil {
		t.Error(`TestNewArrayStack failed`)
	}
	stack := NewArrayStack(5)
	if stack == nil || stack.length != 0 || cap(stack.data) != 5 {
		t.Error(`TestNewArrayStack failed`)
	}
}

func TestArrayStackPush(t *testing.T) {
	stack := NewArrayStack(2)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.length != 3 || stack.data[0] != 1 || stack.data[1] != 2 || stack.data[2] != 3 || cap(stack.data) != 4 {
		t.Error(`TestArrayStackPush failed`)
	}
}

func TestArrayStackPop(t *testing.T) {
	stack := NewArrayStack(2)
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
		t.Error(`TestArrayStackPop failed`)
	}
}

func TestArrayStackPeek(t *testing.T) {
	stack := NewArrayStack(2)
	stack.Push(1)
	stack.Push(2)
	e1 := stack.Peek()
	if stack.length != 2 || e1 != 2 {
		t.Error(`TestArrayStackPeek failed`)
	}
	stack.Pop()
	e2 := stack.Peek()
	if stack.length != 1 || e2 != 1 {
		t.Error(`TestArrayStackPeek failed`)
	}
	stack.Pop()
	e3 := stack.Peek()
	if e3 != nil {
		t.Error(`TestArrayStackPeek failed`)
	}
}

func TestArrayStackEmpty(t *testing.T) {
	stack := NewArrayStack(2)
	if !stack.Empty() {
		t.Error(`TestArrayStackEmpty failed`)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Pop()
	if stack.Empty() {
		t.Error(`TestArrayStackEmpty failed`)
	}
	stack.Pop()
	if !stack.Empty() {
		t.Error(`TestListStackEmpty failed`)
	}
}

func TestArrayStackSearch(t *testing.T) {
	stack := NewArrayStack(2)
	stack.Push(1)
	stack.Push(2)
	if stack.Search(2) != 1 || stack.Search(1) != 2 || stack.Search(3) != -1 {
		t.Error(`TestArrayStackPeek failed`)
	}
	stack.Pop()
	stack.Pop()
	if stack.Search(1) != -1 {
		t.Error(`TestArrayStackPeek failed`)
	}
}
