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
斐波那契数列、
全排列、
⼋皇后、
快速排序；
归并排序、
DFS、
⼆叉树遍历、
链表反转递归实现等
*/
