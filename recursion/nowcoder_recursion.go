package recursion

/**

NC65 斐波那契数列

描述
大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。n≤39

示例1
输入：4
返回值：3
*/

var mapInt map[int]int

/**
 *
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if mapInt == nil {
		mapInt = make(map[int]int)
	}
	if value, ok := mapInt[n]; ok {
		return value
	}
	res := Fibonacci(n-1) + Fibonacci(n-2)
	mapInt[n] = res
	return res
}

/**

NC68 跳台阶

描述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
示例1
输入：
2
复制
返回值：
2
复制
示例2
输入：
7
复制
返回值：
21
复制
*/

var mapInt1 map[int]int

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	if number <= 2 {
		return number
	}
	if mapInt1 == nil {
		mapInt1 = make(map[int]int)
	}
	if value, ok := mapInt1[number]; ok {
		return value
	}
	res := jumpFloor(number-1) + jumpFloor(number-2)
	mapInt1[number] = res
	return res
}

/**
链表反转递归实现
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(pHead *ListNode) *ListNode {
	var pre *ListNode
	for cur := pHead; cur != nil; {
		// pre, cur, cur.Next = cur, cur.Next, pre //这行代码效果等同于下面
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

/**
斐波那契数列、
全排列、
⼋皇后、
快速排序；
归并排序、
DFS、
⼆叉树遍历、
链表反转递归实现等
*/

/**

全排列

给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

示例 1：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

示例 2：
输入：nums = [0,1]
输出：[[0,1],[1,0]]

示例 3：
输入：nums = [1]
输出：[[1]]

提示：

1 <= nums.length <= 6
-10 <= nums[i] <= 10
nums 中的所有整数 互不相同
*/
func permute(nums []int) [][]int {

}

/**

N皇后

设计一种算法，打印 N 皇后在 N × N 棋盘上的各种摆法，其中每个皇后都不同行、不同列，也不在对角线上。这里的“对角线”指的是所有的对角线，不只是平分整个棋盘的那两条对角线。

注意：本题相对原题做了扩展

示例:

 输入：4
 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
 解释: 4 皇后问题存在如下两个不同的解法。
[
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]

*/

func solveNQueens(n int) [][]string {

}
