package nowcoder

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strconv"
)

/**
NC78 反转链表
 算法知识视频讲解
简单  通过率：35.41%  时间限制：1秒  空间限制：64M
知识点
链表
iOS工程师
小米
2021
题目
题解(282)
讨论(2k)
排行
描述
输入一个链表，反转链表后，输出新链表的表头。
示例1
输入：
{1,2,3}
复制
返回值：
{3,2,1}
复制
*/

/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	var pre *ListNode
	for cur := pHead; cur != nil; {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

/**
NC93 设计LRU缓存结构
 算法知识视频讲解
中等  通过率：34.28%  时间限制：2秒  空间限制：256M
知识点
模拟
题目
题解(154)
讨论(321)
排行
描述
设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
1. set(key, value)：将记录(key, value)插入该结构
2. get(key)：返回key对应的value值

提示:
1.某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的，然后都会刷新缓存。
2.当缓存的大小超过K时，移除最不经常使用的记录。
3.输入一个二维数组与K，二维数组每一维有2个或者3个数字，第1个数字为opt，第2，3个数字为key，value
   若opt=1，接下来两个整数key, value，表示set(key, value)
   若opt=2，接下来一个整数key，表示get(key)，若key未出现过或已被移除，则返回-1
   对于每个opt=2，输出一个答案
4.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹

进阶:你是否可以在O(1)的时间复杂度完成set和get操作
示例1
输入：
[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3
复制
返回值：
[1,-1]
复制
说明：
[1,1,1]，第一个1表示opt=1，要set(1,1)，即将(1,1)插入缓存，缓存是{"1"=1}
[1,2,2]，第一个1表示opt=1，要set(2,2)，即将(2,2)插入缓存，缓存是{"1"=1,"2"=2}
[1,3,2]，第一个1表示opt=1，要set(3,2)，即将(3,2)插入缓存，缓存是{"1"=1,"2"=2,"3"=2}
[2,1]，第一个2表示opt=2，要get(1)，返回是[1]，因为get(1)操作，缓存更新，缓存是{"2"=2,"3"=2,"1"=1}
[1,4,4]，第一个1表示opt=1，要set(4,4)，即将(4,4)插入缓存，但是缓存已经达到最大容量3，移除最不经常使用的{"2"=2}，插入{"4"=4}，缓存是{"3"=2,"1"=1,"4"=4}
[2,2]，第一个2表示opt=2，要get(2)，查找不到，返回是[1,-1]
示例2
输入：
[[1,1,1],[1,2,2],[2,1],[1,3,3],[2,2],[1,4,4],[2,1],[2,3],[2,4]],2
复制
返回值：
[1,-1,-1,3,4]
复制
备注：
1≤K≤N≤10^5
−2×10^9≤x,y≤2×10^9
*/
/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	ret := []int{}
	cache := NewLRUCache(k)
	for _, ele := range operators {
		switch ele[0] {
		case 1:
			cache.putCache(ele[1], ele[2])
		case 2:
			ret = append(ret, cache.getCache(ele[1]))
		}
	}
	return ret
}

type LRUCache struct {
	cap   int
	cache map[int]*list.Element
	ll    *list.List
}
type entry struct {
	key   int
	value int
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		cap:   capacity,
		cache: make(map[int]*list.Element),
		ll:    list.New(),
	}
}

func (c *LRUCache) getCache(key int) int {
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		return ee.Value.(*entry).value
	}
	return -1
}

func (c *LRUCache) putCache(key int, value int) {
	if ee, ok := c.cache[key]; ok { // 已存在的修改
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		c.cache[key] = ee
		return
	}

	if c.cap == c.ll.Len() { // 超出上限，则从末尾删除
		ee := c.ll.Back()
		c.ll.Remove(ee)
		delete(c.cache, ee.Value.(*entry).key)
	}

	ee := c.ll.PushFront(&entry{key, value}) // 不存在的新增
	c.cache[key] = ee
}

/**
NC45 实现二叉树先序，中序和后序遍历
 算法知识视频讲解
中等  通过率：42.74%  时间限制：3秒  空间限制：256M
知识点
栈
树
哈希
题目
题解(95)
讨论(222)
排行
描述
分别按照二叉树先序，中序和后序打印所有的节点。
示例1
输入：
{1,2,3}
复制
返回值：
[[1,2,3],[2,1,3],[2,3,1]]
复制
备注：
n≤10^6
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 *
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	ret := [][]int{}
	ret = append(ret, preOrder(root), order(root), postOrder(root))
	return ret
}

func preOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	ret = append(ret, root.Val)
	ret = append(ret, preOrder(root.Left)...)
	ret = append(ret, preOrder(root.Right)...)
	return ret
}

func order(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	ret = append(ret, order(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, order(root.Right)...)
	return ret
}

func postOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}
	ret = append(ret, postOrder(root.Left)...)
	ret = append(ret, postOrder(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}

/**
NC119 最小的K个数
 算法知识视频讲解
中等  通过率：25.62%  时间限制：1秒  空间限制：64M
知识点
堆
排序
分治
题目
题解(191)
讨论(1k)
排行
描述
给定一个数组，找出其中最小的K个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4。
0 <= k <= input.length <= 10000
0 <= input[i] <= 10000

示例1
输入：
[4,5,1,6,2,7,3,8],4
复制
返回值：
[1,2,3,4]
复制
说明：
返回最小的4个数即可，返回[1,3,2,4]也可以
示例2
输入：
[1],0
复制
返回值：
[]
复制
示例3
输入：
[0,1,2,1,2],3
复制
返回值：
[0,1,1]
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param input int整型一维数组
 * @param k int整型
 * @return int整型一维数组
 */
// 1. 内置排序，复杂度为O(NlogN)
func GetLeastNumbers_SelfSort(input []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if k >= len(input) {
		return input
	}

	sort.Ints(input)
	return input[:k]
}

// 2. 大根堆（前K小），小根堆（前K大），复杂度O(NlogK)
func GetLeastNumbers_Heap(input []int, k int) []int {
	if k == 0 || len(input) == 0 {
		return []int{}
	}

	heapSort(input)
	return input[:k]
}

// 堆排序
func heapSort(input []int) {
	if len(input) < 2 {
		return
	}
	for i := 0; i < len(input)-1; i++ {
		minAdjust(input[i:]) // 每次确认第i位置的数
	}
}

func minAdjust(input []int) {
	len := len(input)
	if len < 2 {
		return
	}

	for i := (len - 1) / 2; i >= 0; i-- {
		if 2*i+1 <= len-1 && input[i] > input[2*i+1] {
			input[i], input[2*i+1] = input[2*i+1], input[i]
		}
		if 2*i+2 <= len-1 && input[i] > input[2*i+2] {
			input[i], input[2*i+2] = input[2*i+2], input[i]
		}
	}
}

// 3. 快排仅排好第K小的数，那么它左边的数就是比它小的另外K-1个数（不对整个数组排序），时间复杂度为N+N/2+N/4+...N/N = 2N,O(N)
func GetLeastNumbers_QuickSearch(input []int, k int) []int {
	if k == 0 || len(input) == 0 {
		return []int{}
	}
	return QuickSearch(input, 0, len(input)-1, k-1)
}

