package linked_list

import (
	"fmt"
	"testing"
)

func TestCircularLen(t *testing.T) {
	// Create a new ring of size 5
	r := NewRing(-1)
	if r != nil {
		t.Error(`TestCircularLen failed`)
	}
	r = NewRing(5)
	if r.Len() != 5 {
		t.Error(`TestCircularLen failed`, r.Len())
	}
}

func TestCircularDo(t *testing.T) {
	r := NewRing(5)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
}

func TestCircularMove(t *testing.T) {
	r := new(Ring)
	if r != r.Move(1) {
		t.Error(`TestCircularMove failed`)
	}

	r = NewRing(5)
	n := r.Len()
	for i := n; i > 0; i-- {
		r = r.Prev()
		r.Value = n - i
	}

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

	node := r.Move(2)
	if node.Value != 2 {
		t.Error(`TestCircularMove failed`)
	}

	node = node.Move(4)
	if node.Value != 3 {
		t.Error(`TestCircularMove failed`)
	}

	node = node.Move(-2)
	if node.Value != 0 {
		t.Error(`TestCircularMove failed`)
	}
}

func TestCircularPrevNext(t *testing.T) {
	ring := new(Ring)
	if ring.Prev() != ring {
		t.Error(`TestCircularPrevNext failed`)
	}

	ring = new(Ring)
	if ring.Next() != ring {
		t.Error(`TestCircularPrevNext failed`)
	}
}

func TestCircularLink(t *testing.T) {
	n := 3
	ring1 := NewRing(n)
	for i := 0; i < n; i++ {
		ring1.Value = i
		ring1 = ring1.Next()
	}

	ring1.Link(ring1.Move(2))
	if ring1.Len() != 2 {
		t.Error(`TestCircularPrevNext failed`)
	}

	ring2 := NewRing(n)
	for i := 0; i < n; i++ {
		ring2.Value = i + n
		ring2 = ring2.Next()
	}

	ring1.Do(func(i interface{}) {
		fmt.Println(i)
	})

	ring2.Do(func(i interface{}) {
		fmt.Println(i)
	})

	node := ring1.Link(ring2)
	node.Do(func(i interface{}) {
		fmt.Println(i)
	})
	if node.Len() != 5 || node.Prev().Value != 5 {
		t.Error(`TestCircularPrevNext failed`, node.Prev().Value)
	}
}

func TestCircularUnlink(t *testing.T) {
	n := 3
	r := NewRing(n)
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

	if r.Unlink(-1) != nil {
		t.Error(`TestCircularUnlink failed`)
	}

	r.Unlink(5)

	if r.Len() != (n - 5%n) {
		t.Error(`TestCircularUnlink failed`)
	}
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
}
