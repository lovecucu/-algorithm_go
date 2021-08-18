package stack

type ListNode struct {
	Prev, Next *ListNode
	Value      interface{}
}
type ListStack struct {
	tail   *ListNode
	length int
}

// 初始化栈
func NewListStack() *ListStack {
	stack := new(ListStack)
	return stack
}

// 入栈
func (s *ListStack) Push(e interface{}) {
	node := &ListNode{Value: e}
	if s.Empty() {
		s.tail = node
	} else {
		s.tail.Next = node
		node.Prev = s.tail
		s.tail = node
	}
	s.length++
}

// 出栈
func (s *ListStack) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	item := s.tail
	s.length--
	if s.length == 0 {
		s.tail = nil
	} else {
		s.tail = s.tail.Prev
	}
	return item.Value
}

// 查看top元素，但不删除
func (s *ListStack) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	return s.tail.Value
}

// 是否为空
func (s *ListStack) Empty() bool {
	return s.length == 0
}

// 如果e在栈上，返回从1开始的索引值（top结点为1）
// 如果e不栈上，返回-1
func (s *ListStack) Search(e interface{}) int {
	if s.Empty() {
		return -1
	}

	cur := s.tail
	i := 1
	for cur != nil {
		if cur.Value == e {
			return i
		}
		i++
		cur = cur.Prev
	}
	return -1
}