func QuickSearch(input []int, low, high, k int) []int {
	var j int
	input, j = partition(input, low, high)
	if j == k {
		return input[:k+1]
	}
	if j > k {
		input = QuickSearch(input, low, j-1, k)
	} else {
		input = QuickSearch(input, j+1, high, k)
	}
	return input
}

// 4. 自己实现的快速排序，复杂度为O(NlogN)
func GetLeastNumbers_QuickSort(input []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if k >= len(input) {
		return input
	}

	input = QuickSort(input, 0, len(input)-1)
	return input[:k]
}

func QuickSort(input []int, low, high int) []int {
	if low < high {
		var j int
		input, j = partition(input, low, high)
		input = QuickSort(input, low, j-1)
		input = QuickSort(input, j+1, high)
	}
	return input
}

func partition(input []int, low, high int) ([]int, int) {
	pivot := input[high]
	i := low
	for j := low; j < high; j++ {
		if input[j] < pivot {
			if i != j {
				input[i], input[j] = input[j], input[i]
			}
			i++
		}
	}
	if i != high {
		input[i], input[high] = input[high], input[i]
	}
	return input, i
}

// 5. 计数排序，复杂度为O(N)
func GetLeastNumbers_CounterSort(input []int, k int) []int {
	counter := make([][]int, 10001)
	len1 := len(input)
	for i := 0; i < len1; i++ {
		counter[input[i]] = append(counter[input[i]], i)
	}

	res := []int{}
	for i := 0; i < 10001; i++ {
		for j := 0; j < len(counter[i]) && len(res) < k; j++ {
			res = append(res, input[counter[i][j]])
		}
		if len(res) == k {
			break
		}
	}

	return res
}

/**
NC15 求二叉树的层序遍历
 算法知识视频讲解
中等  通过率：43.77%  时间限制：1秒  空间限制：64M
知识点
树
bfs
题目
题解(97)
讨论(249)
排行
描述
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
例如：
给定的二叉树是{3,9,20,#,#,15,7},

该二叉树层序遍历的结果是
[
[3],
[9,20],
[15,7]
]

示例1
输入：
{1,2}
复制
返回值：
[[1],[2]]
复制
示例2
输入：
{1,2,3,4,#,#,5}
复制
返回值：
[[1],[2,3],[4,5]]
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
 * @return int整型二维数组
 */
