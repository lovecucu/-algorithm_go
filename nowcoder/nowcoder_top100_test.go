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

func TestUniquePaths(t *testing.T) {
	if uniquePaths(1, 1) != 1 || uniquePaths(2, 1) != 1 || uniquePaths(2, 2) != 2 || uniquePaths(4, 4) != 20 {
		t.Error(`TestUniquePaths failed`)
	}
}

func TestMergeInterval(t *testing.T) {
	// [[10,30],[20,60],[80,100],[150,180]]
	intervals := []*Interval{
		{20, 60},
		{10, 30},
		{80, 100},
		{150, 180},
	}

	merged := mergeInterval(intervals)
	if SPrintInterval(merged) != "[[10 60] [80 100] [150 180]]" {
		t.Error(`TestMergeInterval failed`)
	}

	// [[0,10],[10,20]]
	intervals = []*Interval{
		{0, 10},
		{10, 20},
	}
	merged = mergeInterval(intervals)
	if SPrintInterval(merged) != "[[0 20]]" {
		t.Error(`TestMergeInterval failed`)
	}
}

func TestFindMedianinTwoSortedAray(t *testing.T) {
	if findMedianinTwoSortedAray([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}) != 3 {
		t.Error(`TestFindMedianinTwoSortedAray failed`)
	}

	if findMedianinTwoSortedAray([]int{0, 1, 2}, []int{3, 4, 5}) != 2 {
		t.Error(`TestFindMedianinTwoSortedAray failed`)
	}
}

func TestGetMedian(t *testing.T) {
	// [5,2,3,4,1,6,7,0,8]
	Insert(5)
	fmt.Println(GetMedian())
	Insert(2)
	fmt.Println(GetMedian())
	Insert(3)
	fmt.Println(GetMedian())
	Insert(4)
	fmt.Println(GetMedian())
	Insert(1)
	fmt.Println(GetMedian())
	Insert(6)
	fmt.Println(GetMedian())
	Insert(7)
	fmt.Println(GetMedian())
	Insert(0)
	fmt.Println(GetMedian())
	Insert(8)
	fmt.Println(GetMedian())
}
