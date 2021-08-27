package dichotomous

import "testing"

func TestNcDichotomousSqrt(t *testing.T) {
	if sqrt(0) != 0 || sqrt(1) != 1 || sqrt(2) != 1 || sqrt(3) != 1 || sqrt(4) != 2 {
		t.Error(`TestNcDichotomousSqrt failed`)
	}
}

func TestNcDichotomousSearch(t *testing.T) {
	nums := []int{4, 5, 6, 1, 2, 3}
	// if search(nums, 5) != 1 {
	// 	t.Error(`TestNcDichotomousSearch failed`)
	// }
	if search(nums, 2) != 4 {
		t.Error(`TestNcDichotomousSearch failed`)
	}

	nums = []int{5, 6, 1, 2, 3, 4}
	if search(nums, 5) != 0 {
		t.Error(`TestNcDichotomousSearch failed`)
	}
	if search(nums, 2) != 3 {
		t.Error(`TestNcDichotomousSearch failed`)
	}

	nums = []int{1}
	if search(nums, 1) != 0 {
		t.Error(`TestNcDichotomousSearch failed`)
	}
	if search(nums, 2) != -1 {
		t.Error(`TestNcDichotomousSearch failed`)
	}
}

func TestNcDichotomousMinNumberInRotateArray(t *testing.T) {
	nums := []int{3, 4, 5, 1, 2}
	if minNumberInRotateArray(nums) != 1 {
		t.Error(`TestNcDichotomousMinNumberInRotateArray failed`)
	}

	nums = []int{3, 2}
	if minNumberInRotateArray(nums) != 2 {
		t.Error(`TestNcDichotomousMinNumberInRotateArray failed`)
	}

	nums = []int{2, 3, 4, 0, 1}
	if minNumberInRotateArray(nums) != 0 {
		t.Error(`TestNcDichotomousMinNumberInRotateArray failed`)
	}

	nums = []int{1, 2, 2, 2, 2}
	if minNumberInRotateArray(nums) != 1 {
		t.Error(`TestNcDichotomousMinNumberInRotateArray failed`)
	}

	nums = []int{2, 2, 2, 1, 2}
	if minNumberInRotateArray(nums) != 1 {
		t.Error(`TestNcDichotomousMinNumberInRotateArray failed`)
	}
}

func TestNcDichotomousBinarySearch1(t *testing.T) {
	nums := []int{-1, 3, 4, 5, 6, 7}
	if binarySearch1(nums, -1) != 0 {
		t.Error(`TestNcDichotomousBinarySearch1 failed`)
	}

	if binarySearch1(nums, 2) != -1 {
		t.Error(`TestNcDichotomousBinarySearch1 failed`)
	}

	if binarySearch1(nums, 7) != 5 {
		t.Error(`TestNcDichotomousBinarySearch1 failed`)
	}
}

func TestNcDichotomousBinarySearch2(t *testing.T) {
	nums := []int{1, 2, 4, 4, 5}
	if binarySearch2(nums, 4) != 2 {
		t.Error(`TestNcDichotomousBinarySearch2 failed`)
	}

	if binarySearch2(nums, 3) != -1 {
		t.Error(`TestNcDichotomousBinarySearch2 failed`)
	}
	nums = []int{1, 1, 1, 1, 1}
	if binarySearch2(nums, 1) != 0 {
		t.Error(`TestNcDichotomousBinarySearch2 failed`)
	}
}