func levelOrder(root *TreeNode) [][]int {
	vv := [][]int{}

	if root == nil {
		return vv
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	for len(queue) > 0 {
		tmpint := []int{}
		len := len(queue)
		for i := 0; i < len; i++ {
			node := queue[i]
			tmpint = append(tmpint, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[len:]
		vv = append(vv, tmpint)
	}
	return vv
}

/**
NC88 寻找第K大
 算法知识视频讲解
中等  通过率：28.25%  时间限制：3秒  空间限制：64M
知识点
堆
分治
题目
题解(98)
讨论(287)
排行
描述
有一个整数数组，请你根据快速排序的思路，找出数组中第 大的数。

给定一个整数数组 ,同时给定它的大小n和要找的 ，请返回第 大的数(包括重复的元素，不用去重)，保证答案存在。
要求时间复杂度O(N)
示例1
输入：
[1,3,5,2,2],5,3
复制
返回值：
2
复制
示例2
输入：
[10,10,9,9,8,7,5,6,4,3,4,2],12,3
复制
返回值：
9
复制
说明：
去重后的第3大是8，但本题要求包含重复的元素，不用去重，所以输出9
*/
/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	// 第K大，就是第n-K+1小（转换下思路）
	// write code here
	a = findKth_QuickSearch(a, 0, n-1, n-K)
	return a[n-K] // 第n-K+1小
}

// 快速排序变种，复杂度为O(N)
func findKth_QuickSearch(a []int, low, high, k int) []int {
	a, j := partition(a, low, high)
	if j == k {
		return a
	}

	if j > k {
		return findKth_QuickSearch(a, low, j-1, k)
	} else {
		return findKth_QuickSearch(a, j+1, high, k)
	}
}

// 堆排序，复杂度为O(N)
func findKth_HeapSort(a []int, n int, K int) int {
	// 第K大
	for i := 0; i < K; i++ { // 每次定第i+1大的值
		maxAdjust(a[i:])
	}
	return a[K-1]
}

func maxAdjust(a []int) {
	length := len(a)
	if length < 2 {
		return
	}

	for i := (length - 1) / 2; i >= 0; i-- {
		if 2*i+1 < length && a[2*i+1] > a[i] {
			a[2*i+1], a[i] = a[i], a[2*i+1]
		}

		if 2*i+2 < length && a[2*i+2] > a[i] {
			a[2*i+2], a[i] = a[i], a[2*i+2]
		}
	}
}

/**
NC61 两数之和
 算法知识视频讲解
简单  通过率：36.26%  时间限制：2秒  空间限制：64M
知识点
数组
哈希
题目
题解(81)
讨论(358)
排行
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
复制
返回值：
[2,3]
复制
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
	// write code here
	maps := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		if k, ok := maps[target-numbers[i]]; ok {
			return []int{k + 1, i + 1}
		} else {
			maps[numbers[i]] = i
		}
	}
	return []int{}
}

/**
NC33 合并两个排序的链表
 算法知识视频讲解
简单  通过率：29.20%  时间限制：1秒  空间限制：64M
知识点
链表
测试开发工程师
阅文集团
2021
题目
题解(98)
讨论(1k)
排行
描述
输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
示例1
输入：
{1,3,5},{2,4,6}
复制
返回值：
{1,2,3,4,5,6}
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}

	if pHead2 == nil {
		return pHead1
	}

	var pHead *ListNode
	var pre, cur *ListNode
	for cur1, cur2 := pHead1, pHead2; cur1 != nil && cur2 != nil; {
		if cur1.Val <= cur2.Val {
			cur = cur1
			cur1 = cur1.Next
		} else {
			cur = cur2
			cur2 = cur2.Next
		}
		if pHead == nil {
			pHead = cur
		} else {
			pre.Next = cur
		}
		pre = cur
		if cur1 == nil { // cur2还有直接接上
			pre.Next = cur2
			break
		}

		if cur2 == nil {
			pre.Next = cur1
			break
		}
	}

	return pHead
}

/**
NC76 用两个栈实现队列
 算法知识视频讲解
简单  通过率：39.02%  时间限制：1秒  空间限制：64M
知识点
栈
题目
题解(172)
讨论(2k)
排行
描述
用两个栈来实现一个队列，分别完成在队列尾部插入整数(push)和在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

示例:
输入:
["PSH1","PSH2","POP","POP"]
返回:
1,2
解析:
"PSH1":代表将1插入队列尾部
"PSH2":代表将2插入队列尾部
"POP“:代表删除一个元素，先进先出=>返回1
"POP“:代表删除一个元素，先进先出=>返回2
示例1
输入：
["PSH1","PSH2","POP","POP"]
复制
返回值：
1,2
*/
var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 && len(stack1) > 0 {
		for i := len(stack1) - 1; i >= 0; i-- { // 尾部先出
			stack2 = append(stack2, stack1[i])
		}
		stack1 = stack1[len(stack1):]
	}

	if len(stack2) == 0 {
		return -1
	}

	item := stack2[len(stack2)-1] // 从尾部pop
	stack2 = stack2[:len(stack2)-1]
	return item
}

/**
NC68 跳台阶
 算法知识视频讲解
简单  通过率：38.92%  时间限制：1秒  空间限制：64M
知识点
递归
题目
题解(148)
讨论(2k)
排行
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
*/

var jumpCache map[int]int

/**
 *
 * @param number int整型
 * @return int整型
 */
func jumpFloor(number int) int {
	// write code here
	if number <= 2 {
		return number
	}

	if jumpCache == nil { // 初始化
		jumpCache = make(map[int]int)
	}

	if v, ok := jumpCache[number]; ok {
		return v
	}

	cache := jumpFloor(number-1) + jumpFloor(number-2)
	jumpCache[number] = cache
	return cache
}

/**
NC50 链表中的节点每k个一组翻转
 算法知识视频讲解
中等  通过率：36.20%  时间限制：1秒  空间限制：64M
知识点
链表
题目
题解(77)
讨论(235)
排行
描述
将给出的链表中的节点每\ k k 个一组翻转，返回翻转后的链表
如果链表中的节点数不是\ k k 的倍数，将最后剩下的节点保持原样
你不能更改节点中的值，只能更改节点本身。
要求空间复杂度 \ O(1) O(1)
例如：
给定的链表是1\to2\to3\to4\to51→2→3→4→5
对于 \ k = 2 k=2, 你应该返回 2\to 1\to 4\to 3\to 52→1→4→3→5
对于 \ k = 3 k=3, 你应该返回 3\to2 \to1 \to 4\to 53→2→1→4→5

示例1
输入：
{1,2,3,4,5},2
复制
返回值：
{2,1,4,3,5}
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
 * @param k int整型
 * @return ListNode类
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	// write code here
	nextHead := head
	// 寻找可能存在的下一个head结点
	for i := 0; i < k; i++ {
		if nextHead == nil {
			return head
		}
		nextHead = nextHead.Next
	}

	// 下一个head结点递归获取反转结果
	res := reverse(head, nextHead)
	head.Next = reverseKGroup(nextHead, k)
	return res
}

// 反转left到right间的结点，左闭右开
func reverse(left, right *ListNode) *ListNode {
	pre := right
	for left != right {
		tmp := left.Next
		left.Next = pre
		pre = left
		left = tmp
	}
	return pre
}

/**
NC19 子数组的最大累加和问题
 算法知识视频讲解
简单  通过率：59.20%  时间限制：2秒  空间限制：256M
知识点
分治
动态规划
题目
题解(105)
讨论(196)
排行
描述
给定一个数组arr，返回子数组的最大累加和
例如，arr = [1, -2, 3, 5, -2, 6, -1]，所有子数组中，[3, 5, -2, 6]可以累加出最大的和12，所以返回12.
题目保证没有全为负数的数据
[要求]
时间复杂度为O(n)O(n)，空间复杂度为O(1)O(1)

示例1
输入：
[1, -2, 3, 5, -2, 6, -1]
复制
返回值：
12
复制
备注：
1≤N≤10^5
|arr_i|≤100
*/

/**
 * max sum of the subarray
 * @param arr int整型一维数组 the array
 * @return int整型
 */
func maxsumofSubarray(arr []int) int {
	// write code here
	length := len(arr)
	if length == 0 {
		return 0
	}
	dp := make([]int, length)
	dp[0] = arr[0]
	for i := 1; i < length; i++ {
		if arr[i] < dp[i-1]+arr[i] {
			dp[i] = dp[i-1] + arr[i]
		} else {
			dp[i] = arr[i]
		}
	}

	// 遍历获取最大累加和
	res := math.MinInt64
	for i := 0; i < length; i++ {
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

/**
NC41 最长无重复子数组
 算法知识视频讲解
中等  通过率：31.80%  时间限制：2秒  空间限制：256M
知识点
哈希
双指针
数组
题目
题解(126)
讨论(293)
排行
描述
给定一个数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组
示例1
输入：
[2,3,4,5]
复制
返回值：
4
复制
说明：
[2,3,4,5]是最长子数组
示例2
输入：
[2,2,3,4,3]
复制
返回值：
3
复制
说明：
[2,3,4]是最长子数组
示例3
输入：
[9]
复制
返回值：
1
复制
示例4
输入：
[1,2,3,1,2,3,2,2]
复制
返回值：
3
复制
说明：
最长子数组为[1,2,3]
示例5
输入：
[2,2,3,4,8,99,3]
复制
返回值：
5
复制
说明：
最长子数组为[2,3,4,8,99]
备注：
1≤n≤10^5
/**
 *
 * @param arr int整型一维数组 the array
 * @return int整型
*/
func maxLength(arr []int) int {
	// write code here
	window := make(map[int]int)
	left, right, maxLen := 0, 0, 0
	for right < len(arr) {
		c := arr[right]
		right++

		window[c]++
		if window[c] == 1 { // 增长
			if maxLen < right-left {
				maxLen = right - left
			}
		} else { // 收缩
			for window[c] != 1 { // 收缩到window[c] == 1的时候
				d := arr[left]
				left++
				window[d]--
			}
		}
	}
	return maxLen
}

/**
NC4 判断链表中是否有环
 算法知识视频讲解
简单  通过率：37.84%  时间限制：1秒  空间限制：32M
知识点
链表
题目
题解(108)
讨论(320)
排行
描述
判断给定的链表中是否有环。如果有环则返回true，否则返回false。
你能给出空间复杂度的解法么？
输入分为2部分，第一部分为链表，第二部分代表是否有环，然后回组成head头结点传入到函数里面。-1代表无环，其他的数字代表有环，这些参数解释仅仅是为了方便读者自测调试
示例1
输入：
{3,2,0,-4},1
复制
返回值：
true
复制
说明：
第一部分{3,2,0,-4}代表一个链表，第二部分的1表示，-4到位置1，即-4->2存在一个链接，组成传入的head为一个带环的链表 ,返回true
示例2
输入：
{1},-1
复制
返回值：
false
复制
说明：
第一部分{1}代表一个链表，-1代表无环，组成传入head为一个无环的单链表，返回false
示例3
输入：
{-1,-7,7,-4,19,6,-9,-5,-2,-5},6
复制
返回值：
true
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
 * @return bool布尔型
 */
func hasCycle(head *ListNode) bool {
	// write code here
	slow, fast := head, head              // 同一起跑线，快的走2步，慢的走1步，有环会重叠
	for fast != nil && fast.Next != nil { // 无环会因为 fast == nil || fast.Next == nil跳出
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow { // 有环会因为指针指向重叠跳出
			return true
		}
	}
	return false
}

/**
NC22 合并两个有序的数组
 算法知识视频讲解
简单  通过率：32.41%  时间限制：1秒  空间限制：32M
知识点
数组
双指针
题目
题解(92)
讨论(287)
排行
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
复制
返回值：
[1,2,3,4,5,6]
复制
说明：
A数组为[4,5,6]，B数组为[1,2,3]，后台程序会预先将A扩容为[4,5,6,0,0,0]，B还是为[1,2,3]，m=3，n=3，传入到函数merge里面，然后请同学完成merge函数，将B的数据合并A里面，最后后台程序输出A数组
示例2
输入：
[1,2,3],[2,5,6]
复制
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
	// write code here
	index := 0
	A1 := make([]int, n+m)
	i, j := 0, 0
	for i < m && j < n {
		var cur int
		if A[i] <= B[j] {
			cur = A[i]
			i++
		} else {
			cur = B[j]
			j++
		}
		A1[index] = cur
		index++
	}

	for ; i < m; i++ {
		A1[index] = A[i]
		index++
	}

	for ; j < n; j++ {
		A1[index] = B[j]
		index++
	}

	// A1赋给A
	for i := 0; i < n+m; i++ {
		if len(A) <= i {
			A = append(A, A1[i])
		} else {
			A[i] = A1[i]
		}
	}
	return A
}

/**
NC3 链表中环的入口结点
 算法知识视频讲解
中等  通过率：34.25%  时间限制：1秒  空间限制：64M
知识点
链表
题目
题解(123)
讨论(1k)
排行
描述
给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，返回null。

输入描述：
输入分为2段，第一段是入环前的链表部分，第二段是链表环的部分，后台将这2个会组装成一个有环或者无环单链表
返回值描述：
返回链表的环的入口结点即可。而我们后台程序会打印这个节点
示例1
输入：
{1,2},{3,4,5}
复制
返回值：
3
复制
说明：
返回环形链表入口节点，我们后台会打印该环形链表入口节点，即3
示例2
输入：
{1},{}
复制
返回值：
"null"
复制
说明：
没有环，返回null，后台打印"null"
示例3
输入：
{},{2}
复制
返回值：
2
复制
说明：
只有环形链表节点2，返回节点2，后台打印2
*/
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = pHead
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

/**
NC52 括号序列
 算法知识视频讲解
简单  通过率：34.25%  时间限制：1秒  空间限制：64M
知识点
栈
字符串
题目
题解(76)
讨论(239)
排行
描述
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。
示例1
输入：
"["
复制
返回值：
false
复制
示例2
输入：
"[]"
复制
返回值：
true
*/
/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isBracketsValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	// write code here
	maps := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	stack := []rune{}

	for _, char := range s {
		if v, ok := maps[char]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

/**
NC53 删除链表的倒数第n个节点
 算法知识视频讲解
中等  通过率：34.70%  时间限制：1秒  空间限制：64M
知识点
链表
双指针
题目
题解(70)
讨论(191)
排行
描述
给定一个链表，删除链表的倒数第 nn 个节点并返回链表的头指针
例如，
给出的链表为: 1\to 2\to 3\to 4\to 51→2→3→4→5, n= 2n=2.
删除了链表的倒数第 nn 个节点之后,链表变为1\to 2\to 3\to 51→2→3→5.

备注：
题目保证 nn 一定是有效的
请给出时间复杂度为\ O(n) O(n) 的算法
示例1
输入：
{1,2},2
复制
返回值：
{2}
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
 * @param n int整型
 * @return ListNode类
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// write code here
	if head == nil {
		return head
	}

	// write code here
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}

	if fast == nil {
		return head.Next
	}

	// 找到倒数第n个node的前置结点
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 删除slow.Next结点（倒数第n个结点）
	slow.Next = slow.Next.Next

	return head
}

func SprintNode(head *ListNode) string {
	real := []int{}
	for head != nil {
		real = append(real, head.Val)
		head = head.Next
	}
	return fmt.Sprint(real)
}

func PrintNode(head *ListNode) {
	real := []int{}
	for head != nil {
		real = append(real, head.Val)
		head = head.Next
	}
	fmt.Println(real)
}

/**
NC1 大数加法
 算法知识视频讲解
中等  通过率：40.41%  时间限制：1秒  空间限制：256M
知识点
字符串
模拟
题目
题解(83)
讨论(167)
排行
描述
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回。
（字符串长度不大于100000，保证字符串仅由'0'~'9'这10种字符组成）
示例1
输入：
"1","99"
复制
返回值：
"100"
复制
说明：
1+99=100
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 计算两个数之和
 * @param s string字符串 表示第一个整数
 * @param t string字符串 表示第二个整数
 * @return string字符串
 */
func solve(s string, t string) string {
	// write code here
	lens := len(s)
	lent := len(t)

	for lens < lent {
		s = "0" + s
		lens++
	}

	for lent < lens {
		t = "0" + t
		lent++
	}

	str := ""
	incr := false
	for i := 0; i < lent; i++ {
		bitValue := 0
		intNumber, _ := strconv.Atoi(string(s[lens-i-1]))
		bitValue += intNumber

		intNumber, _ = strconv.Atoi(string(t[lent-i-1]))
		bitValue += intNumber

		if incr {
			bitValue += 1
			incr = false
		}
		if bitValue > 9 {
			incr = true
		}
		str = fmt.Sprint(bitValue%10) + str
	}

	if incr {
		str = "1" + str
	}

	return str
}

/**
NC14 按之字形顺序打印二叉树
 算法知识视频讲解
简单  通过率：26.15%  时间限制：1秒  空间限制：64M
知识点
栈
树
题目
题解(92)
讨论(1k)
排行
描述
给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）
例如：
给定的二叉树是{1,2,3,#,#,4,5}

该二叉树之字形层序遍历的结果是
[
[1],
[3,2],
[4,5]
]
示例1
输入：
{1,2,3,#,#,4,5}
复制
返回值：
[[1],[3,2],[4,5]]
复制
示例2
输入：
{8,6,10,5,7,9,11}
复制
返回值：
[[8],[10,6],[5,7,9,11]]
复制
示例3
输入：
{1,2,3,4,5}
复制
返回值：
[[1],[3,2],[4,5]]
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
 * @param pRoot TreeNode类
 * @return int整型二维数组
 */
func Print(pRoot *TreeNode) [][]int {
	// write code here
	ret := [][]int{}
	if pRoot == nil {
		return ret
	}

	queue := []*TreeNode{}
	queue = append(queue, pRoot)
	i := 1
	for len(queue) > 0 {
		length := len(queue)
		tmpInt := []int{}
		for j := 0; j < length; j++ {
			tmpNode := queue[0]
			queue = queue[1:]
			tmpInt = append(tmpInt, tmpNode.Val)
			if tmpNode.Left != nil {
				queue = append(queue, tmpNode.Left)
			}
			if tmpNode.Right != nil {
				queue = append(queue, tmpNode.Right)
			}
		}

		if i%2 == 0 { // 偶数层时反转tmpInt
			for i, j := 0, len(tmpInt)-1; i < j; i, j = i+1, j-1 {
				tmpInt[i], tmpInt[j] = tmpInt[j], tmpInt[i]
			}
			ret = append(ret, tmpInt)
		} else {
			ret = append(ret, tmpInt)
		}
		i++
	}
	return ret
}

/**
NC127 最长公共子串
 算法知识视频讲解
中等  通过率：32.20%  时间限制：5秒  空间限制：500M
知识点
动态规划
题目
题解(65)
讨论(166)
排行
描述
给定两个字符串str1和str2,输出两个字符串的最长公共子串
题目保证str1和str2的最长公共子串存在且唯一。
示例1
输入：
"1AB2345CD","12345EF"
复制
返回值：
"2345"
复制
备注：
1≤∣str1∣,∣str2∣≤5000
*/
/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */
func LCS(str1 string, str2 string) string {
	// write code here
	m, n := len(str1), len(str2)
	maxLength, maxLastIndex := 0, 0
	var dp [][]int
	for i := 0; i <= m; i++ {
		dp = append(dp, make([]int, n+1))
	}

	// dp[i][j]表示str1[0...i-1]和str2[0...j-1]的最长公共子串
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					maxLastIndex = i
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	// fmt.Println(dp)
	return str1[maxLastIndex-maxLength : maxLastIndex]
}

/**
NC66 两个链表的第一个公共结点
 算法知识视频讲解
简单  通过率：35.98%  时间限制：1秒  空间限制：64M
知识点
链表
题目
题解(118)
讨论(1k)
排行
描述
输入两个无环的单链表，找出它们的第一个公共结点。（注意因为传入数据是链表，所以错误测试数据的提示是用其他方式显示的，保证传入数据是正确的）

示例1
输入：
{1,2,3},{4,5},{6,7}
复制
返回值：
{6,7}
复制
说明：
第一个参数{1,2,3}代表是第一个链表非公共部分，第二个参数{4,5}代表是第二个链表非公共部分，最后的{6,7}表示的是2个链表的公共部分
这3个参数最后在后台会组装成为2个两个无环的单链表，且是有公共节点的
示例2
输入：
{1},{2,3},{}
复制
返回值：
{}
复制
说明：
2个链表没有公共节点 ,返回null，后台打印{}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param pHead1 ListNode类
 * @param pHead2 ListNode类
 * @return ListNode类
 */
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	map1 := make(map[*ListNode]struct{})
	for cur1 := pHead1; cur1 != nil; cur1 = cur1.Next {
		map1[cur1] = struct{}{}
	}

	for cur2 := pHead2; cur2 != nil; cur2 = cur2.Next {
		if _, ok := map1[cur2]; ok {
			return cur2
		}
	}
	return nil
}

/**
NC40 两个链表生成相加链表
 算法知识视频讲解
中等  通过率：31.93%  时间限制：2秒  空间限制：512M
知识点
链表
题目
题解(75)
讨论(229)
排行
描述
假设链表中每一个节点的值都在 0 - 9 之间，那么链表整体就可以代表一个整数。
给定两个这种链表，请生成代表两个整数相加值的结果链表。
例如：链表 1 为 9->3->7，链表 2 为 6->3，最后生成新的结果链表为 1->0->0->0。
示例1
输入：
[9,3,7],[6,3]
复制
返回值：
{1,0,0,0}
复制
备注：
1≤n,m≤10^6
0≤a_i,b_i≤9
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head1 ListNode类
 * @param head2 ListNode类
 * @return ListNode类
 */
func addInList(head1 *ListNode, head2 *ListNode) *ListNode {
	// write code here
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	stack1 := []int{}
	for head1 != nil {
		stack1 = append(stack1, head1.Val)
		head1 = head1.Next
	}

	stack2 := []int{}
	for head2 != nil {
		stack2 = append(stack2, head2.Val)
		head2 = head2.Next
	}

	incr := false
	tmpInt := []int{}
	for len(stack1) > 0 || len(stack2) > 0 {
		tmp := 0
		if len(stack1) > 0 {
			tmp += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack2) > 0 {
			tmp += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		if incr {
			tmp += 1
			incr = false
		}
		if tmp > 9 {
			incr = true
		}

		tmpInt = append(tmpInt, tmp%10)
	}

	if incr {
		tmpInt = append(tmpInt, 1)
	}

	var head, pre *ListNode
	for i := len(tmpInt) - 1; i >= 0; i-- {
		cur := &ListNode{Val: tmpInt[i]}
		if head == nil {
			head = cur
		} else {
			pre.Next = cur
		}
		pre = cur
	}
	return head
}

/**
NC102 在二叉树中找到两个节点的最近公共祖先
 算法知识视频讲解
中等  通过率：45.00%  时间限制：1秒  空间限制：256M
知识点
树
题目
题解(67)
讨论(106)
排行
描述
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。
注：本题保证二叉树中每个节点的val值均不相同。
示例1
输入：
[3,5,1,6,2,0,8,#,#,7,4],5,1
复制
返回值：
3
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
 * @param o1 int整型
 * @param o2 int整型
 * @return int整型
 */
func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	if root == nil {
		return -1
	}

	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}

	left, right := lowestCommonAncestor(root.Left, o1, o2), lowestCommonAncestor(root.Right, o1, o2)

	if left != -1 && right != -1 {
		return root.Val
	}

	if left == -1 && right == -1 {
		return -1
	}

	if left != -1 {
		return left
	} else {
		return right
	}
}

/**
NC103 反转字符串
 算法知识视频讲解
入门  通过率：61.31%  时间限制：1秒  空间限制：64M
知识点
双指针
字符串
题目
题解(66)
讨论(201)
排行
描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）
示例1
输入：
"abcd"
复制
返回值：
"dcba"
*/
/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func reverseString(str string) string {
	// write code here
	if len(str) <= 1 {
		return str
	}

	ans := []byte(str)
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return string(ans)
}

/**
NC38 螺旋矩阵
 算法知识视频讲解
入门  通过率：27.43%  时间限制：1秒  空间限制：64M
知识点
数组
题目
题解(56)
讨论(171)
排行
描述
给定一个m x n大小的矩阵（m行，n列），按螺旋的顺序返回矩阵中的所有元素。
示例1
输入：
[[1,2,3],[4,5,6],[7,8,9]]
复制
返回值：
[1,2,3,6,9,8,7,4,5]
*/
/**
 *
 * @param matrix int整型二维数组
 * @return int整型一维数组
 */
func spiralOrder(matrix [][]int) []int {
	// write code here
	ret := []int{}
	if len(matrix) == 0 { // 空数组，直接返回
		return ret
	}
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top++
		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right--
		if top > bottom || left > right { // 提前退出，防止重复输出
			break
		}
		for i := right; i >= left; i-- {
			ret = append(ret, matrix[bottom][i])
		}
		bottom--
		for i := bottom; i >= top; i-- {
			ret = append(ret, matrix[i][left])
		}
		left++
	}
	return ret
}

/**
NC65 斐波那契数列
 算法知识视频讲解
入门  通过率：34.37%  时间限制：1秒  空间限制：64M
知识点
数组
题目
题解(146)
讨论(2k)
排行
描述
大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。
n≤39

示例1
输入：
4
复制
返回值：
3
*/
/**
 *
 * @param n int整型
 * @return int整型
 */
func Fibonacci(n int) int {
	// write code here
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

/**
NC17 最长回文子串
 算法知识视频讲解
中等  通过率：40.07%  时间限制：3秒  空间限制：64M
知识点
字符串
动态规划
题目
题解(97)
讨论(294)
排行
描述
对于一个字符串，请设计一个高效算法，计算其中最长回文子串的长度。

给定字符串A以及它的长度n，请返回最长回文子串的长度。

示例1
输入：
"abc1234321ab",12
复制
返回值：
7
*/

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param A string字符串
 * @param n int整型
 * @return int整型
 */
// 寻找回文串的问题核心思想是：从中间开始向两边扩散来判断回文串
func getLongestPalindrome(A string, n int) int {
	// write code here
	maxLen := 0
	for i := 0; i < len(A); i++ {
		len1 := Palindrome(A, i, i)
		len2 := Palindrome(A, i, i+1)
		if len1 > maxLen {
			maxLen = len1
		}
		if len2 > maxLen {
			maxLen = len2
		}
	}
	return maxLen
}

// 校验是否为回文串
func Palindrome(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] { // 分别向两边扩散
		l--
		r++
	}
	return r - l - 1
}

/**
NC54 数组中相加和为0的三元组
 算法知识视频讲解
中等  通过率：24.71%  时间限制：1秒  空间限制：64M
知识点
数组
双指针
题目
题解(47)
讨论(186)
排行
描述
给出一个有n个元素的数组S，S中是否有元素a,b,c满足a+b+c=0？找出数组S中所有满足条件的三元组。
注意：
三元组（a、b、c）中的元素必须按非降序排列。（即a≤b≤c）
解集中不能包含重复的三元组。
例如，给定的数组 S = {-10 0 10 20 -10 -40},解集为(-10, -10, 20),(-10, 0, 10)
0 <= S.length <= 1000
示例1
输入：
[0]
复制
返回值：
[]
复制
示例2
输入：
[-2,0,1,1,2]
复制
返回值：
[[-2,0,2],[-2,1,1]]
复制
示例3
输入：
[-10,0,10,20,-10,-40]
复制
返回值：
[[-10,-10,20],[-10,0,10]]
*/
/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func threeSum(num []int) [][]int {
	// write code here
	return threeSumTarget(num, 0)
}

func threeSumTarget(num []int, target int) [][]int {
	sort.Ints(num)
	ret := [][]int{}
	n := len(num)
	for i := 0; i < n; i++ {
		tuples := twoSumTarget(num, i+1, target-num[i])
		for _, tuple := range tuples {
			tuple = append([]int{num[i]}, tuple...)
			ret = append(ret, tuple)
		}

		for i < n-1 && num[i] == num[i+1] { // 跳过重复的
			i++
		}
	}

	return ret
}

func twoSumTarget(num []int, start, target int) [][]int {
	lo, hi := start, len(num)-1
	ret := [][]int{}
	for lo < hi {
		sum := num[lo] + num[hi]
		left, right := num[lo], num[hi]
		if sum < target {
			for lo < hi && num[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && num[hi] == right {
				hi--
			}
		} else {
			ret = append(ret, []int{num[lo], num[hi]})
			for lo < hi && num[lo] == left {
				lo++
			}
			for lo < hi && num[hi] == right {
				hi--
			}
		}
	}
	return ret
}

/**
NC12 重建二叉树
 算法知识视频讲解
中等  通过率：26.07%  时间限制：1秒  空间限制：64M
知识点
树
dfs
数组
题目
题解(154)
讨论(2k)
排行
描述
给定某二叉树的前序遍历和中序遍历，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。

提示:
1.0 <= pre.length <= 2000
2.vin.length == pre.length
3.-10000 <= pre[i], vin[i] <= 10000
4.pre 和 vin 均无重复元素
5.vin出现的元素均出现在 pre里
6.只需要返回根结点，系统会自动输出整颗树做答案对比
示例1
输入：
[1,2,4,7,3,5,6,8],[4,7,2,1,5,3,8,6]
复制
返回值：
{1,2,3,4,#,5,6,#,7,#,#,8}
复制
说明：
返回根节点，系统会输出整颗二叉树对比结果
示例2
输入：
[1],[1]
复制
返回值：
{1}
复制
示例3
输入：
[1,2,3,4,5,6,7],[3,2,4,1,6,5,7]
复制
返回值：
{1,2,5,3,4,6,7}
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
 * @param pre int整型一维数组
 * @param vin int整型一维数组
 * @return TreeNode类
 */
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 || len(vin) == 0 {
		return nil
	}
	root := &TreeNode{Val: pre[0]}
	for i := 0; i < len(vin); i++ {
		if vin[i] == pre[0] {
			root.Left = reConstructBinaryTree(pre[1:i+1], vin[:i])
			root.Right = reConstructBinaryTree(pre[i+1:], vin[i+1:])
			break
		}
	}
	return root
}

func PrintTreeNode(root *TreeNode) {
	fmt.Println(preOrder(root))
}

func SprintTreeNode(root *TreeNode) string {
	return fmt.Sprint(preOrder(root))
}

/**
NC91 最长递增子序列
 算法知识视频讲解
中等  通过率：27.17%  时间限制：2秒  空间限制：256M
知识点
二分
动态规划
题目
题解(40)
讨论(93)
排行
描述
给定数组arr，设长度为n，输出arr的最长递增子序列。（如果有多个答案，请输出其中 按数值(注：区别于按单个字符的ASCII码值)进行比较的 字典序最小的那个）
示例1
输入：
[2,1,5,3,6,4,8,9,7]
复制
返回值：
[1,3,4,8,9]
复制
示例2
输入：
[1,2,8,6,4]
复制
返回值：
[1,2,4]
复制
说明：
其最长递增子序列有3个，（1，2，8）、（1，2，6）、（1，2，4）其中第三个 按数值进行比较的字典序 最小，故答案为（1，2，4）
备注：
n≤10^5
1≤arr_i≤10^9
*/
/**
 * retrun the longest increasing subsequence
 * @param arr int整型一维数组 the array
 * @return int整型一维数组
 */
func LIS(arr []int) []int {
	// write code here
	if len(arr) < 2 {
		return arr
	}

	dp := make([][]int, len(arr)) // dp[i]表示以arr[i]结尾的[0:i]中最长递增子序列
	for i := 0; i < len(arr); i++ {
		dp[i] = []int{arr[i]}
	}
	for i := 0; i < len(arr); i++ {
		maxIndex := i
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				if len(dp[j]) == len(dp[maxIndex]) { // 相同，取较小的
					if fmt.Sprint(dp[j]) < fmt.Sprint(dp[maxIndex]) {
						maxIndex = j
					}
				} else if len(dp[j]) > len(dp[maxIndex]) {
					maxIndex = j
				}
			}
		}
		if maxIndex != i {
			// dp[i] = append(dp[maxIndex], arr[i]) // append会改变dp[]maxIndex的值
			dp[i] = make([]int, len(dp[maxIndex]))
			copy(dp[i], dp[maxIndex])
			dp[i] = append(dp[i], arr[i])
		}
	}

	var LIS []int
	for i := len(dp) - 1; i >= 0; i-- {
		if len(dp[i]) > len(LIS) { // 大于则替换
			LIS = dp[i]
		}
	}
	return LIS
}

