package nowcoder

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

/**
NC96 判断一个链表是否为回文结构
 算法知识视频讲解
简单  通过率：36.45%  时间限制：4秒  空间限制：256M
知识点
链表
双指针
题目
题解(49)
讨论(109)
排行
描述
给定一个链表，请判断该链表是否为回文结构。
回文是指该字符串正序逆序完全一致。

数据范围： 链表节点数 ，链表中每个节点的值满足
示例1
输入：
[1]
复制
返回值：
true
复制
示例2
输入：
[2,1]
复制
返回值：
false
复制
说明：
2->1
示例3
输入：
[1,2,2,1]
复制
返回值：
true
复制
说明：
1->2->2->1
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类 the head
 * @return bool布尔型
 */
func isPail(head *ListNode) bool {
	// write code here
	if head == nil || head.Next == nil {
		return true
	}
	/*
		  // 方法一：遍历存储
			maps := make([]int, 0)
			for head != nil {
				maps = append(maps, head.Val)
				head = head.Next
			}

			size := len(maps)
			for i := 0; i < size/2; i++ {
				if maps[i] != maps[size-1-i] {
					return false
				}
			} */

	// 方法二：反转后半部分链表
	slow, fast := head, head
	// 找中间结点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 奇数个节点
	if fast != nil {
		slow = slow.Next
	}

	// 链表反转
	slow = reverseList(slow)

	// 反转后slow从头和head比值
	fast = head
	for slow != nil {
		if slow.Val != fast.Val {
			return false
		}
		slow = slow.Next
		fast = fast.Next
	}

	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for head.Next != nil {
		temp := head.Next
		head.Next = temp.Next

		temp.Next = pre.Next
		pre.Next = temp
	}
	return dummy.Next
}

/**
NC35 最小编辑代价
 算法知识视频讲解
较难  通过率：36.24%  时间限制：3秒  空间限制：512M
知识点
字符串
动态规划
题目
题解(34)
讨论(48)
排行
描述
给定两个字符串str1和str2，再给定三个整数ic，dc和rc，分别代表插入、删除和替换一个字符的代价，请输出将str1编辑成str2的最小代价。

数据范围：，
要求：空间复杂度 ，时间复杂度
示例1
输入：
"abc","adc",5,3,2
复制
返回值：
2
复制
示例2
输入：
"abc","adc",5,3,100
复制
返回值：
8
复制
备注：
1≤∣str1∣,∣str2∣≤5000
1≤ic,dc,rc≤10000
*/
/**
 * min edit cost
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @param ic int整型 insert cost
 * @param dc int整型 delete cost
 * @param rc int整型 replace cost
 * @return int整型
 */
// func minEditCost(str1 string, str2 string, ic int, dc int, rc int) int {
// 	// write code here
// }

/**
NC5 二叉树根节点到叶子节点的所有路径和
 算法知识视频讲解
中等  通过率：35.45%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(38)
讨论(235)
排行
描述
给定一个仅包含数字\ 0-9 0−9 的二叉树，每一条从根节点到叶子节点的路径都可以用一个数字表示。
例如根节点到叶子节点的一条路径是1\to 2\to 31→2→3,那么这条路径就用\ 123 123 来代替。
找出根节点到叶子节点的所有路径表示的数字之和
例如：

这颗二叉树一共有两条路径，
根节点到叶子节点的路径 1\to 21→2 用数字\ 12 12 代替
根节点到叶子节点的路径 1\to 31→3 用数字\ 13 13 代替
所以答案为\ 12+13=25 12+13=25

数据范围：节点数 ,保证结果在32位整型范围内
要求：空间复杂度 ，时间复杂度
示例1
输入：
{1,0}
复制
返回值：
10
复制
示例2
输入：
{1,#,9}
复制
返回值：
19
复制
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
func sumNumbers(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}
	var sum = 0
	return preOrderSum(root, sum)
}

func preOrderSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	sum = 10*sum + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	return preOrderSum(root.Left, sum) + preOrderSum(root.Right, sum)
}

/**
NC8 二叉树根节点到叶子节点和为指定值的路径
 算法知识视频讲解
中等  通过率：32.23%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(59)
讨论(177)
排行
描述
给定一个节点数为 n 的二叉树和一个值 sum ，请找出所有的根节点到叶子节点的节点值之和等于的路径，如果没有则返回空。
例如：
给出如下的二叉树，sum = 22 ，

返回
[
[5,4,11,2],
[5,8,9]
]

数据范围：，每个节点的值  ，
要求： 空间复杂度 ，时间复杂度
示例1
输入：
{1,2},1
复制
返回值：
[]
复制
说明：
此树只有一条路径，即[1,2]大于sum，所以输出空。
示例2
输入：
{1,2},3
复制
返回值：
[[1,2]]
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @param sum int整型
 * @return int整型二维数组
 */

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res = [][]int{}
	path := []int{}

	dfsPathSum(root, sum, path)

	return res
}

