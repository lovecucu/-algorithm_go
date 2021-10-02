package heap

import (
	"container/heap"
)

func HeapSort(input []int) {
	inputLen := len(input)
	for i := 0; i < inputLen; i++ {
		minAjust(input[i:]) // 每次排好第i位的数据
	}
}

func minAjust(input []int) {
	inputLen := len(input)
	if inputLen <= 1 {
		return
	}
	for i := inputLen/2 - 1; i >= 0; i-- { // 小根堆，父结点如果大于子结点，则相反
		if (2*i+1 <= inputLen-1) && (input[i] > input[2*i+1]) { // 左子结点
			input[i], input[2*i+1] = input[2*i+1], input[i]
		}
		if (2*i+2 <= inputLen-1) && (input[i] > input[2*i+2]) { // 右子结点
			input[i], input[2*i+2] = input[2*i+2], input[i]
		}
	}
}

// 大顶堆
type MaxHeapInt []int

func (h MaxHeapInt) Len() int {
	return len(h)
}

func (h MaxHeapInt) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeapInt) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