/**
NC32 求平方根
 算法知识视频讲解
简单  通过率：33.10%  时间限制：1秒  空间限制：64M
知识点
数学
二分
题目
题解(36)
讨论(135)
排行
描述
实现函数 int sqrt(int x).
计算并返回x的平方根（向下取整）
示例1
输入：
2
复制
返回值：
1
*/
/**
 *
 * @param x int整型
 * @return int整型
 */
func sqrt(x int) int {
	// write code here
	ret := 0
	for ret*ret <= x {
		ret++
	}
	return ret - 1
}

/**
NC48 在旋转过的有序数组中寻找目标值
 算法知识视频讲解
简单  通过率：48.72%  时间限制：1秒  空间限制：256M
知识点
二分
题目
题解(29)
讨论(41)
排行
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
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
/* func search(nums []int, target int) int {
	// write code here
} */

/**
NC90 包含min函数的栈
 算法知识视频讲解
简单  通过率：33.58%  时间限制：1秒  空间限制：64M
知识点
栈
题目
题解(89)
讨论(1k)
排行
描述
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的min函数，并且调用 min函数、push函数 及 pop函数 的时间复杂度都是 O(1)
push(value):将value压入栈中
pop():弹出栈顶元素
top():获取栈顶元素
min():获取栈中最小元素

示例:
输入:    ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]
输出:    -1,2,1,-1
解析:
"PSH-1"表示将-1压入栈中，栈中元素为-1
"PSH2"表示将2压入栈中，栈中元素为2，-1
“MIN”表示获取此时栈中最小元素==>返回-1
"TOP"表示获取栈顶元素==>返回2
"POP"表示弹出栈顶元素，弹出2，栈中元素为-1
"PSH-1"表示将1压入栈中，栈中元素为1，-1
"TOP"表示获取栈顶元素==>返回1
“MIN”表示获取此时栈中最小元素==>返回-1

示例1
输入：
 ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]
复制
返回值：
-1,2,1,-1
*/
/* func Push(node int) {
	// write code here
}
func Pop() {
	// write code here
}
func Top() int {
	// write code here
}
func Min() int {
	// write code here
} */

