package nowcoder

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	operators := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 3},
		{2, 1},
		{1, 4, 4},
		{2, 2},
		{1, 3, 5},
		{2, 3},
	}
	result := LRU(operators, 3)
	target := "[1 -1 5]"
	if fmt.Sprint(result) != target {
		t.Error(`TestLRU failed`, fmt.Sprint(result))
	}

	operators = [][]int{{1, 1, 1}, {1, 2, 2}, {2, 1}, {1, 3, 3}, {2, 2}, {1, 4, 4}, {2, 1}, {2, 3}, {2, 4}}
	result = LRU(operators, 2)
	target = "[1 -1 -1 3 4]"
	if fmt.Sprint(result) != target {
		t.Error(`TestLRU failed`, fmt.Sprint(result))
	}
}

func TestThreeOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right

	target := "[[1 2 3] [2 1 3] [2 3 1]]"
	if fmt.Sprint(threeOrders(root)) != target {
		t.Error(`TestThreeOrder failed`)
	}

	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}

	right.Right = &TreeNode{Val: 6}
	target = "[[1 2 4 5 7 8 3 6] [4 2 7 5 8 1 3 6] [4 7 8 5 2 6 3 1]]"
	if fmt.Sprint(threeOrders(root)) != target {
		t.Error(`TestThreeOrder failed`)
	}
}

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 1}
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root.Left = left
	root.Right = right
	left.Left = &TreeNode{Val: 4}
	left.Right = &TreeNode{Val: 5}
	left.Right.Left = &TreeNode{Val: 7}
	left.Right.Right = &TreeNode{Val: 8}

	fmt.Println(levelOrder(root))
}
