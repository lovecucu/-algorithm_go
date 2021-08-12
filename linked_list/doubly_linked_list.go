package linked_list

// node 结构体
type Element struct {
	prev, next *Element
	list       *List // 找出所属的List
	Value      interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root { // 判断p为非头结点
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root { // 判断p为非头结点
		return p
	}
	return nil
}

// 双链表结构体
type List struct {
	root Element
	len  int
}

// new一个List
func New() *List { return new(List).Init() }

func (l *List) Init() *List { // 带头结点的List，root不存值，List的第一个元素通过root.next获得，list判空方法root.next == nil
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// v插到at元素的后面
func (l *List) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// e插到at的后面
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

// v插入到mark后
func (l *List) InsertAfter(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// v插入到mark前
func (l *List) InsertBefore(v interface{}, mark *Element) *Element {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

// 长度
func (l *List) Len() int {
	return l.len
}

// 将e移到mark后
func (l *List) MoveAfter(e, mark *Element) {
	if mark.list != l || e == mark || e.list != l {
		return
	}
	l.move(e, mark)
}

// 将e移动mark前
func (l *List) MoveBefore(e, mark *Element) {
	if mark.list != l || e == mark || e.list != l {
		return
	}
	l.move(e, mark.prev)
}

// 将e移到尾部
func (l *List) MoveToBack(e *Element) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

// 将e移动头部
func (l *List) MoveToFront(e *Element) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// 将e移到at后面，并返回e
func (l *List) move(e, at *Element) *Element {
	if e == at {
		return e
	}
	// 把e两边结点的prev,next处理下
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	return e
}

// 将值v尾插入l，并返回v对应的Element
func (l *List) PushBack(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// 将other尾插入l
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// 将值v头插入l，并返回v对应的Element
func (l *List) PushFront(v interface{}) *Element {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// 将other头插入l
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	// 会将other倒序头插入l
	for i, e := 0, other.Back(); i < other.Len(); i, e = i+1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}

// 延迟初始化
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// 移除e，并返回e的值
func (l *List) Remove(e *Element) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// 从l中移除e，并返回e
func (l *List) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
	return e
}