/**
NC7 买卖股票的最好时机
 算法知识视频讲解
简单  通过率：46.97%  时间限制：1秒  空间限制：64M
知识点
数组
动态规划
题目
题解(68)
讨论(217)
排行
描述
假设你有一个数组，其中第\ i i 个元素是股票在第\ i i 天的价格。
你有一次买入和卖出的机会。（只有买入了股票以后才能卖出）。请你设计一个算法来计算可以获得的最大收益。
示例1
输入：
[1,4,2]
复制
返回值：
3
复制
示例2
输入：
[2,4,1]
复制
返回值：
2
*/
/**
 *
 * @param prices int整型一维数组
 * @return int整型
 */
/* func maxProfit(prices []int) int {
	// write code here
} */

/**
NC51 合并k个已排序的链表
 算法知识视频讲解
较难  通过率：35.41%  时间限制：1秒  空间限制：64M
知识点
堆
链表
分治
题目
题解(46)
讨论(133)
排行
描述
合并\ k k 个已排序的链表并将其作为一个已排序的链表返回。分析并描述其复杂度。
示例1
输入：
[{1,2,3},{4,5,6,7}]
复制
返回值：
{1,2,3,4,5,6,7}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param lists ListNode类一维数组
 * @return ListNode类
 */
/* func mergeKLists(lists []*ListNode) *ListNode {
	// write code here
} */

