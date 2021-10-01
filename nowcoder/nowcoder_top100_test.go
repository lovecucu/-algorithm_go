package nowcoder

import (
	"fmt"
	"testing"
)

func TestIsPail(t *testing.T) {
	head := &ListNode{Val: 1}
	if !isPail(head) || !isPail(nil) {
		t.Error(`TestIsPail failed`)
	}

	head.Next = &ListNode{Val: 2}
	if isPail(head) {
		t.Error(`TestIsPail failed`)
	}

	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
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

func TestPathSum(t *testing.T) {
	tree := &TreeNode{Val: 1}
	tree.Left = &TreeNode{Val: 2}
	if fmt.Sprint(pathSum(tree, 1)) != "[]" {
		t.Error(`TestSumNumbers failed`)
	}
	if fmt.Sprint(pathSum(tree, 3)) != "[[1 2]]" {
		t.Error(`TestSumNumbers failed`)
	}
	tree.Right = &TreeNode{Val: 1}
	tree.Right.Left = &TreeNode{Val: 1}
	if fmt.Sprint(pathSum(tree, 3)) != "[[1 2] [1 1 1]]" {
		t.Error(`TestSumNumbers failed`)
	}
}

func TestReverseBetween(t *testing.T) {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	if SprintNode(reverseBetween(list, 2, 3)) != "[1 3 2]" {
		t.Error(`TestReverseBetween failed`)
	}

	list = &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}
	list.Next.Next.Next = &ListNode{Val: 4}
	list.Next.Next.Next.Next = &ListNode{Val: 5}
	if SprintNode(reverseBetween(list, 2, 4)) != "[1 4 3 2 5]" {
		t.Error(`TestReverseBetween failed`)
	}
}

func TestReverseInt(t *testing.T) {
	if reverseInt(12) != 21 || reverseInt(-123) != -321 || reverseInt(10) != 1 || reverseInt(1147483649) != 0 {
		t.Error(`TestReverseInt failed`, reverseInt(-123))
	}
}
