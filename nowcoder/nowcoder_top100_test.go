package nowcoder

import "testing"

func TestIsPail(t *testing.T) {
	head := &ListNode{Val: 1}
	if !isPail(head) || !isPail(nil) {
		t.Error(`TestIsPail failed`)
	}

	head.Next = &ListNode{Val: 2}
	if isPail(head) {
		t.Error(`TestIsPail failed`)
	}

	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	if !isPail(head) {
		t.Error(`TestIsPail failed`)
	}
}

func TestSumNumbers(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	if sumNumbers(tree) != 12 {
		t.Error(`TestSumNumbers failed`)
	}
	tree.Right = &TreeNode{Val: 3}
	if sumNumbers(tree) != 25 {
		t.Error(`TestSumNumbers failed`)
	}
}