/**
NC121 字符串的排列
 算法知识视频讲解
较难  通过率：22.80%  时间限制：1秒  空间限制：64M
知识点
字符串
动态规划
题目
题解(130)
讨论(1k)
排行
描述
输入一个字符串，打印出该字符串中字符的所有排列，你可以以任意顺序返回这个字符串数组。例如输入字符串abc,则按字典序打印出由字符a,b,c所能排列出来的所有字符串abc,acb,bac,bca,cab和cba。
输入描述：
输入一个字符串,长度不超过9(可能有字符重复),字符只包括大小写字母。
示例1
输入：
"ab"
复制
返回值：
["ab","ba"]
复制
说明：
返回["ba","ab"]也是正确的
示例2
输入：
"aab"
复制
返回值：
["aab","aba","baa"]
复制
示例3
输入：
"abc"
复制
返回值：
["abc","acb","bac","bca","cab","cba"]
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串
 * @return string字符串一维数组
 */
/* func Permutation(str string) []string {
	// write code here
} */

/**
NC128 接雨水问题
 算法知识视频讲解
较难  通过率：35.70%  时间限制：3秒  空间限制：256M
知识点
双指针
题目
题解(70)
讨论(124)
排行
描述
给定一个整形数组arr，已知其中所有的值都是非负的，将这个数组看作一个柱子高度图，计算按此排列的柱子，下雨之后能接多少雨水。

示例1
输入：
[3,1,2,5,2,4]
复制
返回值：
5
复制
说明：
数组 [3,1,2,5,2,4] 表示柱子高度图，在这种情况下，可以接 5个单位的雨水，蓝色的为雨水
示例2
输入：
[4,5,1,3,2]
复制
返回值：
2
复制
备注：
1≤N≤10^6
*/
/**
 * max water
 * @param arr int整型一维数组 the array
 * @return long长整型
 */
