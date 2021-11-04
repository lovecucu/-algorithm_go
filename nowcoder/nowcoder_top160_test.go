package nowcoder

import (
	"fmt"
	"testing"
)

func TestKnapsack2(t *testing.T) {
	if knapsack2(10, 2, [][]int{{1, 3}, {10, 4}}) != 4 {
		t.Error(`TestKnapsack2 failed`)
	}

	if knapsack2(38, 50, [][]int{{43, 50}, {50, 45}, {31, 10}, {3, 32}, {2, 9}, {47, 15}, {27, 46}, {24, 45}, {36, 48}, {38, 29}, {25, 23}, {36, 48}, {10, 7}, {1, 5}, {22, 3}, {12, 13}, {35, 9}, {33, 12}, {3, 15}, {50, 30}, {18, 48}, {20, 19}, {34, 24}, {21, 10}, {25, 9}, {10, 21}, {27, 35}, {12, 27}, {35, 39}, {7, 45}, {25, 46}, {18, 23}, {33, 28}, {22, 50}, {17, 48}, {3, 44}, {34, 13}, {41, 2}, {34, 6}, {47, 46}, {48, 30}, {31, 7}, {34, 32}, {40, 50}, {31, 39}, {10, 41}, {22, 36}, {14, 48}, {46, 17}, {8, 33}}) != 224 {
		t.Error(`TestKnapsack2 failed`)
	}
}

func TestInversePairs(t *testing.T) {
	if InversePairs([]int{1, 2, 3, 4, 5, 6, 7, 0}) != 7 || InversePairs([]int{1, 2, 3}) != 0 || InversePairs([]int{1, 3, 2}) != 1 {
		t.Error(`TestInversePairs failed`)
	}
}

func TestNumberOf1(t *testing.T) {
	if NumberOf1(1025) != 2 || NumberOf1(1) != 1 || NumberOf1(-1) != 32 {
		t.Error(`TestNumberOf1 failed`, NumberOf1(-1))
	}
}

func TestSolveRotateArray(t *testing.T) {
	if fmt.Sprint(solveRotateArray(4, 2, []int{1, 2, 3, 4})) != "[3 4 1 2]" || fmt.Sprint(solveRotateArray(5, 2, []int{1, 2, 3, 4, 5})) != "[4 5 1 2 3]" {
		t.Error(`TestSolveRotateArray failed`)
	}
}

func TestSolveMaxSquare(t *testing.T) {
	if solveMaxSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}) != 4 {
		t.Error(`TestSolveMaxSquare failed`)
	}

	if solveMaxSquare([][]byte{{'1', '0', '0'}, {'0', '0', '0'}, {'0', '0', '0'}}) != 1 {
		t.Error(`TestSolveMaxSquare failed`)
	}

	if solveMaxSquare([][]byte{{'1', '1', '1'}, {'1', '1', '1'}, {'1', '1', '1'}}) != 9 {
		t.Error(`TestSolveMaxSquare failed`)
	}
}

func TestPrintTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	if fmt.Sprint(PrintTree(root)) != "[[1] [2 3] [4 5]]" {
		t.Error(`TestPrintTree failed`)
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	if fmt.Sprint(PrintTree(root)) != "[[1] [2 3] [4 5]]" {
		t.Error(`TestPrintTree failed`)
	}

	root = &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 10}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 11}
	if fmt.Sprint(PrintTree(root)) != "[[8] [6 10] [5 7 9 11]]" {
		t.Error(`TestPrintTree failed`)
	}
}

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	if fmt.Sprint(inorderTraversal(root)) != "[2 1 4 3 5]" {
		t.Error(`TestInorderTraversal failed`)
	}

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	if fmt.Sprint(inorderTraversal(root)) != "[4 2 5 1 3]" {
		t.Error(`TestInorderTraversal failed`)
	}
}

func TestSearchEasy(t *testing.T) {
	if searchEasy([]int{-1, 0, 3, 4, 6, 10, 13, 14}, 13) != 6 || searchEasy([]int{}, 3) != -1 || searchEasy([]int{-1, 0, 3, 4, 6, 10, 13, 14}, 2) != -1 {
		t.Error(`TestSearchEasy failed`)
	}
}

func TestMaxProfit134(t *testing.T) {
	if maxProfit134([]int{5, 4, 3, 2, 1}) != 0 || maxProfit134([]int{1, 2, 3, 4, 5}) != 4 {
		t.Error(`TestMaxProfit134 failed`)
	}
}

func TestReOrderArray(t *testing.T) {
	if fmt.Sprint(reOrderArray([]int{1, 2, 3, 4})) != "[1 3 2 4]" || fmt.Sprint(reOrderArray([]int{2, 4, 6, 5, 7})) != "[5 7 2 4 6]" || fmt.Sprint(reOrderArray([]int{1, 3, 5, 6, 7})) != "[1 3 5 7 6]" {
		t.Error(`TestReOrderArray failed`)
	}
}

