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
