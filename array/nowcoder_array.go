package array

/**

合并两个有序数组

描述
给出一个整数数组 和有序的整数数组 ，请将数组 合并到数组 中，变成一个有序的升序数组
注意：
1.可以假设 数组有足够的空间存放 数组的元素， 和 中初始的元素数目分别为 和 ，的数组空间大小为 +
2.不要返回合并的数组，返回是空的，将数组 的数据合并到里面就好了
3.数组在[0,m-1]的范围也是有序的

例1:
A: [4,5,6,0,0,0]，m=3
B: [1,2,3]，n=3
合并过后A为:
A: [1,2,3,4,5,6]

示例1
输入：
[4,5,6],[1,2,3]
返回值：
[1,2,3,4,5,6]
说明：
A数组为[4,5,6]，B数组为[1,2,3]，后台程序会预先将A扩容为[4,5,6,0,0,0]，B还是为[1,2,3]，m=3，n=3，传入到函数merge里面，然后请同学完成merge函数，将B的数据合并A里面，最后后台程序输出A数组

示例2
输入：
[1,2,3],[2,5,6]
返回值：
[1,2,2,3,5,6]

*/

/**
 *
 * @param A int整型一维数组
 * @param B int整型一维数组
 * @return void
 */
func merge(A []int, m int, B []int, n int) []int {
	idxA, idxB := 0, 0
	newA := make([]int, 0)
	for idxA < m && idxB < n {
		if A[idxA] > B[idxB] { // A大，则B前移
			newA = append(newA, B[idxB])
			idxB++
		} else { // 相反，则A前移
			newA = append(newA, A[idxA])
			idxA++
		}
	}

	if idxB < n { // b没遍历完，直接把后续加入到A后面
		newA = append(newA, B[idxB:n]...)
	}

	if idxA < m {
		newA = append(newA, A[idxA:m]...)
	}

	return newA
}

/**

两数之和

描述
给出一个整数数组，请在数组中找出两个加起来等于目标值的数，
你给出的函数twoSum 需要返回这两个数字的下标（index1，index2），需要满足 index1 小于index2.。注意：下标是从1开始的
假设给出的数组中只存在唯一解
例如：
给出的数组为 {20, 70, 110, 150},目标值为90
输出 index1=1, index2=2


示例1
输入：
[3,2,4],6
返回值：
[2,3]
说明：
因为 2+4=6 ，而 2的下标为2 ， 4的下标为3 ，又因为 下标2 < 下标3 ，所以输出[2,3]

*/
/**
 *
 * @param numbers int整型一维数组
 * @param target int整型
 * @return int整型一维数组
 */
func twoSum(numbers []int, target int) []int {
	dict := make(map[int]int)
	for i, n := range numbers {
		if value, ok := dict[target-n]; ok {
			return []int{value + 1, i + 1}
		} else {
			dict[n] = i
		}
	}
	return []int{}
}

/**

数字在升序数组中出现的次数

描述
统计一个数字在升序数组中出现的次数。
示例1
输入：
[1,2,3,3,3,3,4,5],3
返回值：
4
*/
/**
 *
 * @param data int整型一维数组
 * @param k int整型
 * @return int整型
 */
func GetNumberOfK(data []int, k int) int {
	num := 0
	idx := binarySearch(data, k)
	// idx := binarySearch2(data, 0, len(data)-1, k)
	if idx == -1 {
		return 0
	}

	num = 1
	for i := idx + 1; i < len(data); i++ {
		if data[i] != k {
			break
		} else {
			num++
		}
	}

	for i := idx - 1; i >= 0; i-- {
		if data[i] != k {
			break
		} else {
			num++
		}
	}

	return num
}

// 非递归二分查找
func binarySearch(arr []int, hkey int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == hkey {
			return mid
		} else if hkey < arr[mid] {
			high = mid - 1
		} else if hkey > arr[mid] {
			low = mid + 1
		}
	}
	return -1
}

// 递归二分查找
func binarySearch2(arr []int, low, high, hkey int) int {
	if low > high {
		return -1
	}
	mid := low + (high-low)/2

	if arr[mid] > hkey {
		return binarySearch2(arr, low, mid-1, hkey)
	} else if arr[mid] < hkey {
		return binarySearch2(arr, mid+1, high, hkey)
	}

	return mid
}

/**

数组中出现次数超过一半的数字

描述
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。你可以假设数组是非空的，并且给定的数组总是存在多数元素。1<=数组长度<=50000，0<=数组元素<=10000
示例1
输入：
[1,2,3,2,2,2,5,4,2]
复制
返回值：
2
复制
示例2
输入：
[3,3,3,3,2,2,2]
复制
返回值：
3
复制
示例3
输入：
[1]
复制
返回值：
1
复制
*/
/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
func MoreThanHalfNum(numbers []int) int {
	dict := make(map[int]int)
	max, times := 0, 0
	for _, n := range numbers {
		k, ok := dict[n]
		if ok {
			dict[n]++
		} else {
			dict[n] = 1
		}

		if k+1 > times {
			times = k + 1
			max = n
		}
	}

	return max
}

/**

三个数的最大乘积

描述
给定一个长度为  的无序数组  ，包含正数、负数和 0 ，请从中找出 3 个数，使得乘积最大，返回这个乘积。

要求时间复杂度： O(n) ，空间复杂度： O(1)。

数据范围：
3 <= n <= 2 * 10^5
-10^6<=A[i]<=10^6

示例1
输入：
[3,4,1,2]
复制
返回值：
24
复制

*/

/**
 * 最大乘积
 * @param A int整型一维数组
 * @return long长整型
 */
func solve(A []int) int64 {
	var top, mid, low int64
	for _, n := range A {
		n64 := int64(n)
		if n64 >= top {
			low = mid
			mid = top
			top = n64
		} else if n64 >= mid {
			low = mid
			mid = n64
		} else if n64 >= low {
			low = n64
		}
	}
	return top * mid * low
}
