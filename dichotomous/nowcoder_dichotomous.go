package dichotomous

/**

NC32 求平方根

描述
实现函数 int sqrt(int x).
计算并返回x的平方根（向下取整）
示例1
输入：
2
返回值：
1
*/

/**
 *
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	var sqrt int
	for i := 0; i*i <= x; i++ {
		sqrt = i
	}
	return sqrt
}

/**

NC48 在旋转过的有序数组中寻找目标值

描述
给定一个整数数组nums，按升序排序，数组中的元素各不相同。
nums数组在传递给search函数之前，会在预先未知的某个下标 t（0 <= t <= nums.length-1）上进行旋转，让数组变为[nums[t], nums[t+1], ..., nums[nums.length-1], nums[0], nums[1], ..., nums[t-1]]。
比如，数组[0,2,4,6,8,10]在下标3处旋转之后变为[6,8,10,0,2,4]
现在给定一个旋转后的数组nums和一个整数target，请你查找这个数组是不是存在这个target，如果存在，那么返回它的下标，如果不存在，返回-1
示例1
输入：
[6,8,10,0,2,4],10
复制
返回值：
2
复制
示例2
输入：
[6,8,10,0,2,4],3
复制
返回值：
-1
复制
示例3
输入：
[2],1
复制
返回值：
-1
复制
备注：
1 <= nums.length <= 4000

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
*/
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] { // 左半升序
			if nums[mid] > target && nums[left] <= target { // 目标值在左边
				right = mid - 1
			} else { // 目标值不在左边
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] { //目标值大于中间值且小于最右边的值
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

/**

NC71 旋转数组的最小数字

描述
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序的数组的一个旋转，输出旋转数组的最小元素。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
示例1
输入：
[3,4,5,1,2]
复制
返回值：
1
复制
*/
/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
 */
func minNumberInRotateArray(rotateArray []int) int {
	/*
	   二分查找、右端点作为target

	   情况1，arr[mid] > target：4 5 6 1 2 3
	   arr[mid] 为 6， target为右端点 3， arr[mid] > target, 说明[first ... mid] 都是 >= target 的，因为原始数组是非递减，所以可以确定答案为 [mid+1...last]区间,所以 first = mid + 1

	   情况2，arr[mid] < target:5 6 1 2 3 4
	   arr[mid] 为 1， target为右端点 4， arr[mid] < target, 说明答案肯定不在[mid+1...last]，但是arr[mid] 有可能是答案,所以答案在[first, mid]区间，所以last = mid;

	   情况3，arr[mid] == target:
	   如果是 1 0 1 1 1， arr[mid] = target = 1, 显然答案在左边
	   如果是 1 1 1 0 1, arr[mid] = target = 1, 显然答案在右边
	   所以这种情况，不能确定答案在左边还是右边，那么就让last = last - 1;慢慢缩少区间，同时也不会错过答案。
	*/

	var len = len(rotateArray)
	if len == 0 {
		return 0
	}

	left := 0
	right := len - 1
	for left < right {
		mid := left + (right-left)/2
		if rotateArray[mid] < rotateArray[right] {
			right = mid
		} else if rotateArray[mid] > rotateArray[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return rotateArray[left]
}

/**

NC160 二分查找-I

描述
请实现无重复数字的升序数组的二分查找

给定一个 元素升序的、无重复数字的整型数组  和一个目标值  ，写一个函数搜索  中的 ，如果目标值存在返回下标（下标从 0 开始），否则返回 -1

数据范围：
0 \le len(nums) \le 10^40≤len(nums)≤10
4


复杂度要求：
时间
空间

示例1
输入：
[-1,0,3,4,6,10,13,14],13
复制
返回值：
6
复制
说明：
13 出现在nums中并且下标为 6
示例2
输入：
[],3
复制
返回值：
-1
复制
说明：
nums为空，返回-1
示例3
输入：
[-1,0,3,4,6,10,13,14],2
复制
返回值：
-1
复制
说明：
2 不存在nums中因此返回 -1
备注：
数组元素长度在[0,10000]之间
数组每个元素都在 [-9999, 9999]之间。
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarySearch1(nums []int, target int) int {
	len := len(nums)
	if len == 0 {
		return -1
	}
	left := 0
	right := len - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/*

NC105 二分查找-II

描述
请实现有重复数字的升序数组的二分查找
给定一个 元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的第一个出现的target，如果目标值存在返回下标，否则返回 -1
示例1
输入：
[1,2,4,4,5],4
复制
返回值：
2
复制
说明：
从左到右，查找到第1个为4的，下标为2，返回2
示例2
输入：
[1,2,4,4,5],3
复制
返回值：
-1
复制
示例3
输入：
[1,1,1,1,1],1
复制
返回值：
0
复制
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 如果目标值存在返回下标，否则返回 -1
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func binarySearch2(nums []int, target int) int {
	index := binarySearch1(nums, target)
	if index == -1 {
		return index
	}

	for i := 0; i < index; i++ {
		if nums[i] == target {
			index = i
			break
		}
	}
	return index
}
