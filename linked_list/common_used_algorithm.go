package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// ------------------------------ 反转链表 -------------------------------

// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]
// 输入：head = []
// 输出：[]

// 提示：
// 链表中节点的数目范围是 [0, 5000]
// -5000 <= Node.val <= 5000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归解法
func reverseListRecursion(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		next := cur.Next // 暂存下一个结点
		cur.Next = pre   // cur指向pre
		pre = cur        // cur变成最新的pre
		cur = next       // 指针移到下一个结点
	}
	return pre
}

// 迭代解法
func reverseListIteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 获取head.Next的头结点
	newHead := reverseListIteration(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// ------------------------------ 合并两个升序链表 -------------------------------

// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 输入：l1 = [], l2 = []
// 输出：[]
// 输入：l1 = [], l2 = [0]
// 输出：[0]

// 提示：
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列

/**
 * Definition for singly-linked list.

 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var merge, cur, pre *ListNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val { // l1小，则l1指针右移
			cur = l1
			l1 = l1.Next
		} else { // 相反，则l2指针右移
			cur = l2
			l2 = l2.Next
		}
		if merge == nil { // 初始化merge
			merge = cur
		} else {
			pre.Next = cur
		}
		pre = cur

		// 有一个nil，则把另外一个接到最后
		if l1 == nil {
			pre.Next = l2
			break
		} else if l2 == nil {
			pre.Next = l1
			break
		}
	}
	return merge
}

// ------------------------------ 求链表的中间结点 -------------------------------

// 输入：[1,2,3,4,5]
// 输出：此列表中的结点 3 (序列化形式：[3,4,5])
// 返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
// 注意，我们返回了一个 ListNode 类型的对象 ans，这样：
// ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.

// 输入：[1,2,3,4,5,6]
// 输出：此列表中的结点 4 (序列化形式：[4,5,6])
// 由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。

// 提示：
// 给定链表的结点数介于 1 和 100 之间。

// 快慢指针解法（还有数组法、单指针法-N/2）
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil { // fast不是最后一个节点
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// ------------------------------ 删除链表的倒数第 N 个结点 -------------------------------

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 输入：head = [1], n = 1
// 输出：[]
// 输入：head = [1,2], n = 1
// 输出：[1]

// 提示：
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}
	var nthNode *ListNode

	// 找nthNode
	i := 0
	for tmp := head; i < n && tmp != nil; {
		nthNode = tmp
		i++
		tmp = tmp.Next
	}
	if i < n { // 说明list长度小于n
		return head
	}

	slow := head      // 1
	var pre *ListNode // nil
	fast := nthNode   // 5

	if fast.Next == nil { // 最后一位
		head = head.Next
	} else {
		// nthNode先到终点，此时slow对应的就是倒数n结点
		for fast.Next != nil {
			pre = slow
			slow = slow.Next
			fast = fast.Next
		}
		pre.Next = slow.Next
		slow.Next = nil
	}

	return head
}
