package sort

/**
1.
熟练掌握冒泡、
插⼊、
选择、
快速、
归并排序的原理、代码实现；
2.
熟练掌握快速、
归并排序的时间和空间复杂度分析；
3.
掌握桶排序、
计数排序、
基数排序的原理；

线性排序
桶排序：所有元素尽量均匀分配到各个桶中。可用于限定内存下对大数据排序
计数排序：桶排序的特例，每个桶内元素值是相同的。可用于年龄排序、考分排序
基数排序：要求数据可以分割独立的“位”来比较，位之间有递进关系，每一位的数据范围不能太大，要可以使用线性排序（桶排序或计数排序）
*/

/**
输入：[5,2,3,1,4]
返回值：[1,2,3,4,5]

输入：[5,1,6,2,5]
返回值：[1,2,5,5,6]
*/

// 原地排序且稳定：每次冒泡可使至少一个元素移动到它应该在的位置，从大到小
func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		flag := false                // 提前退出冒泡循环的标志位
		for j := 0; j < n-i-1; j++ { // 每次冒泡完，可确定arr[n-i-1]的值
			if arr[j] > arr[j+1] { // 把大数后移
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return arr
}

// 原地排序且稳定：取未排序区间中的元素，在已排序区间中倒序找插入位置
func InsertSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- { // 从已排序区间，从大往小找
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
	return arr
}

// 原地排序但不稳定：每次会从未排序区间中找到最小的元素，将其放到已排序区间的末尾
func SelectSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min { // 当前值就是未排序元素中最小的
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

// 非原地排序但稳定：将大数组从中间分成两个小数组，小数组排序后，再把两个小数组合并
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left := MergeSort(arr[:len(arr)/2])
	right := MergeSort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(l, r []int) []int {
	final := []int{}
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			final = append(final, l[i])
			i++
		} else {
			final = append(final, r[j])
			j++
		}
	}

	for ; i < len(l); i++ {
		final = append(final, l[i])
	}

	for ; j < len(r); j++ {
		final = append(final, r[j])
	}
	return final
}

// 原地排序但不稳定：使用arr[lo]将大数组切分成两个小数组，左数组元素都小于等于arr[lo]，右数组元素都大于arr[lo]
func QuickSort(arr []int) []int {
	return QuickSortWithIndex(arr, 0, len(arr)-1)
}

func QuickSortWithIndex(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high) // 找出pivot
		arr = QuickSortWithIndex(arr, low, p-1)
		arr = QuickSortWithIndex(arr, p+1, high)
	}
	return arr
}

// 找出pivot，并保证pivot比左边都大，比右边都小或等于
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
