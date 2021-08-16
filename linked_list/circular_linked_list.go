package linked_list

// 双向循环链表
type Ring struct {
	prev, next *Ring
	Value      interface{}
}

// new一个长度为n的循环链表
func NewRing(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p} // 创建下一个Ring节点
		p = p.next              // 指针移到下个Ring节点
	}
	p.next = r // 尾部节点的next指向头部节点
	r.prev = p // 头部节点的prev指向尾部节点
	return r   // 返回头节点
}

// 初始化
func (r *Ring) init() *Ring {
	r.prev = r
	r.next = r
	return r
}

// 返回链表长度，不care值
func (r *Ring) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.Next(); p != r; p = p.Next() {
			n++
		}
	}
	return n
}

// 返回r.next，如果r为空链表返回r
func (r *Ring) Next() *Ring {
	if r.next == nil {
		r.init()
	}
	return r.next
}

// 返回r.prev，如果r为空链表返回r
func (r *Ring) Prev() *Ring {
	if r.prev == nil {
		r.init()
	}
	return r.prev
}

// r与s指向同一个ring时，link会删除r和s之前的元素
// r与s指向不同ring时，link会把r.next指向s，同时p指向n
func (r *Ring) Link(s *Ring) *Ring {
	n := r.Next()
	if s != nil {
		p := s.Prev() // 保留s.prev
		r.next = s
		s.prev = r
		n.prev = p
		p.next = n
	}
	return n
}

// 删除r 与 Len() % n距离节点之前的元素
func (r *Ring) Unlink(n int) *Ring {
	if n <= 0 {
		return nil
	}
	return r.Link(r.Move(n + 1))
}

// r移动n，n < 0左移，n > 0右移
func (r *Ring) Move(n int) *Ring {
	if r.next == nil {
		return r.init()
	}
	switch {
	case n < 0:
		for ; n < 0; n++ {
			r = r.prev
		}
	case n > 0:
		for ; n > 0; n-- {
			r = r.next
		}
	}
	return r
}

// 遍历执行函数
func (r *Ring) Do(f func(interface{})) {
	if r != nil {
		f(r.Value)
		for p := r.Next(); p != r; p = p.Next() {
			f(p.Value)
		}
	}
}
