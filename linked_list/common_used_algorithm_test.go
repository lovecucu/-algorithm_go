package linked_list

import (
	"fmt"
	"testing"
)

func TestReverseListRecursion(t *testing.T) {
	head := initHeadNode(0)
	reverse := reverseListRecursion(head)
	if head != nil || reverse != nil {
		t.Error(`TestReverseListIteration failed`)
	}

	head = initHeadNode(5)
	reverse = reverseListRecursion(head)
	target := "54321"
	real := getString(reverse)
	if real != target {
		t.Error(`TestReverseListRecursion failed`)
	}
}

func TestReverseListIteration(t *testing.T) {
	head := initHeadNode(0)
	reverse := reverseListIteration(head)
	if head != nil || reverse != nil {
		t.Error(`TestReverseListIteration failed`)
	}

	head = initHeadNode(5)
	reverse = reverseListIteration(head)
	target := "54321"
	real := getString(reverse)
	if real != target {
		t.Error(`TestReverseListIteration failed`)
	}
}

func TestMiddleNode(t *testing.T) {
	head := initHeadNode(0)
	middle := middleNode(head)
	if head != nil || middle != nil {
		t.Error(`TestMiddleNode failed`)
	}

	head = initHeadNode(5)
	middle = middleNode(head)
	target := "345"
	real := getString(middle)
	if real != target {
		t.Error(`TestMiddleNode failed`, real)
	}

	head = initHeadNode(6)
	middle = middleNode(head)
	target = "456"
	real = getString(middle)
	if real != target {
		t.Error(`TestMiddleNode failed`, real)
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	head := initHeadNode(0)
	remove := removeNthFromEnd(head, 0)
	if head != nil || remove != nil {
		t.Error(`TestRemoveNthFromEnd failed`)
	}

	head = initHeadNode(5)
	remove = removeNthFromEnd(head, 7)
	target := "12345"
	real := getString(remove)
	if real != target {
		t.Error(`TestRemoveNthFromEnd failed`, real)
	}

	head = initHeadNode(5)
	remove = removeNthFromEnd(head, 2)
	target = "1235"
	real = getString(remove)
	if real != target {
		t.Error(`TestRemoveNthFromEnd failed`, real)
	}

	head = initHeadNode(5)
	remove = removeNthFromEnd(head, 5)
	target = "2345"
	real = getString(remove)
	if real != target {
		t.Error(`TestRemoveNthFromEnd failed`, real)
	}
}

func TestMergeTwoLists(t *testing.T) {
	l1 := initHeadNode(4)
	l2 := initHeadNode(2)
	merge := mergeTwoLists(l1, l2)
	target := "112234"
	real := getString(merge)
	if target != real {
		t.Error(`TestMergeTwoLists failed`, real)
	}

	l3 := initHeadNodeWithStartStep(3, 1, 2)
	l4 := initHeadNodeWithStartStep(4, 0, 2)
	target = "0123456"
	real = getString(mergeTwoLists(l3, l4))
	if target != real {
		t.Error(`TestMergeTwoLists failed`, real)
	}

	l5 := initHeadNode(2)
	l6 := initHeadNode(0)
	if l5 != mergeTwoLists(l5, l6) {
		t.Error(`TestMergeTwoLists failed`, real)
	}

	l7 := initHeadNode(0)
	l8 := initHeadNode(1)
	if l8 != mergeTwoLists(l7, l8) {
		t.Error(`TestMergeTwoLists failed`, real)
	}
}

func initHeadNode(n int) *ListNode {
	if n < 1 {
		return nil
	}
	head := &ListNode{Val: 1}
	cur := head
	for i := 2; i <= n; i++ {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	return head
}

func initHeadNodeWithStartStep(n, start, step int) *ListNode {
	if n < 1 {
		return nil
	}

	if start < 0 {
		start = 0
	}

	if step < 1 {
		step = 1
	}

	head := &ListNode{Val: start}
	cur := head
	for i := start + step; i < start+n*step; i += step {
		cur.Next = &ListNode{Val: i}
		cur = cur.Next
	}
	return head
}

func getString(head *ListNode) string {
	real := ""
	for tmp := head; tmp != nil; tmp = tmp.Next {
		real += fmt.Sprint(tmp.Val)
	}
	return real
}