func TestIsContains(t *testing.T) {
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 4}
	root1.Left.Right = &TreeNode{Val: 5}
	root1.Right.Left = &TreeNode{Val: 6}
	root1.Right.Right = &TreeNode{Val: 7}
	root1.Left.Left.Right = &TreeNode{Val: 8}
	root1.Left.Right.Left = &TreeNode{Val: 9}

	root2 := &TreeNode{Val: 2}
	root2.Left = &TreeNode{Val: 4}
	root2.Right = &TreeNode{Val: 5}
	root2.Left.Right = &TreeNode{Val: 8}
	root2.Right.Left = &TreeNode{Val: 9}

	if !isContains(root1, root2) {
		t.Error(`TestIsContains failed`)
	}
}

func TestIsContinuous(t *testing.T) {
	if !IsContinuous([]int{6, 0, 2, 0, 4}) || !IsContinuous([]int{0, 3, 2, 6, 4}) || IsContinuous([]int{1, 0, 0, 1, 0}) || IsContinuous([]int{13, 12, 11, 0, 1}) {
		t.Error(`TestIsContinuous failed`)
	}
}

func TestFirstNotRepeatingChar(t *testing.T) {
	if FirstNotRepeatingChar("google") != 4 || FirstNotRepeatingChar("aa") != -1 {
		t.Error(`TestFirstNotRepeatingChar failed`)
	}
}

func TestMergeTrees(t *testing.T) {
	r1 := &TreeNode{Val: 1}
	r1.Left = &TreeNode{Val: 3}
	r1.Right = &TreeNode{Val: 2}
	r1.Left.Left = &TreeNode{Val: 5}
	if mergeTrees(r1, nil) != r1 {
		t.Error(`TestMergeTrees failed`)
	}

	r2 := &TreeNode{Val: 2}
	r2.Left = &TreeNode{Val: 1}
	r2.Right = &TreeNode{Val: 3}
	r2.Left.Right = &TreeNode{Val: 4}
	r2.Right.Right = &TreeNode{Val: 7}
	if SprintTreeNode(mergeTrees(r1, r2)) != "[3 4 5 4 5 7]" {
		t.Error(`TestMergeTrees failed`)
	}
}

func TestFindError(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	if fmt.Sprint(findError(root)) != "[1 2]" {
		t.Error(`TestFindError failed`)
	}

	root = &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	if fmt.Sprint(findError(root)) != "[2 3]" {
		t.Error(`TestFindError failed`)
	}
}

func TestSortedArrayToBST(t *testing.T) {
	if sortedArrayToBST([]int{}) != nil {
		t.Error(`TestSortedArrayToBST failed`)
	}
	root := sortedArrayToBST([]int{-1, 0, 1, 2})
	if SprintTreeNode(root) != "[1 0 -1 2]" {
		t.Error(`TestSortedArrayToBST failed`)
	}
}

func TestMinNumberInRotateArray(t *testing.T) {
	if minNumberInRotateArray([]int{3, 4, 5, 1, 2}) != 1 || minNumberInRotateArray([]int{3, 100, 200, 3}) != 3 || minNumberInRotateArray([]int{4, 3, 3}) != 3 {
		t.Error(`TestMinNumberInRotateArray failed`)
	}
}

func TestGetNumberOfK(t *testing.T) {
	if GetNumberOfK([]int{1, 2, 3, 3, 3, 3, 4, 5}, 3) != 4 || GetNumberOfK([]int{1, 3, 4, 5}, 6) != 0 {
		t.Error(`TestGetNumberOfK failed`)
	}
}

func TestMaxlenEqualK(t *testing.T) {
	if maxlenEqualK([]int{1, -2, 1, 1, 1}, 0) != 3 || maxlenEqualK([]int{0, 1, 2, 3}, 3) != 3 {
		t.Error(`TestMaxlenEqualK failed`)
	}
}

func TestSolveThreeMaxProduct(t *testing.T) {
	if solveThreeMaxProduct([]int{3, 4, 1, 2}) != 24 || solveThreeMaxProduct([]int{-1, -2, -3, -4}) != -6 || solveThreeMaxProduct([]int{-3, -4, 3, 2, 4}) != 48 {
		t.Error(`TestSolveThreeMaxProduct failed`)
	}
}

func TestSolveRotateString(t *testing.T) {
	if solveRotateString("youzan", "zanyou") == false || solveRotateString("youzan", "youzan") == false || solveRotateString("ba", "aa") == true {
		t.Error(`TestSolveRotateString failed`)
	}
}
