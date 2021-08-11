/*

func main() {
	array := NewArray(5)
	dump(array)
	array.Add(0, 2)
	dump(array)
	array.AddFirst(1)
	dump(array)
	array.AddFirst(0)
	dump(array)
	array.AddLast(3)
	dump(array)
	array.AddLast(4)
	dump(array)
	array.AddLast(5)
	dump(array)
	array.Remove(2)
	dump(array)
	array.RemoveLast()
	dump(array)
	array.RemoveFirst()
	dump(array)
	array.RemoveFirst()
	dump(array)
	array.RemoveFirst()
	dump(array)
	fmt.Println(array.Contains(4))
	fmt.Println(array.RemoveLast())

		Array: size = 0, capacity = 5
		[]
		Array: size = 1, capacity = 5
		[2]
		Array: size = 2, capacity = 5
		[1,2]
		Array: size = 3, capacity = 5
		[0,1,2]
		Array: size = 4, capacity = 5
		[0,1,2,3]
		Array: size = 5, capacity = 5
		[0,1,2,3,4]
		Array: size = 6, capacity = 10
		[0,1,2,3,4,5]
		Array: size = 5, capacity = 10
		[0,1,3,4,5]
		Array: size = 4, capacity = 10
		[0,1,3,4]
		Array: size = 3, capacity = 10
		[1,3,4]
		Array: size = 2, capacity = 5
		[3,4]
		Array: size = 1, capacity = 2
		[4]
		true
		4
}

func dump(a *Array) {
	fmt.Println(a)
}

*/

package array

import "testing"

func TestNewArray(t *testing.T) {
	newArray := NewArray(5)
	if newArray.size != 0 || cap(newArray.data) != 5 {
		t.Error(`TestNewArray(5) failed`)
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
