package array

import (
	"fmt"
	"testing"
)

func TestNcArrayMerge(t *testing.T) {
	A := []int{4, 5, 6}
	B := []int{1, 2, 3}
	C := merge(A, 3, B, 3)
	if len(C) != 6 {
		t.Error(`TestNcArrayMerge failed`)
	}
	target := "[1 2 3 4 5 6]"
	real := fmt.Sprint(C)
	if target != real {
		t.Error(`TestNcArrayMerge failed`)
	}

	A = []int{4, 5, 6, 0, 0, 0}
	B = []int{1, 2, 3}
	C = merge(A, 3, B, 3)
	if len(C) != 6 {
		t.Error(`TestNcArrayMerge failed`)
	}
	target = "[1 2 3 4 5 6]"
	real = fmt.Sprint(C)
	if target != real {
		t.Error(`TestNcArrayMerge failed`)
	}

	A = []int{1, 2, 3}
	B = []int{2, 5, 6}
	C = merge(A, 3, B, 3)
	if len(C) != 6 {
		t.Error(`TestNcArrayMerge failed`)
	}
	target = "[1 2 2 3 5 6]"
	real = fmt.Sprint(C)
	if target != real {
		t.Error(`TestNcArrayMerge failed`)
	}
}

func TestNcArrayTwoSum(t *testing.T) {
	A := []int{3, 2, 4}
	sum := 6
	target := "[2 3]"
	real := twoSum(A, sum)
	if target != fmt.Sprint(real) {
		t.Error(`TestNcArrayTwoSum failed`)
	}

	A = []int{4, 5, 6}
	sum = 1
	target = "[]"
	real = twoSum(A, sum)
	if target != fmt.Sprint(real) {
		t.Error(`TestNcArrayTwoSum failed`)
	}
}

func TestNcArrayGetNumberOfK(t *testing.T) {
	A := []int{1, 2, 3, 3, 3, 3, 4, 5}
	if GetNumberOfK(A, 3) != 4 {
		t.Error(`TestNcArrayGetNumberOfK failed`, GetNumberOfK(A, 3))
	}

	if GetNumberOfK(A, 1) != 1 {
		t.Error(`TestNcArrayGetNumberOfK failed`, GetNumberOfK(A, 1))
	}

	if GetNumberOfK(A, 2) != 1 {
		t.Error(`TestNcArrayGetNumberOfK failed`, GetNumberOfK(A, 2))
	}

	if GetNumberOfK(A, 4) != 1 {
		t.Error(`TestNcArrayGetNumberOfK failed`, GetNumberOfK(A, 4))
	}

	if GetNumberOfK(A, 5) != 1 {
		t.Error(`TestNcArrayGetNumberOfK failed`, GetNumberOfK(A, 5))
	}
}

func TestNcArrayBinarySearch2(t *testing.T) {
	A := []int{1, 2, 3, 3, 3, 3, 4, 5}
	if binarySearch2(A, 0, 7, 3) != 3 {
		t.Error(`TestNcArrayBinarySearch2 failed`)
	}

	if binarySearch2(A, 0, 7, 2) != 1 {
		t.Error(`TestNcArrayBinarySearch2 failed`)
	}

	if binarySearch2(A, 0, 7, 1) != 0 {
		t.Error(`TestNcArrayBinarySearch2 failed`)
	}

	if binarySearch2(A, 0, 7, 4) != 6 {
		t.Error(`TestNcArrayBinarySearch2 failed`)
	}

	if binarySearch2(A, 0, 7, 5) != 7 {
		t.Error(`TestNcArrayBinarySearch2 failed`)
	}
}

func TestNcArrayMoreThanHalfNum(t *testing.T) {
	if MoreThanHalfNum([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}) != 2 {
		t.Error(`TestMoreThanHalfNum failed`)
	}

	if MoreThanHalfNum([]int{3, 3, 3, 3, 2, 2, 2}) != 3 {
		t.Error(`TestMoreThanHalfNum failed`)
	}

	if MoreThanHalfNum([]int{1}) != 1 {
		t.Error(`TestMoreThanHalfNum failed`)
	}
}

func TestNcArrayMaxThreeMultiple(t *testing.T) {
	if solve([]int{3, 4, 1, 2}) != 24 {
		t.Error(`TestNcArrayMaxThreeMultiple failed`)
	}

	if solve([]int{5, 14, 100, 30, 5}) != 42000 {
		t.Error(`TestNcArrayMaxThreeMultiple failed`)
	}

	if solve([]int{1000000, 999999, 100, 1, 5}) != 99999900000000 {
		t.Error(`TestNcArrayMaxThreeMultiple failed`)
	}

	if solve([]int{1, -5, -2, 3}) != 30 {
		t.Error(`TestNcArrayMaxThreeMultiple failed`)
	}

	mul := solve([]int{3472, 7789, 7955, -7098, -9281, 6101, 5051, 7778, 3090, 7423, -7151, 5652, 1595, -8094, 677, -8324, 8347, -2482, 9313, -9338, -3157, 8559, 6945, 3618, 3087, 121, -8468, 3225, 1356, 6939, 2799, -7231, -6309, -5453, 633, -8689, -4776, 2714, -2743, -1409, 5918, -3333, 1803, 8330, -2206, -6117, -4486, -7903, -4375, -3739, 2897, 8056, -5864, -522, 7451, -4541, -2813, 5790, -532, -6517, 925})
	if mul != 807120253114 {
		t.Error(`TestNcArrayMaxThreeMultiple failed`, mul)
	}
}
