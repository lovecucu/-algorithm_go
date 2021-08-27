package dichotomous

import "testing"

func TestNcDichotomousBinarySearchFirst(t *testing.T) {
	nums := []int{1, 2, 4, 4, 5}
	if binarySearchFirst(nums, 4) != 2 {
		t.Error(`TestNcDichotomousBinarySearchFirst failed`)
	}

	if binarySearchFirst(nums, 3) != -1 {
		t.Error(`TestNcDichotomousBinarySearchFirst failed`)
	}
	nums = []int{1, 1, 1, 1, 1}
	if binarySearchFirst(nums, 1) != 0 {
		t.Error(`TestNcDichotomousBinarySearchFirst failed`)
	}
}

func TestNcDichotomousBinarySearchLast(t *testing.T) {
	nums := []int{1, 2, 4, 4, 5}
	if binarySearchLast(nums, 4) != 3 {
		t.Error(`TestNcDichotomousBinarySearchLast failed`)
	}

	if binarySearchLast(nums, 3) != -1 {
		t.Error(`TestNcDichotomousBinarySearchLast failed`)
	}
	nums = []int{1, 1, 1, 1, 1}
	if binarySearchLast(nums, 1) != 4 {
		t.Error(`TestNcDichotomousBinarySearchLast failed`)
	}
}

func TestNcDichotomousBinarySearchFirstGreaterEqual(t *testing.T) {
	nums := []int{1, 2, 4, 4, 5}
	if binarySearchFirstGreaterEqual(nums, 1) != 0 {
		t.Error(`TestNcDichotomousBinarySearchFirstGreaterEqual failed`)
	}

	if binarySearchFirstGreaterEqual(nums, 3) != 2 {
		t.Error(`TestNcDichotomousBinarySearchFirstGreaterEqual failed`)
	}

	if binarySearchFirstGreaterEqual(nums, 5) != 4 {
		t.Error(`TestNcDichotomousBinarySearchFirstGreaterEqual failed`)
	}

	if binarySearchFirstGreaterEqual(nums, 6) != -1 {
		t.Error(`TestNcDichotomousBinarySearchFirstGreaterEqual failed`)
	}

	nums = []int{1, 1, 1, 1, 1}
	if binarySearchFirstGreaterEqual(nums, 1) != 0 {
		t.Error(`TestNcDichotomousBinarySearchFirstGreaterEqual failed`)
	}
}

func TestNcDichotomousBinarySearchLastLowerEqual(t *testing.T) {
	nums := []int{1, 2, 4, 4, 5}
	if binarySearchLastLowerEqual(nums, 0) != -1 {
		t.Error(`TestNcDichotomousBinarySearchLastLowerEqual failed`)
	}

	if binarySearchLastLowerEqual(nums, 1) != 0 {
		t.Error(`TestNcDichotomousBinarySearchLastLowerEqual failed`)
	}

	if binarySearchLastLowerEqual(nums, 3) != 1 {
		t.Error(`TestNcDichotomousBinarySearchLastLowerEqual failed`)
	}

	if binarySearchLastLowerEqual(nums, 5) != 4 {
		t.Error(`TestNcDichotomousBinarySearchLastLowerEqual failed`)
	}

	if binarySearchLastLowerEqual(nums, 6) != 4 {
		t.Error(`TestNcDichotomousBinarySearchLastLowerEqual failed`)
	}

	nums = []int{1, 1, 1, 1, 1}
	if binarySearchLastLowerEqual(nums, 1) != 4 {
		t.Error(`TestNcDichotomousBinarySearchLastLowerEqual failed`)
	}
}