/* func maxWater(arr []int) int64 {
	// write code here
} */

/**
NC136 输出二叉树的右视图
 算法知识视频讲解
中等  通过率：52.45%  时间限制：1秒  空间限制：256M
知识点
树
题目
题解(31)
讨论(77)
排行
描述
请根据二叉树的前序遍历，中序遍历恢复二叉树，并打印出二叉树的右视图

示例1
输入：
[1,2,4,5,3],[4,2,5,1,3]
复制
返回值：
[1,3,5]
复制
备注：
二叉树每个节点的值在区间[1,10000]内，且保证每个节点的值互不相同。
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 求二叉树的右视图
 * @param xianxu int整型一维数组 先序遍历
 * @param zhongxu int整型一维数组 中序遍历
 * @return int整型一维数组
 */
/* func solve(xianxu []int, zhongxu []int) []int {
	// write code here
}*/

/**
NC109 岛屿数量
 算法知识视频讲解
中等  通过率：36.78%  时间限制：1秒  空间限制：256M
知识点
dfs
bfs
搜索
题目
题解(39)
讨论(98)
排行
描述
给一个01矩阵，1代表是陆地，0代表海洋， 如果两个1相邻，那么这两个1属于同一个岛。我们只考虑上下左右为相邻。
岛屿: 相邻陆地可以组成一个岛屿（相邻:上下左右） 判断岛屿个数。
示例1
输入：
[[1,1,0,0,0],[0,1,0,1,1],[0,0,0,1,1],[0,0,0,0,0],[0,0,1,1,1]]
复制
返回值：
3
复制
备注：
01矩阵范围<=200*200
*/
/**
 * 判断岛屿数量
 * @param grid char字符型二维数组
 * @return int整型
 */
