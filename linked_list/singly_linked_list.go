package main

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

// 插入指定node后
func (this *SingleLinkedList) InsertAfter(n *ListNode, v interface{}) bool {
	if !this.Contains(n) {
		return false
	}

	// 将newNode插入到n后面
	newNode := NewListNode(v)
	oldNext := n.next
	newNode.next = oldNext
	n.next = newNode
	this.length++
	return true
}

// 插入指定node前
func (this *SingleLinkedList) InsertBefore(n *ListNode, v interface{}) bool {
	if !this.Contains(n) {
		return false
	}

	if this.isHeadNode(n) { // n为head结点，则换head
		newNode := NewListNode(v)
		this.head = newNode
		newNode.next = n
		this.length++
	} else { // n为非head结点时，找出pre
		pre := this.getPreNode(n) // 获取pre结点
		if nil == pre {
			return false
		}
		newNode := NewListNode(v)
		pre.next = newNode
		newNode.next = n
		this.length++
	}

	return true
}

// 头插
func (this *SingleLinkedList) InsertToHead(v interface{}) bool {
	newNode := NewListNode(v)
	if this.IsEmpty() { // 链表为空，则初始化head结点
		this.head = newNode
		this.length++
		return true
	}
	return this.InsertBefore(this.head, v)
}

// 尾插
func (this *SingleLinkedList) InsertToTail(v interface{}) bool {
	newNode := NewListNode(v)
	if this.IsEmpty() { // 链表为空，则初始化head结点
		this.head = newNode
		this.length++
		return true
	}
	return this.InsertAfter(this.getTailNode(), v)
}

// 根据索引获取结点
func (this *SingleLinkedList) FindByIndex(idx int) (n *ListNode) {
	if idx >= this.length {
		return
	}
	n = this.head
	for i := 0; i < idx; i++ {
		if i == idx {
			return
		}
		n = n.next
	}
	return
}

// 删除结点
func (this *SingleLinkedList) DeleteNode(n *ListNode) bool {
	if !this.Contains(n) {
		return false
	}

	if this.isHeadNode(n) { // 删除head结点
		this.head = this.head.next
	} else {
		pre := this.getPreNode(n)
		pre.next = n.next
	}
	this.length--
	return true
}

// 是否包含
func (this *SingleLinkedList) Contains(n *ListNode) bool {
	if n == nil || this.IsEmpty() {
		return false
	}

	cur := this.head
	for nil != cur {
		if n == cur {
			return true
		}
		cur = cur.next
	}
	return false
}

// 链表长度
func (this *SingleLinkedList) Len() int {
	return this.length
}

// 链表是否为空
func (this *SingleLinkedList) IsEmpty() bool {
	return this.length == 0
}

// string化
func (this *SingleLinkedList) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("SingleList: size = %d\n", this.Len()))
	buffer.WriteString("[")
	cur := this.head
	for i := 0; i < this.length; i++ {
		buffer.WriteString(fmt.Sprint(cur.value))
		if i != this.length-1 {
			buffer.WriteString(",")
		}
		cur = cur.next
	}
	buffer.WriteString("]")
	return buffer.String()
}

// 是否为head结点
func (this *SingleLinkedList) isHeadNode(n *ListNode) bool {
	return n != nil && n == this.head
}

// 获取结点的pre结点
func (this *SingleLinkedList) getPreNode(n *ListNode) (pre *ListNode) {
	if !this.Contains(n) {
		return
	}

	cur := this.head
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
func (this *SingleLinkedList) getTailNode() (tail *ListNode) {
	if this.IsEmpty() {
		return
	}

	tail = this.head
	for nil != tail {
		tail = tail.next
	}
	return
}

var _ SingleLinkedListInterface = (*SingleLinkedList)(nil)
