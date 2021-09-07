package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr1 := []int{5, 2, 3, 1, 4}
	target1 := "[1 2 3 4 5]"
	if fmt.Sprint(BubbleSort(arr1)) != target1 {
		t.Error(`TestBubbleSort failed`)
	}

	arr2 := []int{5, 1, 6, 2, 5}
	target2 := "[1 2 5 5 6]"
	if fmt.Sprint(BubbleSort(arr2)) != target2 {
		t.Error(`TestBubbleSort failed`)
	}

	if fmt.Sprint(BubbleSort([]int{1})) != "[1]" {
		t.Error(`TestBubbleSort failed`)
	}
}

func TestInsertSort(t *testing.T) {
	arr1 := []int{5, 2, 3, 1, 4}
	target1 := "[1 2 3 4 5]"
	if fmt.Sprint(InsertSort(arr1)) != target1 {
		t.Error(`TestInsertSort failed`)
	}

	arr2 := []int{5, 1, 6, 2, 5}
	target2 := "[1 2 5 5 6]"
	if fmt.Sprint(InsertSort(arr2)) != target2 {
		t.Error(`TestInsertSort failed`)
	}

	if fmt.Sprint(InsertSort([]int{1})) != "[1]" {
		t.Error(`TestInsertSort failed`)
	}
}

func TestSelectSort(t *testing.T) {
	arr1 := []int{5, 2, 3, 1, 4}
	target1 := "[1 2 3 4 5]"
	if fmt.Sprint(SelectSort(arr1)) != target1 {
		t.Error(`TestSelectSort failed`)
	}

	arr2 := []int{5, 1, 6, 2, 5}
	target2 := "[1 2 5 5 6]"
	if fmt.Sprint(SelectSort(arr2)) != target2 {
		t.Error(`TestSelectSort failed`)
	}

	if fmt.Sprint(SelectSort([]int{1})) != "[1]" {
		t.Error(`TestSelectSort failed`)
	}
}

func TestMergeSort(t *testing.T) {
	arr1 := []int{5, 2, 3, 1, 4}
	target1 := "[1 2 3 4 5]"
	if fmt.Sprint(MergeSort(arr1)) != target1 {
		t.Error(`TestMergeSort failed`)
	}

	arr2 := []int{5, 1, 6, 2, 5}
	target2 := "[1 2 5 5 6]"
	if fmt.Sprint(MergeSort(arr2)) != target2 {
		t.Error(`TestMergeSort failed`)
	}

	if fmt.Sprint(MergeSort([]int{1})) != "[1]" {
		t.Error(`TestMergeSort failed`)
	}
}

func TestQuickSort(t *testing.T) {
	arr1 := []int{5, 2, 3, 1, 4}
	target1 := "[1 2 3 4 5]"
	if fmt.Sprint(QuickSort(arr1)) != target1 {
		t.Error(`TestQuickSort failed`)
	}

	arr2 := []int{5, 1, 6, 2, 5}
	target2 := "[1 2 5 5 6]"
	if fmt.Sprint(QuickSort(arr2)) != target2 {
		t.Error(`TestQuickSort failed`)
	}

	if fmt.Sprint(QuickSort([]int{1})) != "[1]" {
		t.Error(`TestQuickSort failed`)
	}
}