/* func solve(grid [][]byte) int {
	// write code here
} */

/**
NC13 二叉树的最大深度
 算法知识视频讲解
简单  通过率：57.44%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(45)
讨论(209)
排行
描述
求给定二叉树的最大深度，
最大深度是指树的根结点到最远叶子结点的最长路径上结点的数量。

示例1
输入：
{1,2}
复制
返回值：
2
复制
示例2
输入：
{1,2,3,4,#,#,5}
复制
返回值：3
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
/* func maxDepth(root *TreeNode) int {
	// write code here
}
*/
/**
NC141 判断回文
 算法知识视频讲解
入门  通过率：61.46%  时间限制：1秒  空间限制：256M
知识点
字符串
题目
题解(42)
讨论(103)
排行
描述
给定一个字符串，请编写一个函数判断该字符串是否回文。如果回文请返回true，否则返回false。
示例1
输入：
"absba"
复制
返回值：
true
复制
示例2
输入：
"ranko"
复制
返回值：
false
复制
示例3
输入：
"yamatomaya"
复制
返回值：
false
复制
示例4
输入：
"a"
复制
返回值：
true
复制
备注：
字符串长度不大于1000000，且仅由小写字母组成
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param str string字符串 待判断的字符串
 * @return bool布尔型
 */
/* func judge(str string) bool {
	// write code here
} */

/**
NC70 单链表的排序
 算法知识视频讲解
简单  通过率：42.11%  时间限制：2秒  空间限制：256M
知识点
链表
排序
题目
题解(43)
讨论(107)
排行
描述
给定一个无序单链表，实现单链表的排序(按升序排序)。
示例1
输入：
[1,3,2,4,5]
复制
返回值：
{1,2,3,4,5}
*/
/*
 * type ListNode struct{
 *   Val int
 *   Next *ListNode
 * }
 */

/**
 *
 * @param head ListNode类 the head node
 * @return ListNode类
 */
/* func sortInList(head *ListNode) *ListNode {
	// write code here
} */

/**
NC62 平衡二叉树
 算法知识视频讲解
简单  通过率：37.88%  时间限制：1秒  空间限制：64M
知识点
树
dfs
题目
题解(94)
讨论(1k)
排行
描述
输入一棵二叉树，判断该二叉树是否是平衡二叉树。
在这里，我们只需要考虑其平衡性，不需要考虑其是不是排序二叉树
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树。

注：我们约定空树是平衡二叉树。
示例1
输入：
{1,2,3,4,5,6,7}
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
 * @param pRoot TreeNode类
 * @return bool布尔型
 */
/* func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
} */

/**
NC73 数组中出现次数超过一半的数字
 算法知识视频讲解
简单  通过率：31.10%  时间限制：1秒  空间限制：64M
知识点
哈希
数组
题目
题解(121)
讨论(2k)
排行
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
*/
/**
 *
 * @param numbers int整型一维数组
 * @return int整型
 */
/* func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
} */

/**
NC59 矩阵的最小路径和
 算法知识视频讲解
中等  通过率：46.61%  时间限制：1秒  空间限制：256M
知识点
数组
动态规划
题目
题解(44)
讨论(102)
排行
描述
给定一个 n * m 的矩阵 a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是路径和，输出所有的路径中最小的路径和。
示例1
输入：
[[1,3,5,9],[8,1,3,4],[5,0,6,1],[8,8,4,0]]
复制
返回值：
12
复制
备注：
1≤n,m≤2000
1≤arr_i,j≤100
*/
/**
 *
 * @param matrix int整型二维数组 the matrix
 * @return int整型
 */
/* func minPathSum(matrix [][]int) int {
	// write code here
} */

/**
NC137 表达式求值
 算法知识视频讲解
中等  通过率：46.93%  时间限制：1秒  空间限制：256M
知识点
栈
递归
题目
题解(40)
讨论(64)
排行
描述
请写一个整数计算器，支持加减乘三种运算和括号。
示例1
输入：
"1+2"
复制
返回值：
3
复制
示例2
输入：
"(2*(3-4))*5"
复制
返回值：
-10
复制
示例3
输入：
"3+2*3*4-1"
复制
返回值：
26
*/
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 返回表达式的值
 * @param s string字符串 待计算的表达式
 * @return int整型
 */
/* func solve(s string) int {
	// write code here
} */

/**
NC97 字符串出现次数的TopK问题
 算法知识视频讲解
中等  通过率：40.58%  时间限制：2秒  空间限制：256M
知识点
哈希
题目
题解(38)
讨论(75)
排行
描述
给定一个字符串数组，再给定整数k，请返回出现次数前k名的字符串和对应的次数。
返回的答案应该按字符串出现频率由高到低排序。如果不同的字符串有相同出现频率，按字典序排序。
对于两个字符串，大小关系取决于两个字符串从左到右第一个不同字符的 ASCII 值的大小关系。
比如"ah1x"小于"ahb"，"231"<”32“
字符仅包含数字和字母

[要求]
如果字符串数组长度为N，时间复杂度请达到O(N \log K)O(NlogK)

示例1
输入：
["a","b","c","b"],2
复制
返回值：
[["b","2"],["a","1"]]
复制
说明：
"b"出现了2次，记["b","2"]，"a"与"c"各出现1次，但是a字典序在c前面，记["a","1"]，最后返回[["b","2"],["a","1"]]

示例2
输入：
["123","123","231","32"],2
复制
返回值：
[["123","2"],["231","1"]]
复制
说明：
 "123"出现了2次，记["123","2"]，"231"与"32"各出现1次，但是"231"字典序在"32"前面，记["231","1"]，最后返回[["123","2"],["231","1"]]
示例3
输入：
["abcd","abcd","abcd","pwb2","abcd","pwb2","p12"],3
复制
返回值：
[["abcd","4"],["pwb2","2"],["p12","1"]]
*/
/**
 * return topK string
 * @param strings string字符串一维数组 strings
 * @param k int整型 the k
 * @return string字符串二维数组
 */
/* func topKstrings(strings []string, k int) [][]string {
	// write code here
} */

/**
NC112 进制转换
 算法知识视频讲解
简单  通过率：34.15%  时间限制：1秒  空间限制：64M
知识点
位运算
题目
题解(19)
讨论(75)
排行
描述
给定一个十进制数 M ，以及需要转换的进制数 N 。将十进制数 M 转化为 N 进制数。

当 N 大于 10 以后， 应在结果中使用大写字母表示大于 10 的一位，如 'A' 表示此位为 10 ， 'B' 表示此位为 11 。

若 M 为负数，应在结果中保留负号。
示例1
输入：
7,2
复制
返回值：
"111"
复制
备注：
M是32位整数，2<=N<=16.
*/
/**
 * 进制转换
 * @param M int整型 给定整数
 * @param N int整型 转换到的进制
 * @return string字符串
 */
/* func solve(M int, N int) string {
	// write code here
} */
