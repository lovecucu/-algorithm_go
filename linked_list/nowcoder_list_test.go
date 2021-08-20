package linked_list

import "testing"

func TestNcListReverseList(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	reverse := ReverseList(head)
	if reverse.Val != 3 || reverse.Next.Val != 2 || reverse.Next.Next.Val != 1 {
		t.Error(`TestNcListReverseList failed`)
	}
}

func TestNcListHasCycle(t *testing.T) {
	node2 := &ListNode{Val: 2}
	head := &ListNode{Val: 1}
	head.Next = node2
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = node2
	if !hasCycle(head) {
		t.Error(`TestNcListHasCycle failed`)
	}
}
