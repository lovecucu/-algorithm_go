package array

import "testing"

func TestNewArray(t *testing.T) {
	newArray := NewArray(5)
	if newArray.size != 0 || cap(newArray.data) != 5 {
		t.Error(`TestNewArray failed`)
	}
}

func TestAdd(t *testing.T) {
	newArray := NewArray(5)
	newArray.Add(0, 10)
	if newArray.data[0] != 10 {
		t.Error(`TestAdd failed`)
	}
}

func TestGet(t *testing.T) {
	newArray := NewArray(5)
	newArray.Add(0, 10)
	newArray.Add(0, 9)
	newArray.Add(1, 8)
	if newArray.Get(0) != 9 || newArray.Get(1) != 8 || newArray.Get(2) != 10 || newArray.size != 3 {
		t.Error(`TestAdd failed`)
	}
}

func TestAddFirst(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddFirst(10)
	newArray.AddFirst(9)
	newArray.AddFirst(8)
	if newArray.Get(0) != 8 || newArray.Get(1) != 9 || newArray.Get(2) != 10 || newArray.size != 3 {
		t.Error(`TestAddFirst failed`)
	}
}

func TestAddLast(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(10)
	newArray.AddLast(9)
	newArray.AddLast(8)
	if newArray.Get(0) != 10 || newArray.Get(1) != 9 || newArray.Get(2) != 8 || newArray.size != 3 {
		t.Error(`TestAddLast failed`)
	}
}

func TestRemove(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(10)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.Remove(1)
	if newArray.Get(0) != 10 || newArray.Get(1) != 8 || newArray.size != 2 {
		t.Error(`TestRemove failed`)
	}
}

func TestRemoveFirst(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(10)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.RemoveLast()
	if newArray.Get(0) != 10 || newArray.Get(1) != 9 || newArray.size != 2 {
		t.Error(`TestRemoveFirst failed`)
	}
}

func TestRemoveLast(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(10)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.RemoveFirst()
	if newArray.Get(0) != 9 || newArray.Get(1) != 8 || newArray.size != 2 {
		t.Error(`TestRemoveLast failed`)
	}
}

func TestFind(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.AddLast(8)
	if newArray.Find(8) != 1 {
		t.Error(`TestFind failed`)
	}
}

func TestFindAll(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.AddLast(8)
	idxs := newArray.FindAll(8)
	if len(idxs) != 2 || idxs[0] != 1 || idxs[1] != 2 {
		t.Error(`TestFindAll failed`)
	}
}

func TestContains(t *testing.T) {
	newArray := NewArray(5)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.AddLast(8)
	if newArray.Contains(8) == false || newArray.Contains(10) == true {
		t.Error(`TestContains failed`)
	}
}

func TestResize(t *testing.T) {
	newArray := NewArray(2)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.AddLast(8)
	newArray.AddLast(7)
	if len(newArray.data) != 4 {
		t.Error(`TestResize failed, cap() = `, len(newArray.data))
	}
	newArray.RemoveFirst()
	newArray.RemoveFirst()
	newArray.RemoveFirst()
	if len(newArray.data) != 2 {
		t.Error(`TestResize failed, cap() = `, len(newArray.data))
	}
}

func TestSet(t *testing.T) {
	newArray := NewArray(2)
	newArray.AddLast(9)
	newArray.AddLast(8)
	newArray.Set(1, 2)
	if newArray.Get(0) != 9 || newArray.Get(1) != 2 {
		t.Error(`TestSet failed`)
	}
}

func TestSize(t *testing.T) {
	newArray := NewArray(2)
	if newArray.Size() != 0 {
		t.Error(`TestSize failed`)
	}
	newArray.AddFirst(10)
	newArray.AddFirst(9)
	newArray.AddFirst(8)
	if newArray.Size() != 3 {
		t.Error(`TestSize failed`)
	}
}

func TestIsEmpty(t *testing.T) {
	newArray := NewArray(2)
	if newArray.IsEmpty() == false {
		t.Error(`TestIsEmpty failed`)
	}
	newArray.AddFirst(10)
	if newArray.IsEmpty() == true {
		t.Error(`TestIsEmpty failed`)
	}
}

func TestString(t *testing.T) {
	want := "Array: size = 3, capacity = 4\n[1,2,3]"
	newArray := NewArray(2)
	newArray.Add(0, 1)
	newArray.Add(1, 2)
	newArray.Add(2, 3)
	if newArray.String() != want {
		t.Error(`TestString failed`)
	}
}
