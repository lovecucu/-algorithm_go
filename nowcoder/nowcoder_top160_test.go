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
	if fmt.Sprint(solveRotateArray(4, 2, []int{1, 2, 3, 4})) != "[3 4 1 2]" || fmt.Sprint(solveRotateArray(5, 2, []int{1, 2, 3, 4, 5})) != "[3 4 5 1 2]" {
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
