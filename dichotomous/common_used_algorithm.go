package dichotomous

/**
二分查找及变体
变体⼀：查找第⼀个值等于给定值的元素
变体⼆：查找最后⼀个值等于给定值的元素
变体三：查找第⼀个⼤于等于给定值的元素
变体四：查找最后⼀个⼩于等于给定值的元素
*/

/**
 * 变体⼀：查找第⼀个值等于给定值的元素
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarySearchFirst(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != nums[mid] {
				return mid
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

/**
 * 变体⼆：查找最后⼀个值等于给定值的元素
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarySearchLast(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			if mid == high || nums[mid+1] != nums[mid] {
				return mid
			} else {
				low = mid + 1
			}
		}
	}
	return -1
}

/**
 * 变体三：查找第⼀个⼤于等于给定值的元素
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarySearchFirstGreaterEqual(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			} else {
				high = mid - 1
			}
		} else {
			low = mid + 1
		}
	}
	return -1
}

/**
 * 变体四：查找最后⼀个⼩于等于给定值的元素
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarySearchLastLowerEqual(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] <= target {
			if mid == high || nums[mid+1] > target { // mid已经是最后一个，或下一个元素大于target
				return mid
			} else {
				low = mid + 1
			}
		} else { // target在左边
			high = mid - 1
		}
	}
	return -1
}
