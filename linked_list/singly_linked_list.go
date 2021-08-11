package linked_list

import (
	"bytes"
	"fmt"
)

// 单链表 接口
type SingleLinkedListInterface interface {
	InsertAfter(*ListNode, interface{}) bool
	InsertBefore(*ListNode, interface{}) bool
	InsertToHead(interface{}) bool
	InsertToTail(interface{}) bool
	FindByIndex(int) *ListNode
	DeleteNode(*ListNode) bool
	Contains(*ListNode) bool
	Len() int
	IsEmpty() bool
}

// node 结构体
type ListNode struct {
	next  *ListNode
	value interface{}
}

// 根据值new一个Node
func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

// 单链表结构体
type SingleLinkedList struct {
	head   *ListNode
	length int
}

// new一个linkedList
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{nil, 0}
}

// 头插
func (s *SingleLinkedList) InsertToHead(v interface{}) bool {
	newNode := NewListNode(v)
	if s.IsEmpty() { // 链表为空，则初始化head结点
		s.head = newNode
		s.length++
		return true
	}
	return s.InsertBefore(s.head, v)
}

// 尾插
func (s *SingleLinkedList) InsertToTail(v interface{}) bool {
	newNode := NewListNode(v)
	if s.IsEmpty() { // 链表为空，则初始化head结点
		s.head = newNode
		s.length++
		return true
	}
	return s.InsertAfter(s.getTailNode(), v)
}

// 插入指定node后
func (s *SingleLinkedList) InsertAfter(n *ListNode, v interface{}) bool {
	if !s.Contains(n) {
		return false
	}

	// 将newNode插入到n后面
	newNode := NewListNode(v)
	oldNext := n.next
	newNode.next = oldNext
	n.next = newNode
	s.length++
	return true
}

// 插入指定node前
func (s *SingleLinkedList) InsertBefore(n *ListNode, v interface{}) bool {
	if !s.Contains(n) {
		return false
	}

	if s.isHeadNode(n) { // n为head结点，则换head
		newNode := NewListNode(v)
		s.head = newNode
		newNode.next = n
		s.length++
	} else { // n为非head结点时，找出pre
		pre := s.getPreNode(n) // 获取pre结点
		if nil == pre {
			return false
		}
		newNode := NewListNode(v)
		pre.next = newNode
		newNode.next = n
		s.length++
	}

	return true
}

// 根据索引获取结点
func (s *SingleLinkedList) FindByIndex(idx int) (n *ListNode) {
	if idx >= s.length {
		return
	}
	n = s.head
	for i := 0; i < idx; i++ {
		if i == idx {
			return
		}
		n = n.next
	}
	return
}

// 删除结点
func (s *SingleLinkedList) DeleteNode(n *ListNode) bool {
	if !s.Contains(n) {
		return false
	}

	if s.isHeadNode(n) { // 删除head结点
		s.head = s.head.next
	} else {
		pre := s.getPreNode(n)
		pre.next = n.next
	}
	s.length--
	return true
}

// 是否包含
func (s *SingleLinkedList) Contains(n *ListNode) bool {
	if n == nil || s.IsEmpty() {
		return false
	}

	cur := s.head
	for nil != cur {
		if n == cur {
			return true
		}
		cur = cur.next
	}
	return false
}

// 链表长度
func (s *SingleLinkedList) Len() int {
	return s.length
}

// 链表是否为空
func (s *SingleLinkedList) IsEmpty() bool {
	return s.length == 0
}

// string化
func (s *SingleLinkedList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("SingleLinkedList: size = %d\n", s.Len()))
	buffer.WriteString("[")
	cur := s.head
	for i := 0; i < s.length; i++ {
		buffer.WriteString(fmt.Sprint(cur.value))
		if i != s.length-1 {
			buffer.WriteString(",")
		}
		cur = cur.next
	}
	buffer.WriteString("]")
	return buffer.String()
}

// 是否为head结点
func (s *SingleLinkedList) isHeadNode(n *ListNode) bool {
	return n != nil && n == s.head
}

// 获取结点的pre结点
func (s *SingleLinkedList) getPreNode(n *ListNode) (pre *ListNode) {
	if !s.Contains(n) {
		return
	}

	cur := s.head
	for nil != cur {
		if cur == n {
			break
		}
		pre = cur
		cur = cur.next
	}
	return
}

// 获取结点的tail结点
func (s *SingleLinkedList) getTailNode() (tail *ListNode) {
	if s.IsEmpty() {
		return
	}

	cur := s.head
	for nil != cur {
		tail = cur
		cur = cur.next
	}
	return
}

var _ SingleLinkedListInterface = (*SingleLinkedList)(nil)
