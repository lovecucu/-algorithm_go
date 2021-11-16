package crackingcodinginterview

import (
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	if SprintNode(removeDuplicateNodes(root)) != "[1 2 3]" {
		t.Error(`TestRemoveDuplicateNodes failed`)
	}
}

func TestKthToLast(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	if kthToLast(root, 1) != 4 || kthToLast(root, 2) != 3 || kthToLast(root, 3) != 2 || kthToLast(root, 4) != 1 {
		t.Error(`TestKthToLast failed`)
	}
}
