package crackingcodinginterview

import (
	"fmt"
	"sort"
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

func TestDeleteNode(t *testing.T) {
	root := &ListNode{Val: 1}
	second := &ListNode{Val: 2}
	root.Next = second
	third := &ListNode{Val: 3}
	root.Next.Next = third
	root.Next.Next.Next = &ListNode{Val: 4}
	deleteNode(second)
	if SprintNode(root) != "[1 3 4]" {
		t.Error(`TestDeleteNode failed`)
	}

	deleteNode(root.Next)
	if SprintNode(root) != "[1 4]" {
		t.Error(`TestDeleteNode failed`, SprintNode(root))
	}
}

func TestPartition(t *testing.T) {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 4}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	// 保证升序输出
	ret := Node2Array(partition(root, 3))
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	if fmt.Sprint(ret) != "[1 2 2 3 4 5]" {
		t.Error(`TestPartition failed`)
	}

	root = &ListNode{Val: 2}
	root.Next = &ListNode{Val: 1}
	if SprintNode(partition(root, 2)) != "[1 2]" {
		t.Error(`TestPartition failed`)
	}
}
