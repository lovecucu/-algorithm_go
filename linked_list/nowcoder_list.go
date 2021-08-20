package linked_list

/*

反转链表

描述:
输入一个链表，反转链表后，输出新链表的表头。

示例1
输入：{1,2,3}
返回值：{3,2,1}

**/

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
	var pre *ListNode
	for cur := pHead; cur != nil; {
		tmp := cur.Next
		if pre == nil {
			pre = cur
		} else {
			cur.Next = pre
			pre = cur
		}
		cur = tmp
	}
	return pre
}

/*

判断链表中是否有环

描述:
判断给定的链表中是否有环。如果有环则返回true，否则返回false。
你能给出空间复杂度的解法么？
输入分为2部分，第一部分为链表，第二部分代表是否有环，然后回组成head头结点传入到函数里面。-1代表无环，其他的数字代表有环，这些参数解释仅仅是为了方便读者自测调试

示例1
输入：{3,2,0,-4},1
返回值：true
说明：
第一部分{3,2,0,-4}代表一个链表，第二部分的1表示，-4到位置1，即-4->2存在一个链接，组成传入的head为一个带环的链表 ,返回true

示例2
输入：{1},-1
返回值：false
说明：
第一部分{1}代表一个链表，-1代表无环，组成传入head为一个无环的单链表，返回false

示例3
输入：{-1,-7,7,-4,19,6,-9,-5,-2,-5},6
返回值：true

**/

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
	return false
}

/**

合并两个有序链表


描述
输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
示例1
输入：
{1,3,5},{2,4,6}
返回值：
{1,2,3,4,5,6}

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
// func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
// 	// write code here
// }

/**

链表排序

描述
给定一个无序单链表，实现单链表的排序(按升序排序)。
示例1
输入：
[1,3,2,4,5]
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
// func sortInList(head *ListNode) *ListNode {
// 	// write code here
// }

/**

判断链表是否为回文结构

描述
给定一个链表，请判断该链表是否为回文结构。
示例1
输入：
[1]
返回值：
true

示例2
输入：
[2,1]
返回值：
false
说明：
2->1

示例3
输入：
[1,2,2,1]
返回值：
true
说明：
1->2->2->1

备注：
1≤n≤10^6

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
// func isPail(head *ListNode) bool {
// 	// write code here
// }

/**

两个链表中第一个公共结点

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
// func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
// 	// write code here
// }

/**

删除有序链表中重复的元素

描述
删除给出链表中的重复元素（链表中元素从小到大有序），使链表中的所有元素都只出现一次
例如：
给出的链表为1\to1\to21→1→2,返回1 \to 21→2.
给出的链表为1\to1\to 2 \to 3 \to 31→1→2→3→3,返回1\to 2 \to 31→2→3.

示例1
输入：
{1,1,2}
复制
返回值：
{1,2}
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
 * @param head ListNode类
 * @return ListNode类
 */
// func deleteDuplicates(head *ListNode) *ListNode {
// 	// write code here
// }
