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

func TestFindElement(t *testing.T) {
	mat := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	if fmt.Sprint(findElement(mat, 2, 3, 6)) != "[1 2]" {
		t.Error(`TestFindElement failed`)
	}
	mat = [][]int{
		{1, 2, 3},
	}
	if fmt.Sprint(findElement(mat, 1, 3, 2)) != "[0 1]" {
		t.Error(`TestFindElement failed`)
	}
	mat = [][]int{
		{1, 4, 8},
		{2, 5, 9},
	}
	if fmt.Sprint(findElement(mat, 2, 3, 5)) != "[1 1]" {
		t.Error(`TestFindElement failed`)
	}
}

func TestLCSPlus1(t *testing.T) {
	if LCSPlus1("1A2C3D4B56", "B1D23A456A") != "123456" || LCSPlus1("abc", "def") != "-1" || LCSPlus1("abc", "abc") != "abc" || LCSPlus1("abc", "") != "-1" {
		t.Error(`TestLCSPlus1 failed`)
	}
}

func TestJudgeIt(t *testing.T) {
	if fmt.Sprint(judgeIt(nil)) != "[true true]" {
		t.Error("TestJudgeIt failed")
	}

	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}

	if fmt.Sprint(judgeIt(root)) != "[true true]" {
		t.Error("TestJudgeIt failed")
	}

	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	if fmt.Sprint(judgeIt(root)) != "[true false]" {
		t.Error("TestJudgeIt failed")
	}

	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	if fmt.Sprint(judgeIt(root)) != "[false true]" {
		t.Error("TestJudgeIt failed")
	}

	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 8}
	if fmt.Sprint(judgeIt(root)) != "[true false]" {
		t.Error("TestJudgeIt failed", fmt.Sprint(judgeIt(root)))
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 8}
	if fmt.Sprint(judgeIt(root)) != "[false false]" {
		t.Error("TestJudgeIt failed", fmt.Sprint(judgeIt(root)))
	}
}

func TestDeleteDuplicates(t *testing.T) {
	root := &ListNode{Val: 1}
	if SprintNode(deleteDuplicates(root)) != "[1]" {
		t.Error(`TestDeleteDuplicates failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicates(root)) != "[2]" {
		t.Error(`TestDeleteDuplicates failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicates(root)) != "[]" {
		t.Error(`TestDeleteDuplicates failed`)
	}
}

func TestDeleteDuplicatesEasy(t *testing.T) {
	root := &ListNode{Val: 1}
	if SprintNode(deleteDuplicatesEasy(root)) != "[1]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicatesEasy(root)) != "[1 2]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}

	root = &ListNode{Val: 1}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 2}
	root.Next.Next.Next = &ListNode{Val: 2}
	if SprintNode(deleteDuplicatesEasy(root)) != "[1 2]" {
		t.Error(`deleteDuplicatesEasy failed`)
	}
}

func TestMinNumberDisappeared(t *testing.T) {
	if minNumberDisappeared([]int{1, 0, 2}) != 3 {
		t.Error(`TestMinNumberDisappeared failed`)
	}

	if minNumberDisappeared([]int{-2, 3, 4, 1, 5}) != 2 {
		t.Error(`TestMinNumberDisappeared failed`)
	}

	if minNumberDisappeared([]int{4, 5, 6, 8, 9}) != 1 {
		t.Error(`TestMinNumberDisappeared failed`)
	}
}

func TestAtoi(t *testing.T) {
	if atoi("  010") != 10 || atoi("+12") != 12 || atoi("-12ab") != -12 || atoi("123") != 123 || atoi("123e123") != 123 || atoi("e123") != 0 {
		t.Error(`TestAtoi failed`)
	}
}
