package queue

type ArrayQueue struct {
	data        []interface{}
	cap, length int
}

// 初始化长度为capacity的栈
func NewArrayQueue(capacity int) *ArrayQueue {
	if capacity < 1 {
		return nil
	}
	stack := new(ArrayQueue)
	stack.data = make([]interface{}, capacity)
	stack.cap = capacity
	return stack
}

// 查看top元素，但不删除
func (s *ArrayQueue) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[0]
}

// 出队列
func (s *ArrayQueue) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	item := s.data[0]
	s.data = s.data[1:]
	s.length--
	if s.length == s.cap/4 && s.length != 0 { // 1/4时缩容
		s.resize(s.cap / 2)
	}
	return item
}

// 入队列
func (s *ArrayQueue) Push(e interface{}) {
	if !s.Empty() && s.length == s.cap { // 不够时扩容成2倍
		s.resize(s.cap * 2)
	}
	s.data[s.length] = e
	s.length++
}

// 是否为空
func (s *ArrayQueue) Empty() bool {
	return s.length == 0
}

// 如果e在队列上，返回从1开始的索引值（top结点为1）
// 如果e不队列上，返回-1
func (s *ArrayQueue) Search(e interface{}) int {
	if s.Empty() {
		return -1
	}

	for i := 0; i < s.length; i++ {
		if s.data[i] == e {
			return i + 1
		}
	}
	return -1
}

// 扩缩容
func (s *ArrayQueue) resize(capacity int) {
	newData := make([]interface{}, capacity)
	for i := 0; i < s.length; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
	s.cap = capacity
}
