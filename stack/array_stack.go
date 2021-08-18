package stack

type ArrayStack struct {
	data   []interface{}
	length int
}

// 初始化长度为capacity的栈
func NewArrayStack(capacity int) *ArrayStack {
	if capacity < 1 {
		return nil
	}
	stack := new(ArrayStack)
	stack.data = make([]interface{}, capacity)
	return stack
}

// 查看top元素，但不删除
func (s *ArrayStack) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	return s.data[s.length-1]
}

// 出栈
func (s *ArrayStack) Pop() interface{} {
	if s.Empty() {
		return nil
	}

	if s.length == cap(s.data)/4 { // 1/4时缩容
		s.resize(cap(s.data) / 2)
	}

	item := s.data[s.length-1]
	s.data = s.data[:s.length-1]
	s.length--
	return item
}

// 入栈
func (s *ArrayStack) Push(e interface{}) {
	if !s.Empty() && s.length == cap(s.data) { // 不够时扩容成2倍
		s.resize(cap(s.data) * 2)
	}
	s.data[s.length] = e
	s.length++
}

// 是否为空
func (s *ArrayStack) Empty() bool {
	return s.length == 0
}

// 如果e在栈上，返回从1开始的索引值（top结点为1）
// 如果e不栈上，返回-1
func (s *ArrayStack) Search(e interface{}) int {
	if s.Empty() {
		return -1
	}

	for i := 0; i < s.length; i++ {
		if s.data[i] == e {
			return s.length - i
		}
	}
	return -1
}

// 扩缩容
func (s *ArrayStack) resize(capacity int) {
	newData := make([]interface{}, capacity)
	for i := 0; i < s.length; i++ {
		newData[i] = s.data[i]
	}
	s.data = newData
}
