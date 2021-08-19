package queue

type ListNode struct {
	Prev, Next *ListNode
	Value      interface{}
}
type ListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

// 初始化栈
func NewListQueue() *ListQueue {
	stack := new(ListQueue)
	return stack
}

// 入栈
func (s *ListQueue) Push(e interface{}) {
	node := &ListNode{Value: e}
	if s.Empty() {
		s.head = node
		s.tail = node
	} else {
		s.tail.Next = node
		node.Prev = s.tail
		s.tail = node
	}
	s.length++
}

// 出栈
func (s *ListQueue) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	item := s.head
	s.length--
	if s.length == 0 {
		s.head = nil
		s.tail = nil
	} else { // head变成下一个结点
		s.head = s.head.Next
	}
	return item.Value
}

// 查看top元素，但不删除
func (s *ListQueue) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	return s.head.Value
}

// 是否为空
func (s *ListQueue) Empty() bool {
	return s.length == 0
}

// 如果e在队列上，返回从1开始的索引值（top结点为1）
// 如果e不队列上，返回-1
func (s *ListQueue) Search(e interface{}) int {
	if s.Empty() {
		return -1
	}

	cur := s.head
	i := 1
	for cur != nil {
		if cur.Value == e {
			return i
		}
		i++
		cur = cur.Next
	}
	return -1
}
