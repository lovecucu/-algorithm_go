package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	input := []int{3, 2, 4, 1, 6, 5, 1}
	target := "[1 1 2 3 4 5 6]"
	HeapSort(input)
	if fmt.Sprint(input) != target {
		t.Error(`TestHeapSort failed`, input)
	}
}

func TestMaxHeapInt(t *testing.T) {
	h := make(MaxHeapInt, 0)
	heap.Init(&h)
	heap.Push(&h, 2)
	heap.Push(&h, 4)
	heap.Push(&h, 5)
	heap.Push(&h, 3)
	heap.Push(&h, 1)
	heap.Fix(&h, 6)

	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
	fmt.Println(heap.Pop(&h))
}

func TestPriorityQueue(t *testing.T) {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s \n", item.priority, item.value)
	}
}