func dfsPathSum(root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && root.Val == sum {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	dfsPathSum(root.Left, sum-root.Val, path)
	dfsPathSum(root.Right, sum-root.Val, path)
}

/**
NC21 链表内指定区间反转
 算法知识视频讲解
中等  通过率：33.27%  时间限制：1秒  空间限制：64M
知识点
链表
题目
题解(40)
讨论(150)
排行
描述
将一个节点数为 size 链表 m 位置到 n 位置之间的区间反转，要求时间复杂度 O(n)O(n)，空间复杂度 O(1)O(1)。
例如：
给出的链表为 1→2→3→4→5→NULL,m=2,n=4,
返回 1→4→3→2→5→NULL.

数据范围： 链表长度 ，，链表中每个节点的值满足
要求：时间复杂度  ，空间复杂度
进阶：时间复杂度 ，空间复杂度
示例1
输入：
{1,2,3,4,5},2,4
复制
返回值：
{1,4,3,2,5}
复制
示例2
输入：
{5},1,1
复制
返回值：
{5}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @param m int整型
 * @param n int整型
 * @return ListNode类
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	// write code here
	if head == nil || head.Next == nil || m == n {
		return head
	}

	// 增加头结点，防止m=1时pre为nil的情况
	dummy := &ListNode{Next: head}

	// 找出m前一个的结点
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}

	// 头插法反转m-n之间的结点
	cur := pre.Next // cur为m结点
	for i := m; i < n; i++ {
		// 把[m+1，n]之间结点一一摘除，头插到m-1和m之间
		temp := cur.Next
		cur.Next = temp.Next

		// 将temp插入pre和pre.Next之前
		temp.Next = pre.Next
		pre.Next = temp
	}

	return dummy.Next
}

/**
NC34 求路径
 算法知识视频讲解
简单  通过率：46.22%  时间限制：1秒  空间限制：64M
知识点
动态规划
题目
题解(35)
讨论(156)
排行
描述
一个机器人在m×n大小的地图的左上角（起点）。
机器人每次可以向下或向右移动。机器人要到达地图的右下角（终点）。
可以有多少种不同的路径从起点走到终点？

备注：m和n小于等于100,并保证计算结果在int范围内

数据范围：，保证计算结果在32位整型范围内
要求：空间复杂度 ，时间复杂度
进阶：空间复杂度 ，时间复杂度
示例1
输入：
2,1
复制
返回值：
1
复制
示例2
输入：
2,2
复制
返回值：
2
*/
/**
 *
 * @param m int整型
 * @param n int整型
 * @return int整型
 */
func uniquePaths(m int, n int) int {
	// write code here
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

/**
NC37 合并区间
 算法知识视频讲解
中等  通过率：30.31%  时间限制：1秒  空间限制：64M
知识点
排序
数组
题目
题解(48)
讨论(160)
排行
描述
给出一组区间，请合并所有重叠的区间。
请保证合并后的区间按区间起点升序排列。

数据范围：区间组数 ，区间内 的值都满足
要求：空间复杂度 ，时间复杂度
进阶：空间复杂度 ，时间复杂度
示例1
输入：
[[10,30],[20,60],[80,100],[150,180]]
复制
返回值：
[[10,60],[80,100],[150,180]]
复制
示例2
输入：
[[0,10],[10,20]]
复制
返回值：
[[0,20]]
*/

type Interval struct {
	Start int
	End   int
}

/**
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func mergeInterval(intervals []*Interval) []*Interval {
	// write code here
	size := len(intervals)
	var ret []*Interval

	// 先排序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// 再通过二分
	i := 0
	for i < size {
		left := intervals[i].Start
		right := intervals[i].End
		for i < size-1 && intervals[i+1].Start <= right { // 判断下一个区间是否重合，重合取较大的做right
			right = int(math.Max(float64(right), float64(intervals[i+1].End)))
			i++
		}
		ret = append(ret, &Interval{left, right})
		i++
	}

	return ret
}

func SPrintInterval(intervals []*Interval) string {
	size := len(intervals)
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		tmp := make([]int, 2)
		tmp[0] = intervals[i].Start
		tmp[1] = intervals[i].End
		dp[i] = tmp
	}
	return fmt.Sprint(dp)
}

/**
NC92 最长公共子序列-II
 算法知识视频讲解
中等  通过率：36.81%  时间限制：3秒  空间限制：256M
知识点
动态规划
题目
题解(37)
讨论(86)
排行
描述
给定两个字符串str1和str2，输出两个字符串的最长公共子序列。如果最长公共子序列为空，则返回"-1"。目前给出的数据，仅仅会存在一个最长的公共子序列

数据范围：
要求：空间复杂度  ，时间复杂度
示例1
输入：
"1A2C3D4B56","B1D23A456A"
复制
返回值：
"123456"
复制
示例2
输入：
"abc","def"
复制
返回值：
"-1"
复制
示例3
输入：
"abc","abc"
复制
返回值：
"abc"
复制
示例4
输入：
"ab",""
复制
返回值：
"-1"
*/
/**
 * longest common subsequence
 * @param s1 string字符串 the string
 * @param s2 string字符串 the string
 * @return string字符串
 */
func LCSPlus1(s1 string, s2 string) string {
	// write code here
	m, n := len(s1), len(s2)
	dp := make([][]string, m+1)
	for i := range dp {
		dp[i] = make([]string, n+1)
	}
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + string(c1)
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	if len(dp[m][n]) == 0 {
		return "-1"
	}
	return dp[m][n]
}

func max(a, b string) string {
	if len(a) > len(b) {
		return a
	}
	return b
}

/**
NC36 在两个长度相等的排序数组中找到上中位数
 算法知识视频讲解
较难  通过率：42.40%  时间限制：2秒  空间限制：256M
知识点
数组
二分
分治
题目
题解(23)
讨论(55)
排行
描述
给定两个递增数组arr1和arr2，已知两个数组的长度都为N，求两个数组中所有数的上中位数。
上中位数：假设递增序列长度为n，若n为奇数，则上中位数为第n/2+1个数；否则为第n/2个数

数据范围：，

要求：时间复杂度O(N) ，空间复杂度O(1)
进阶：时间复杂度为O(logN)，空间复杂度为O(1)

示例1
输入：
[1,2,3,4],[3,4,5,6]
复制
返回值：
3
复制
说明：
总共有8个数，上中位数是第4小的数，所以返回3。
示例2
输入：
[0,1,2],[3,4,5]
复制
返回值：
2
复制
说明：
总共有6个数，那么上中位数是第3小的数，所以返回2
示例3
输入：
[1],[2]
复制
返回值：
1
复制
备注：
1≤N≤10^5
0≤arr_1i,arr_2i≤10^9
*/
/**
 * find median in two sorted array
 * @param arr1 int整型一维数组 the array1
 * @param arr2 int整型一维数组 the array2
 * @return int整型
 */
func findMedianinTwoSortedAray(arr1 []int, arr2 []int) int {
	// write code here
	n := len(arr1)
	i, j := 0, 0
	var lastItem int
	for i < n {
		if i+j == n {
			break
		}
		if arr1[i] < arr2[j] {
			lastItem = arr1[i]
			i++
		} else {
			lastItem = arr2[j]
			j++
		}
	}
	return lastItem
}

/**
NC60 判断一棵二叉树是否为搜索二叉树和完全二叉树
 算法知识视频讲解
中等  通过率：29.02%  时间限制：5秒  空间限制：256M
知识点
树
dfs
题目
题解(30)
讨论(94)
排行
描述
给定一棵二叉树，已知其中的节点没有重复值，请判断该二叉树是否为搜索二叉树和完全二叉树。
输出描述：分别输出是否为搜索二叉树、完全二叉树。


数据范围：二叉树节点数满足  ，二叉树上的值满足
要求：空间复杂度 ，时间复杂度

注意：空子树我们认为同时符合搜索二叉树和完全二叉树。
示例1
输入：
{2,1,3}
复制
返回值：
[true,true]
复制
示例2
输入：
{1,#,2}
复制
返回值：
{true,false}
复制
说明：
由于节点的右儿子大于根节点，无左子树，所以是搜索二叉树但不是完全二叉树
示例3
输入：
{}
复制
返回值：
{true,true}
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类 the root
 * @return bool布尔型一维数组
 */
// func judgeIt( root *TreeNode ) []bool {
//     // write code here
// }

/**
NC24 删除有序链表中重复的元素-II
 算法知识视频讲解
中等  通过率：34.21%  时间限制：1秒  空间限制：64M
知识点
链表
题目
题解(51)
讨论(181)
排行
描述
给出一个升序排序的链表，删除链表中的所有重复出现的元素，只保留原链表中只出现一次的元素。
例如：
给出的链表为1 \to 2\to 3\to 3\to 4\to 4\to51→2→3→3→4→4→5, 返回1\to 2\to51→2→5.
给出的链表为1\to1 \to 1\to 2 \to 31→1→1→2→3, 返回2\to 32→3.

数据范围：链表长度 ，链表中的值满足
要求：空间复杂度 ，时间复杂度
进阶：空间复杂度 ，时间复杂度



示例1
输入：
{1,2,2}
复制
返回值：
{1}
复制
示例2
输入：
{}
复制
返回值：
{}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @return ListNode类
 */
// func deleteDuplicates( head *ListNode ) *ListNode {
//     // write code here
// }

/**
NC100 将字符串转化为整数
 算法知识视频讲解
较难  通过率：19.37%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(24)
讨论(115)
排行
描述
实现函数 atoi 。函数的功能为将字符串转化为整数
提示：仔细思考所有可能的输入情况。这个问题没有给出输入的限制，你需要自己考虑所有可能的情况。


atoi函数位于cpp <stdlib.h>库。该函数舍弃字符串前全部空白符，直至找到一个合法的数字或者正负号后开始读取，然后尽可能读取多的字符组成合法的整数表示，并返回这个合法的整数，值得注意的是，当读取到一个非法字符后将直接停止读取数字并返回当前已转换结果。注意，数学中e表示数的幂在该函数定义中是非法字符。当输入数字超出返回数据类型范围是一个未定义行为（undefined behavior），本题保证不会出现这种情况。

数据范围：字符串长度满足
要求：空间复杂度 ，时间复杂度
示例1
输入：
"123"
复制
返回值：
123
复制
示例2
输入：
"123e123"
复制
返回值：
"123"
复制
示例3
输入：
"e123"
复制
返回值：
"0"
*/
/**
 *
 * @param str string字符串
 * @return int整型
 */
// func atoi(str string) int {
// 	// write code here
// }

/**
NC57 反转数字
 算法知识视频讲解
简单  通过率：48.35%  时间限制：1秒  空间限制：64M
知识点
数学
题目
题解(30)
讨论(174)
排行
描述
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
你有注意到翻转后的整数可能溢出吗？因为给出的是32位整数，则其数值范围为[−2^{31}, 2^{31} − 1][−2
31
 ,2
31
 −1]。翻转可能会导致溢出，如果反转后的结果会溢出就返回 0。

数据范围：
要求：空间复杂度 ，时间复杂度
示例1
输入：
12
复制
返回值：
21
复制
示例2
输入：
-123
复制
返回值：
-321
复制
示例3
输入：
10
复制
返回值：
1
复制
示例4
输入：
1147483649
复制
返回值：
0
*/
/**
 *
 * @param x int整型
 * @return int整型
 */
func reverseInt(x int) int {
	// write code here
	ret := 0
	for x/10 != 0 {
		if checkOverflow(10*ret + x%10) {
			return 0
		}
		ret = 10*ret + x%10
		x = x / 10
	}

	if checkOverflow(10*ret + x%10) {
		return 0
	}

	ret = 10*ret + x%10

	return ret
}

func checkOverflow(v int) bool {
	return v > math.MaxInt32 || v < math.MinInt32
}

/**
NC86 矩阵元素查找
 算法知识视频讲解
中等  通过率：45.55%  时间限制：3秒  空间限制：64M
知识点
二分
分治
题目
题解(19)
讨论(89)
排行
描述
已知int一个有序矩阵mat，同时给定矩阵的大小n和m以及需要查找的元素x，且矩阵的行和列都是从小到大有序的。设计查找算法返回所查找元素的二元数组，代表该元素的行号和列号(均从零开始)。保证元素互异。

数据范围：，矩阵中的任何元素满足
要求：空间复杂度 ，时间复杂度


示例1
输入：
[[1,2,3],[4,5,6]],2,3,6
复制
返回值：
[1,2]
复制
示例2
输入：
[[1,2,3]],1,3,2
复制
返回值：
[0,1]
*/
/**
 *
 * @param mat int整型二维数组
 * @param n int整型
 * @param m int整型
 * @param x int整型
 * @return int整型一维数组
 */
func findElement(mat [][]int, n int, m int, x int) []int {
	// write code here
	// 横向二分
	pos, left, right := 0, 0, m-1
	for left <= right {
		mid := (right + left) >> 1
		if mat[0][mid] == x {
			return []int{0, mid}
		}

		if mat[0][mid] < x {
			pos = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// 纵向二分
	top, bottom := 0, n-1
	for top <= bottom {
		mid := (top + bottom) >> 1
		if mat[mid][pos] == x {
			return []int{mid, pos}
		}

		if mat[mid][pos] < x {
			top = mid + 1
		} else {
			bottom = mid - 1
		}
	}

	return []int{0, 0}
}

/**
NC30 缺失的第一个正整数
 算法知识视频讲解
中等  通过率：57.81%  时间限制：1秒  空间限制：256M
题目
题解
讨论(2)
排行
描述
给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数

进阶:你能用O(n)时间复杂度并且O(1)的空间复杂度的方法实现吗

数据范围:
-231<=nums[i]<=231-1
0<=len(nums)<=5*105
示例1
输入：
[1,0,2]
复制
返回值：
3
复制
示例2
输入：
[-2,3,4,1,5]
复制
返回值：
2
复制
示例3
输入：
[4,5,6,8,9]
复制
返回值：
1
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
// func minNumberDisappeared(nums []int) int {
// 	// write code here
// }

/**
NC101 缺失数字
 算法知识视频讲解
入门  通过率：44.27%  时间限制：1秒  空间限制：64M
知识点
位运算
数组
数学
二分
题目
题解(48)
讨论(111)
排行
描述
从 0,1,2,...,n 这 n+1 个数中选择 n 个数，选择出的数字依然保持有序，找出这 n 个数中缺失的那个数，要求 O(n) 或 O(log(n)) 并尽可能小。

数据范围：
要求： 空间复杂度 ，时间复杂度
进阶：空间复杂度，时间复杂度
示例1
输入：
[0,1,2,3,4,5,7]
复制
返回值：
6
复制
示例2
输入：
[0,2,3]
复制
返回值：
1
复制
示例3
输入：
[0,1,2,3,4]
复制
返回值：
5
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 找缺失数字
 * @param a int整型一维数组 给定的数字串
 * @return int整型
 */
// func solve(a []int) int {
// 	// write code here
// }

/**
NC133 链表的奇偶重排
 算法知识视频讲解
中等  通过率：52.04%  时间限制：1秒  空间限制：256M
知识点
链表
排序
题目
题解(33)
讨论(52)
排行
描述
给定一个单链表，请设定一个函数，将链表的奇数位节点和偶数位节点分别放在一起，重排后输出。
注意是节点的编号而非节点的数值。

数据范围：节点数量满足 ，节点中的值都满足
要求：空间复杂度 ，时间复杂度
示例1
输入：
{1,2,3,4,5,6}
复制
返回值：
{1,3,5,2,4,6}
复制
说明：
1->2->3->4->5->6->NULL
重排后为
1->3->5->2->4->6->NULL

示例2
输入：
{1,4,6,3,7}
复制
返回值：
{1,6,7,4,3}
复制
说明：
1->4->6->3->7->NULL
重排后为
1->6->7->4->3->NULL
奇数位节点有1,6,7，偶数位节点有4,3。重排后为1,6,7,4,3

备注：
链表长度不大于200000。每个数范围均在int内。
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param head ListNode类
 * @return ListNode类
 */
// func oddEvenList( head *ListNode ) *ListNode {
//     // write code here
// }

/**
NC6 二叉树的最大路径和
 算法知识视频讲解
较难  通过率：30.40%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(30)
讨论(105)
排行
描述
给定一个二叉树，请计算节点值之和最大的路径的节点值之和是多少。
这个路径的开始节点和结束节点可以是二叉树中的任意节点
例如：
给出以下的二叉树，

返回的结果为6

数据范围：节点数满足  ，节点上的值满足
要求：空间复杂度 ，时间复杂度
示例1
输入：
{-2,1}
复制
返回值：
1
复制
示例2
输入：
{-2,#,-3}
复制
返回值：
-2
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return int整型
 */
// func maxPathSum( root *TreeNode ) int {
//     // write code here
// }

/**
NC16 判断二叉树是否对称
 算法知识视频讲解
简单  通过率：41.40%  时间限制：1秒  空间限制：64M
知识点
树
dfs
bfs
题目
题解(41)
讨论(178)
排行
描述
给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）
例如：                                 下面这棵二叉树是对称的

                                              下面这棵二叉树不对称。

数据范围：节点数满足 ，节点上的值满足
要求：空间复杂度 ，时间复杂度
备注：
你可以用递归和迭代两种方法解决这个问题
示例1
输入：
{1,2,2}
复制
返回值：
true
复制
示例2
输入：
{1,2,3,3,#,2,#}
复制
返回值：
false
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @return bool布尔型
 */
// func isSymmetric(root *TreeNode) bool {
// 	// write code here
// }

/**
NC26 括号生成
 算法知识视频讲解
中等  通过率：49.79%  时间限制：1秒  空间限制：64M
知识点
回溯
题目
题解(23)
讨论(84)
排行
描述
给出n对括号，请编写一个函数来生成所有的由n对括号组成的合法组合。
例如，给出n=3，解集为：
"((()))", "(()())", "(())()", "()()()", "()(())"

 数据范围：
 要求：空间复杂度 ，时间复杂度
示例1
输入：
1
复制
返回值：
["()"]
复制
示例2
输入：
2
复制
返回值：
["(())","()()"]
*/
/**
 *
 * @param n int整型
 * @return string字符串一维数组
 */
// func generateParenthesis( n int ) []string {
//     // write code here
// }

/**
NC18 顺时针旋转矩阵
 算法知识视频讲解
中等  通过率：45.26%  时间限制：3秒  空间限制：64M
知识点
数组
题目
题解(36)
讨论(224)
排行
描述
有一个NxN整数矩阵，请编写一个算法，将矩阵顺时针旋转90度。

给定一个NxN的矩阵，和矩阵的阶数N,请返回旋转后的NxN矩阵。

数据范围：，矩阵中的值满足

要求：空间复杂度 ，时间复杂度
进阶：空间复杂度 ，时间复杂度
示例1
输入：
[[1,2,3],[4,5,6],[7,8,9]],3
复制
返回值：
[[7,4,1],[8,5,2],[9,6,3]]
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param mat int整型二维数组
 * @param n int整型
 * @return int整型二维数组
 */
// func rotateMatrix(mat [][]int, n int) [][]int {
// 	// write code here
// }

/**
NC20 数字字符串转化成IP地址
 算法知识视频讲解
中等  通过率：33.83%  时间限制：1秒  空间限制：64M
知识点
字符串
回溯
题目
题解(14)
讨论(107)
排行
描述
现在有一个只包含数字的字符串，将该字符串转化成IP地址的形式，返回所有可能的情况。
例如：
给出的字符串为"25525522135",
返回["255.255.22.135", "255.255.221.35"]. (顺序没有关系)

数据范围：字符串长度
要求：空间复杂度 ,时间复杂度

注意：ip地址是由四段数字组成的数字序列，格式如 "x.x.x.x"，其中 x 的范围应当是 [0,255]。

示例1
输入：
"25525522135"
复制
返回值：
["255.255.22.135","255.255.221.35"]
复制
示例2
输入：
"1111"
复制
返回值：
["1.1.1.1"]
复制
示例3
输入：
"000256"
复制
返回值：
"[]"
*/
/**
 *
 * @param s string字符串
 * @return string字符串一维数组
 */
// func restoreIpAddresses( s string ) []string {
//     // write code here
// }

/**
NC94 LFU缓存结构设计
 算法知识视频讲解
较难  通过率：35.85%  时间限制：2秒  空间限制：256M
知识点
模拟
题目
题解(18)
讨论(30)
排行
描述
一个缓存结构需要实现如下功能。
set(key, value)：将记录(key, value)插入该结构
get(key)：返回key对应的value值
但是缓存结构中最多放K条记录，如果新的第K+1条记录要加入，就需要根据策略删掉一条记录，然后才能把新记录加入。这个策略为：在缓存结构的K条记录中，哪一个key从进入缓存结构的时刻开始，被调用set或者get的次数最少，就删掉这个key的记录；
如果调用次数最少的key有多个，上次调用发生最早的key被删除
这就是LFU缓存替换算法。实现这个结构，K作为参数给出

数据范围：，
要求：get和set的时间复杂度都是 ，空间复杂度是


若opt=1，接下来两个整数x, y，表示set(x, y)
若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1

对于每个操作2，返回一个答案
示例1
输入：
[[1,1,1],[1,2,2],[1,3,2],[1,2,4],[1,3,5],[2,2],[1,4,4],[2,1]],3
复制
返回值：
[4,-1]
复制
说明：
在执行"1 4 4"后，"1 1 1"被删除。因此第二次询问的答案为-1
备注：
1≤k≤n≤10^5
−2×10^9≤x,y≤2×10^9
*/
/**
 * lfu design
 * @param operators int整型二维数组 ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
// func LFU( operators [][]int ,  k int ) []int {
//     // write code here
// }

/**
描述
将给定的单链表\ L L： L_0→L_1→…→L_{n-1}→L_ nL
0
​
 →L
1
​
 →…→L
n−1
​
 →L
n
​

重新排序为：L_0→L_n →L_1→L_{n-1}→L_2→L_{n-2}→…L
0
​
 →L
n
​
 →L
1
​
 →L
n−1
​
 →L
2
​
 →L
n−2
​
 →…
要求使用原地算法，不能只改变节点内部的值，需要对实际的节点进行交换。

数据范围：链表长度  ，链表中每个节点的值满足

要求：空间复杂度  并在链表上进行操作而不新建链表，时间复杂度
进阶：空间复杂度  ， 时间复杂度
示例1
输入：
{1,2,3,4}
复制
返回值：
{1,4,2,3}
复制
说明：
给定head链表1->2->3->4, 重新排列为 1->4->2->3,会取head链表里面的值打印输出
示例2
输入：
{1,2,3,4,5}
复制
返回值：
{1,5,2,4,3}
复制
说明：
给定head链表1->2->3->4->5, 重新排列为 1->5>2->4->3,会取head链表里面的值打印输出
示例3
输入：
{}
复制
返回值：
{}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @return void
 */
// func reorderList( head *ListNode )  {
//     // write code here
// }

/**
NC25 删除有序链表中重复的元素-I
 算法知识视频讲解
简单  通过率：40.00%  时间限制：1秒  空间限制：64M
知识点
链表
题目
题解(36)
讨论(150)
排行
描述
删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
例如：
给出的链表为1\to1\to21→1→2,返回1 \to 21→2.
给出的链表为1\to1\to 2 \to 3 \to 31→1→2→3→3,返回1\to 2 \to 31→2→3.

数据范围：链表长度满足 ，链表中任意节点的值满足
进阶：空间复杂度 ，时间复杂度
示例1
输入：
{1,1,2}
复制
返回值：
{1,2}
复制
示例2
输入：
{}
复制
返回值：
{}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类
 * @return ListNode类
 */
// func deleteDuplicates( head *ListNode ) *ListNode {
//     // write code here
// }

/**
NC42 有重复项数字的所有排列
 算法知识视频讲解
中等  通过率：39.88%  时间限制：1秒  空间限制：64M
知识点
回溯
题目
题解(17)
讨论(85)
排行
描述
给出一组可能包含重复项的数字，返回该组数字的所有排列。
示例1
输入：
[1,1,2]
复制
返回值：
[[1,1,2],[1,2,1],[2,1,1]]
*/
/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
// func permuteUnique(num []int) [][]int {
// 	// write code here
// }

/**
NC82 滑动窗口的最大值
 算法知识视频讲解
较难  通过率：25.95%  时间限制：1秒  空间限制：64M
知识点
堆
双指针
队列
题目
题解(97)
讨论(1k)
排行
描述
给定一个长度为 n 的数组 num 和滑动窗口的大小 size ，找出所有滑动窗口里数值的最大值。

例如，如果输入数组{2,3,4,2,6,2,5,1}及滑动窗口的大小3，那么一共存在6个滑动窗口，他们的最大值分别为{4,4,6,6,6,5}； 针对数组{2,3,4,2,6,2,5,1}的滑动窗口有以下6个： {[2,3,4],2,6,2,5,1}， {2,[3,4,2],6,2,5,1}， {2,3,[4,2,6],2,5,1}， {2,3,4,[2,6,2],5,1}， {2,3,4,2,[6,2,5],1}， {2,3,4,2,6,[2,5,1]}。

窗口大于数组长度或窗口长度为0的时候，返回空。

数据范围： ，，数组中每个元素的值满足
要求：空间复杂度 ，时间复杂度


示例1
输入：
[2,3,4,2,6,2,5,1],3
复制
返回值：
[4,4,6,6,6,5]
复制
示例2
输入：
[9,10,9,-7,-3,8,2,-6],5
复制
返回值：
[10,10,9,8]
复制
示例3
输入：
[1,2,3,4],5
复制
返回值：
[]
*/
/**
 *
 * @param num int整型一维数组
 * @param size int整型
 * @return int整型一维数组
 */
// func maxInWindows( num []int ,  size int ) []int {
//     // write code here
// }

/**
NC9 二叉树中是否存在节点和为指定值的路径
 算法知识视频讲解
简单  通过率：38.04%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(26)
讨论(137)
排行
描述
给定一个二叉树和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径，
例如：
给出如下的二叉树，\ sum=22 sum=22，

返回true，因为存在一条路径 5\to 4\to 11\to 25→4→11→2的节点值之和为 22

数据范围：树上的节点数满足  ，每 个节点的值都满足
进阶：空间复杂度 ，时间复杂度
示例1
输入：
{1,2},0
复制
返回值：
false
复制
示例2
输入：
{1,2},3
复制
返回值：
true
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param root TreeNode类
 * @param sum int整型
 * @return bool布尔型
 */
// func hasPathSum( root *TreeNode ,  sum int ) bool {
//     // write code here
// }

/**
NC46 加起来和为目标值的组合
 算法知识视频讲解
中等  通过率：35.23%  时间限制：1秒  空间限制：64M
知识点
数组
回溯
题目
题解(23)
讨论(85)
排行
描述
给出一组候选数 c 和一个目标数 t ，找出候选数中起来和等于 t 的所有组合。

 c 中的每个数字在一个组合中只能使用一次。

注意：
1. 题目中所有的数字（包括目标数 t  ）都是正整数
2. 组合中的数字 ( a_1, a_2, … , a_ka
1
​
 ,a
2
​
 ,…,a
k
​
  ) 要按非递减排序 ( a_1 \leq a_2 \leq … \leq a_ka
1
​
 ≤a
2
​
 ≤…≤a
k
​
  ).
3. 结果中不能包含重复的组合
4. 组合之间的排序按照索引从小到大依次比较，小的排在前面，如果索引相同的情况下数值相同，则比较下一个索引。

数据范围:  ， ，
要求：空间复杂度  ， 时间复杂度
示例1
输入：
[100,10,20,70,60,10,50],80
复制
返回值：
[[10,10,60],[10,20,50],[10,70],[20,60]]
复制
说明：
给定的候选数集是[100,10,20,70,60,10,50]，目标数是80
示例2
输入：
[2],1
复制
返回值：
[]
*/
/**
 *
 * @param num int整型一维数组
 * @param target int整型
 * @return int整型二维数组
 */
// func combinationSum2( num []int ,  target int ) [][]int {
//     // write code here
// }

/**
NC49 最长的括号子串
 算法知识视频讲解
较难  通过率：26.89%  时间限制：1秒  空间限制：64M
知识点
字符串
动态规划
题目
题解(13)
讨论(66)
排行
描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

例1: 对于字符串 "(()" 来说，最长的格式正确的子串是 "()" ，长度为 2 .
例2：对于字符串 ")()())" , 来说, 最长的格式正确的子串是 "()()" ，长度为 4 .

字符串长度：

要求时间复杂度  ,空间复杂度 .
示例1
输入：
"(()"
复制
返回值：
2
复制
示例2
输入：
"(())"
复制
返回值：
4
*/
/**
 *
 * @param s string字符串
 * @return int整型
 */
// func longestValidParentheses(s string) int {
// 	// write code here
// }

/**
NC55 最长公共前缀
 算法知识视频讲解
简单  通过率：36.42%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(35)
讨论(126)
排行
描述
给你一个大小为 n 的字符串数组 strs ，其中包含n个字符串 , 编写一个函数来查找字符串数组中的最长公共前缀，返回这个公共前缀。

数据范围： ，
进阶：空间复杂度 ，时间复杂度
示例1
输入：
["abca","abc","abca","abc","abcc"]
复制
返回值：
"abc"
复制
示例2
输入：
["abc"]
复制
返回值：
“abc"
*/
/**
 *
 * @param strs string字符串一维数组
 * @return string字符串
 */
// func longestCommonPrefix( strs []string ) string {
//     // write code here
// }

/**
NC56 回文数字
 算法知识视频讲解
简单  通过率：45.74%  时间限制：1秒  空间限制：64M
题目
题解(19)
讨论(104)
排行
描述
在不使用额外的内存空间的条件下判断一个整数是否是回文。
回文指逆序和正序完全相同。


数据范围：
进阶： 空间复杂度 ，时间复杂度

提示：
负整数可以是回文吗？（比如-1）
如果你在考虑将数字转化为字符串的话，请注意一下不能使用额外空间的限制
你可以将整数翻转。但是，如果你做过题目“反转数字”，你会知道将整数翻转可能会出现溢出的情况，你怎么处理这个问题？
示例1
输入：
121
复制
返回值：
true
复制
示例2
输入：
122
复制
返回值：
false
*/
/**
 *
 * @param x int整型
 * @return bool布尔型
 */
// func isPalindrome( x int ) bool {
//     // write code here
// }

/**
NC105 二分查找-II
 算法知识视频讲解
中等  通过率：30.56%  时间限制：1秒  空间限制：256M
知识点
二分
题目
题解(100)
讨论(143)
排行
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
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 如果目标值存在返回下标，否则返回 -1
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
// func search( nums []int ,  target int ) int {
//     // write code here
// }

/**
NC87 丢棋子问题
 算法知识视频讲解
中等  通过率：24.24%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(10)
讨论(19)
排行
描述

一座大楼有 n+1 层，地面算作第0层，最高的一层为第 n 层。已知棋子从第0层掉落肯定不会摔碎，从第 i 层掉落可能会摔碎，也可能不会摔碎。

给定整数 n 作为楼层数，再给定整数 k 作为棋子数，返回如果想找到棋子不会摔碎的最高层数，即使在最差的情况下扔的最小次数。一次只能扔一个棋子。

数据范围：
要求：空间复杂度 ，时间复杂度 (m是最终返回的结果)
示例1
输入：
10,1
复制
返回值：
10
复制
说明：
因为只有1棵棋子，所以不得不从第1层开始一直试到第10层，在最差的情况下，即第10层是不会摔坏的最高层，最少也要扔10次
示例2
输入：
3,2
复制
返回值：
2
复制
说明：
先在2层扔1棵棋子，如果碎了，试第1层，如果没碎，试第3层
示例3
输入：
105,2
复制
返回值：
14
复制
说明：
第一个棋子先在14层扔，碎了则用仅存的一个棋子试1~13层
若没碎，第一个棋子继续在27层扔，碎了则用仅存的一个棋子试15~26层
若没碎，第一个棋子继续在39层扔，碎了则用仅存的一个棋子试28~38层
若没碎，第一个棋子继续在50层扔，碎了则用仅存的一个棋子试40~49层
若没碎，第一个棋子继续在60层扔，碎了则用仅存的一个棋子试51~59层
若没碎，第一个棋子继续在69层扔，碎了则用仅存的一个棋子试61~68层
若没碎，第一个棋子继续在77层扔，碎了则用仅存的一个棋子试70~76层
若没碎，第一个棋子继续在84层扔，碎了则用仅存的一个棋子试78~83层
若没碎，第一个棋子继续在90层扔，碎了则用仅存的一个棋子试85~89层
若没碎，第一个棋子继续在95层扔，碎了则用仅存的一个棋子试91~94层
若没碎，第一个棋子继续在99层扔，碎了则用仅存的一个棋子试96~98层
若没碎，第一个棋子继续在102层扔，碎了则用仅存的一个棋子试100、101层
若没碎，第一个棋子继续在104层扔，碎了则用仅存的一个棋子试103层
若没碎，第一个棋子继续在105层扔，若到这一步还没碎，那么105便是结果
备注：
0≤N,K≤10^60≤N,K≤10
6
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回最差情况下扔棋子的最小次数
 * @param n int整型 楼层数
 * @param k int整型 棋子数
 * @return int整型
 */
// func solve( n int ,  k int ) int {
//     // write code here
// }

/**
NC123 序列化二叉树
 算法知识视频讲解
较难  通过率：23.60%  时间限制：1秒  空间限制：64M
知识点
队列
树
题目
题解(64)
讨论(509)
排行
描述
请实现两个函数，分别用来序列化和反序列化二叉树，不对序列化之后的字符串进行约束，但要求能够根据序列化之后的字符串重新构造出一棵与原二叉树相同的树。

二叉树的序列化(Serialize)是指：把一棵二叉树按照某种遍历方式的结果以某种格式保存为字符串，从而使得内存中建立起来的二叉树可以持久保存。序列化可以基于先序、中序、后序、层序的二叉树等遍历方式来进行修改，序列化的结果是一个字符串，序列化时通过 某种符号表示空节点（#）

二叉树的反序列化(Deserialize)是指：根据某种遍历顺序得到的序列化字符串结果str，重构二叉树。

例如，可以根据层序遍历的方案序列化，如下图:

层序序列化(即用函数Serialize转化)如上的二叉树转为"{1,2,3,#,#,6,7}"，再能够调用反序列化(Deserialize)将"{1,2,3,#,#,6,7}"构造成如上的二叉树。

当然你也可以根据满二叉树结点位置的标号规律来序列化，还可以根据先序遍历和中序遍历的结果来序列化。不对序列化之后的字符串进行约束，所以欢迎各种奇思妙想。

数据范围：节点数 ，树上每个节点的值满足
要求：序列化和反序列化都是空间复杂度 ，时间复杂度
示例1
输入：
{1,2,3,#,#,6,7}
复制
返回值：
{1,2,3,#,#,6,7}
复制
说明：
如题面图
示例2
输入：
{8,6,10,5,7,9,11}
复制
返回值：
{8,6,10,5,7,9,11}
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param root TreeNode类
 * @return TreeNode类
 */
// func Serialize( root *TreeNode ) string {
//     // write code here
// }
// func Deserialize( s string ) *TreeNode {
//     // write code here
// }

/**
NC81 二叉搜索树的第k个结点
 算法知识视频讲解
简单  通过率：25.18%  时间限制：1秒  空间限制：64M
知识点
树
题目
题解(98)
讨论(905)
排行
描述
给定一棵节点数为 n 二叉搜索树，请找出其中的第 k 小的TreeNode结点。

数据范围： ，，树上每个节点的值满足
要求：空间复杂度 ，时间复杂度
注意：不是返回结点的值
输入描述：
提示：当n为0或者k为0时返回空。
示例1
输入：
{5,3,7,2,4,6,8},3
复制
返回值：
4
复制
说明：
按结点数值大小顺序第三小结点的值为4
示例2
输入：
{},1
复制
返回值：
"null"
*/
/*
 * type TreeNode struct {
 *   Val int
 *   Left *TreeNode
 *   Right *TreeNode
 * }
 */

/**
 *
 * @param pRoot TreeNode类
 * @param k int整型
 * @return TreeNode类
 */
// func KthNode( pRoot *TreeNode ,  k int ) *TreeNode {
//     // write code here
// }

/**
NC89 字符串变形
 算法知识视频讲解
简单  通过率：24.20%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(16)
讨论(86)
排行
描述
对于一个长度为 n 字符串，我们需要对它做一些变形。

首先这个字符串中包含着一些空格，就像"Hello World"一样，然后我们要做的是把这个字符串中由空格隔开的单词反序，同时反转每个字符的大小写。

比如"Hello World"变形后就变成了"wORLD hELLO"。

数据范围:  , 字符串中包括大写英文字母、小写英文字母、空格。
进阶：空间复杂度  ， 时间复杂度
输入描述：
给定一个字符串s以及它的长度n(1 ≤ n ≤ 10^6)
返回值描述：
请返回变形后的字符串。题目保证给定的字符串均由大小写字母和空格构成。
示例1
输入：
"This is a sample",16
复制
返回值：
"SAMPLE A IS tHIS"
复制
示例2
输入：
"nowcoder“
复制
返回值：
”NOWCODER“
复制
示例3
输入：
"iOS"
复制
返回值：
"Ios"
*/
/**
 *
 * @param s string字符串
 * @param n int整型
 * @return string字符串
 */
// func trans( s string ,  n int ) string {
//     // write code here
// }

/**
NC95 数组中的最长连续子序列
 算法知识视频讲解
较难  通过率：40.27%  时间限制：2秒  空间限制：256M
知识点
并查集
数组
题目
题解(32)
讨论(72)
排行
描述
给定无序数组arr，返回其中最长的连续序列的长度(要求值连续，位置可以不连续,例如 3,4,5,6为连续的自然数）

数据范围： ，数组中的值满足
要求：空间复杂度 ，时间复杂度
示例1
输入：
[100,4,200,1,3,2]
复制
返回值：
4
复制
示例2
输入：
[1,1,1]
复制
返回值：
1
复制
备注：
1 \leq n \leq 10^51≤n≤10
5

1 \leq arr_i \leq 10^81≤arr
i
​
 ≤10
8
*/
/**
 * max increasing subsequence
 * @param arr int整型一维数组 the array
 * @return int整型
 */
// func MLS( arr []int ) int {
//     // write code here
// }

/**
NC132 环形链表的约瑟夫问题
 算法知识视频讲解
中等  通过率：52.36%  时间限制：2秒  空间限制：256M
知识点
链表
数学
题目
题解(21)
讨论(42)
排行
描述
编号为 1 到 n 的 n 个人围成一圈。从编号为 1 的人开始报数，报到 m 的人离开。
下一个人继续从 1 开始报数。
n-1 轮结束以后，只剩下一个人，问最后留下的这个人编号是多少？

数据范围：
进阶：空间复杂度 ，时间复杂度
示例1
输入：
5,2
复制
返回值：
3
复制
说明：
开始5个人 1，2，3，4，5 ，从1开始报数，1->1，2->2编号为2的人离开
1，3，4，5，从3开始报数，3->1，4->2编号为4的人离开
1，3，5，从5开始报数，5->1，1->2编号为1的人离开
3，5，从3开始报数，3->1，5->2编号为5的人离开
最后留下人的编号是3
示例2
输入：
1,1
复制
返回值：
1
复制
备注：
1 \leq n, m \leq 100001≤n,m≤10000
*/
/**
 *
 * @param n int整型
 * @param m int整型
 * @return int整型
 */
// func ysf( n int ,  m int ) int {
//     // write code here
// }

/**
NC99 树的直径
 算法知识视频讲解
较难  通过率：36.47%  时间限制：1秒  空间限制：128M
知识点
树
题目
题解(10)
讨论(29)
排行
描述
给定一棵树，求出这棵树的直径，即树上最远两点的距离。
包含n个结点，n-1条边的连通图称为树。
示例1的树如下图所示。其中4到5之间的路径最长，是树的直径，距离为5+2+4=11

数据范围：，保证最终结果满足
要求：空间复杂度：O(N)，时间复杂度O(N)
*/
/*
 * type Interval struct {
 *   Start int
 *   End int
 * }
 */

/**
 * 树的直径
 * @param n int整型 树的节点个数
 * @param Tree_edge Interval类一维数组 树的边
 * @param Edge_value int整型一维数组 边的权值
 * @return int整型
 */
// func solve( n int ,  Tree_edge []*Interval ,  Edge_value []int ) int {
//     // write code here
// }

/**
NC111 最大数
 算法知识视频讲解
中等  通过率：38.18%  时间限制：1秒  空间限制：64M
知识点
排序
题目
题解(32)
讨论(56)
排行
描述
给定一个nums数组由一些非负整数组成，现需要将他们进行排列并拼接，每个数不可拆分，使得最后的结果最大，返回值需要是string类型，否则可能会溢出

提示:
1 <= nums.length <= 100
0 <= nums[i] <= 10000

示例1
输入：
[30,1]
复制
返回值：
"301"
复制
示例2
输入：
[2,20,23,4,8]
复制
返回值：
"8423220"
复制
示例3
输入：
[2]
复制
返回值：
"2"
复制
示例4
输入：
[10]
复制
返回值：
"10"
*/
/**
 * 最大数
 * @param nums int整型一维数组
 * @return string字符串
 */
// func solve( nums []int ) string {
//     // write code here
// }

/**
NC113 验证IP地址
 算法知识视频讲解
中等  通过率：39.79%  时间限制：1秒  空间限制：64M
知识点
字符串
题目
题解(15)
讨论(50)
排行
描述
编写一个函数来验证输入的字符串是否是有效的 IPv4 或 IPv6 地址

IPv4 地址由十进制数和点来表示，每个地址包含4个十进制数，其范围为 0 - 255， 用(".")分割。比如，172.16.254.1；
同时，IPv4 地址内的数不会以 0 开头。比如，地址 172.16.254.01 是不合法的。

IPv6 地址由8组16进制的数字来表示，每组表示 16 比特。这些组数字通过 (":")分割。比如,  2001:0db8:85a3:0000:0000:8a2e:0370:7334 是一个有效的地址。而且，我们可以加入一些以 0 开头的数字，字母可以使用大写，也可以是小写。所以， 2001:db8:85a3:0:0:8A2E:0370:7334 也是一个有效的 IPv6 address地址 (即，忽略 0 开头，忽略大小写)。

然而，我们不能因为某个组的值为 0，而使用一个空的组，以至于出现 (::) 的情况。 比如， 2001:0db8:85a3::8A2E:0370:7334 是无效的 IPv6 地址。
同时，在 IPv6 地址中，多余的 0 也是不被允许的。比如， 02001:0db8:85a3:0000:0000:8a2e:0370:7334 是无效的。

说明: 你可以认为给定的字符串里没有空格或者其他特殊字符。

数据范围：字符串长度满足
进阶：空间复杂度 ，时间复杂度
示例1
输入：
"172.16.254.1"
复制
返回值：
"IPv4"
复制
说明：
这是一个有效的 IPv4 地址, 所以返回 "IPv4"
示例2
输入：
"2001:0db8:85a3:0:0:8A2E:0370:7334"
复制
返回值：
"IPv6"
复制
说明：
这是一个有效的 IPv6 地址, 所以返回 "IPv6"
示例3
输入：
"256.256.256.256"
复制
返回值：
"Neither"
复制
说明：
这个地址既不是 IPv4 也不是 IPv6 地址
备注：
ip地址的类型，可能为
IPv4,   IPv6,   Neither
*/
/**
 * 验证IP地址
 * @param IP string字符串 一个IP地址字符串
 * @return string字符串
 */
// func solve( IP string ) string {
//     // write code here
// }

/**
NC10 大数乘法
 算法知识视频讲解
中等  通过率：43.72%  时间限制：1秒  空间限制：256M
知识点
字符串
题目
题解(17)
讨论(38)
排行
描述
以字符串的形式读入两个数字，编写一个函数计算它们的乘积，以字符串形式返回。

数据范围： 读入的数字大小满足
要求：空间复杂度 ，时间复杂度
示例1
输入：
"11","99"
复制
返回值：
"1089"
复制
说明：
11*99=1089
示例2
输入：
"1","0"
复制
返回值：
"0"
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param s string字符串 第一个整数
 * @param t string字符串 第二个整数
 * @return string字符串
 */
// func solve( s string ,  t string ) string {
//     // write code here
// }

/**
NC27 集合的所有子集
 算法知识视频讲解
中等  通过率：30.91%  时间限制：1秒  空间限制：64M
知识点
位运算
数组
回溯
题目
题解(24)
讨论(117)
排行
描述
现在有一个没有重复元素的整数集合S，求S的所有子集
注意：
你给出的子集中的元素必须按升序排列
给出的解集中不能出现重复的元素

数据范围：，集合中的任意元素满足
要求：空间复杂度 ，时间复杂度
示例1
输入：
[1,2,3]
复制
返回值：
[[],[1],[2],[3],[1,2],[1,3],[2,3],[1,2,3]]
复制
示例2
输入：
[]
复制
返回值：
[]
*/
/**
 *
 * @param A int整型一维数组
 * @return int整型二维数组
 */
// func subsets( A []int ) [][]int {
//     // write code here
// }

/**
NC43 没有重复项数字的所有排列
 算法知识视频讲解
中等  通过率：45.10%  时间限制：1秒  空间限制：64M
知识点
回溯
题目
题解(9)
讨论(85)
排行
描述
给出一组数字，返回该组数字的所有排列
例如：
[1,2,3]的所有排列如下
[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2], [3,2,1].
（以数字在数组中的位置靠前为优先级，按字典序排列输出。）

示例1
输入：
[1,2,3]
复制
返回值：
[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
*/
/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
// func permute( num []int ) [][]int {
//     // write code here
// }

/**
NC69 链表中倒数最后k个结点
 算法知识视频讲解
中等  通过率：36.38%  时间限制：1秒  空间限制：256M
知识点
链表
题目
题解(77)
讨论(129)
排行
描述
输入一个长度为的链表，设链表中的元素的值为a_ia
i
​
 ，输出一个链表，该输出链表包含原链表中从倒数第k个结点至尾节点的全部节点。
如果该链表长度小于k，请返回一个长度为 0 的链表。


提示:
0 \leq n \leq 10^50≤n≤10
5

0 \leq a_i \leq 10^90≤a
i
​
 ≤10
9

0 \leq k \leq 10^90≤k≤10
9

示例1
输入：
{1,2,3,4,5},3
复制
返回值：
{3,4,5}
复制
示例2
输入：
{2},8
复制
返回值：
{}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param pHead ListNode类
 * @param k int整型
 * @return ListNode类
 */
// func FindKthToTail( pHead *ListNode ,  k int ) *ListNode {
//     // write code here
// }

/**
NC126 换钱的最少货币数
 算法知识视频讲解
简单  通过率：43.80%  时间限制：2秒  空间限制：256M
知识点
动态规划
题目
题解(29)
讨论(37)
排行
描述
给定数组arr，arr中所有的值都为正整数且不重复。每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，求组成aim的最少货币数。
如果无解，请返回-1.
【要求】
时间复杂度O(n \times aim)O(n×aim)，空间复杂度On。
示例1
输入：
[5,2,3],20
复制
返回值：
4
复制
示例2
输入：
[5,2,3],0
复制
返回值：
0
复制
示例3
输入：
[3,5],2
复制
返回值：
-1
复制
备注：
0 \leq n \leq 1\,0000≤n≤1000
0 \leq aim \leq 5\,0000≤aim≤5000
*/
/**
 * 最少货币数
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
 */
// func minMoney( arr []int ,  aim int ) int {
//     // write code here
// }

/**
NC107 寻找峰值
 算法知识视频讲解
入门  通过率：33.92%  时间限制：1秒  空间限制：256M
知识点
数组
题目
题解(29)
讨论(75)
排行
描述
山峰元素是指其值大于或等于左右相邻值的元素。给定一个输入数组nums，任意两个相邻元素值不相等，数组可能包含多个山峰。找到索引最大的那个山峰元素并返回其索引。

假设 nums[-1] = nums[n] = -∞。


提示:
1 <= 数组长度 <= 1000
0 <= 数组元素的值 <= 1000
示例1
输入：
[2,4,1,2,7,8,4]
复制
返回值：
5
*/
/**
 * 寻找最后的山峰
 * @param a int整型一维数组
 * @return int整型
 */
// func solve( a []int ) int {
//     // write code here
// }

/**
NC28 最小覆盖子串
 算法知识视频讲解
较难  通过率：32.23%  时间限制：1秒  空间限制：64M
知识点
哈希
双指针
字符串
题目
题解(19)
讨论(71)
排行
描述
给出两个字符串 SS 和 TT，要求在O(n)O(n)的时间复杂度内在 SS 中找出最短的包含 TT 中所有字符的子串。
例如：

S ="XDOYEZODEYXNZ"S="XDOYEZODEYXNZ"
T ="XYZ"T="XYZ"
找出的最短子串为"YXNZ""YXNZ".

注意：
如果 SS 中没有包含 TT 中所有字符的子串，返回空字符串 “”；
满足条件的子串可能有很多，但是题目保证满足条件的最短的子串唯一。

示例1
输入：
"XDOYEZODEYXNZ","XYZ"
复制
返回值：
"YXNZ"
*/
/**
 *
 * @param S string字符串
 * @param T string字符串
 * @return string字符串
 */
// func minWindow( S string ,  T string ) string {
//     // write code here
// }

/**
NC29 二维数组中的查找
 算法知识视频讲解
中等  通过率：25.92%  时间限制：1秒  空间限制：64M
知识点
数组
题目
题解(202)
讨论(3k)
排行
描述
在一个二维数组array中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
[
[1,2,8,9],
[2,4,9,12],
[4,7,10,13],
[6,8,11,15]
]
给定 target = 7，返回 true。

给定 target = 3，返回 false。

0 <= array.length <= 500
0 <= array[0].length <= 500

你能给出时间复杂度为  的解法吗？（n,m为矩阵的长和宽）
示例1
输入：
7,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
复制
返回值：
true
复制
说明：
存在7，返回true
示例2
输入：
3,[[1,2,8,9],[2,4,9,12],[4,7,10,13],[6,8,11,15]]
复制
返回值：
false
复制
说明：
不存在3，返回false
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
// func Find( target int ,  array [][]int ) bool {
//     // write code here
// }

/**
NC131 数据流中的中位数
 算法知识视频讲解
中等  通过率：27.23%  时间限制：1秒  空间限制：64M
知识点
排序
堆
题目
题解(59)
讨论(663)
排行
描述
如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。我们使用Insert()方法读取数据流，使用GetMedian()方法获取当前读取数据的中位数。
示例1
输入：
[5,2,3,4,1,6,7,0,8]
复制
返回值：
"5.00 3.50 3.00 3.50 3.00 3.50 4.00 3.50 4.00 "
复制
说明：
数据流里面不断吐出的是5,2,3...,则得到的平均数分别为5,(5+2)/2,3...
*/

type MinHeapInt []int

func (h MinHeapInt) Len() int {
	return len(h)
}

func (h MinHeapInt) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] < h[j]
}

func (h *MinHeapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MinHeapInt) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *MinHeapInt) Top() interface{} {
	return (*h)[0]
}

type MaxHeapInt []int

func (h MaxHeapInt) Len() int {
	return len(h)
}

func (h MaxHeapInt) Less(i, j int) bool {
	// 由于是最大堆，所以使用大于号
	return h[i] > h[j]
}

func (h *MaxHeapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeapInt) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出最后一个元素
func (h *MaxHeapInt) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h *MaxHeapInt) Top() interface{} {
	return (*h)[0]
}

var minHeap = make(MinHeapInt, 0)
var maxHeap = make(MaxHeapInt, 0)

func Insert(num int) {
	if maxHeap.Len() == 0 {
		heap.Init(&maxHeap)
		heap.Init(&minHeap)
		heap.Push(&maxHeap, num)
		return
	}

	left := maxHeap.Top().(int)
	if num < left { // 小进大堆
		heap.Push(&maxHeap, num)
	} else { // 大进小堆
		heap.Push(&minHeap, num)
	}

	for maxHeap.Len()-minHeap.Len() > 1 {
		heap.Push(&minHeap, heap.Pop(&maxHeap))
	}

	for minHeap.Len() > maxHeap.Len() {
		heap.Push(&maxHeap, heap.Pop(&minHeap))
	}
}

func GetMedian() float64 {
	if maxHeap.Len() == 0 {
		return 0
	}

	if maxHeap.Len() == minHeap.Len() {
		return (float64(maxHeap.Top().(int)) + float64(minHeap.Top().(int))) / 2
	} else {
		return float64(maxHeap.Top().(int))
	}
}

/**
NC145 01背包
 算法知识视频讲解
简单  通过率：46.82%  时间限制：1秒  空间限制：256M
知识点
动态规划
题目
题解(15)
讨论(31)
排行
描述
已知一个背包最多能容纳物体的体积为

现有 n 个物品，第 i 个物品的体积为 v_i, 重量为 w_i
​
求当前背包最多能装多大重量的物品?

数据范围：
1≤V≤5000
1≤n≤5000
1≤v_i≤5000
1≤w_i≤5000

复杂度要求：
O(n⋅V)
示例1
输入：
10,2,[[1,3],[10,4]]
复制
返回值：
4
复制
说明：
第一个物品的体积为1，重量为3，第二个物品的体积为10，重量为4。只取第二个物品可以达到最优方案，取物重量为4
备注：
1≤V≤5000
1≤n≤5000
1≤v_i≤5000
1≤w_i≤5000

O(n⋅V)
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算01背包问题的结果
 * @param V int整型 背包的体积
 * @param n int整型 物品的个数
 * @param vw int整型二维数组 第一维度为n,第二维度为2的二维数组,vw[i][0],vw[i][1]分别描述i+1个物品的vi,wi
 * @return int整型
 */
func knapsack(V int, n int, vw [][]int) int {
	// write code here
	return 0
}
