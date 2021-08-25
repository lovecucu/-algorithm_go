package stack

import "fmt"

/**

两个栈实现队列

描述:
用两个栈来实现一个队列，分别完成在队列尾部插入整数(push)和在队列头部删除整数(pop)的功能。 队列中的元素为int类型。保证操作合法，即保证pop操作时队列内已有元素。

示例:
输入:["PSH1","PSH2","POP","POP"]
返回:1,2
解析:
"PSH1":代表将1插入队列尾部
"PSH2":代表将2插入队列尾部
"POP“:代表删除一个元素，先进先出=>返回1
"POP“:代表删除一个元素，先进先出=>返回2

示例1
输入：["PSH1","PSH2","POP","POP"]
返回值：1,2
*/

var stack1 []int
var stack2 []int

func QueuePush(node int) {
	stack1 = append(stack1, node)
}

func QueuePop() int {
	len1 := len(stack1)
	if len(stack2) == 0 && len1 > 0 {
		for i := len1 - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = stack1[0:0]
	}
	if len(stack2) > 0 {
		node := stack2[len(stack2)-1]
		stack2 = stack2[0 : len(stack2)-1]
		return node
	}
	return -1
}

/**

括号序列

描述
给出一个仅包含字符'(',')','{','}','['和']',的字符串，判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列，但"(]"和"([)]"不合法。

示例1
输入：
"["
返回值：
false

示例2
输入：
"[]"
返回值：
true

*/

/**
 *
 * @param s string字符串
 * @return bool布尔型
 */
func isValid(s string) bool {
	maps := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	var stack1 []rune

	for _, b := range s {
		if v, ok := maps[b]; ok {
			if len(stack1) == 0 || stack1[len(stack1)-1] != v {
				return false
			} else {
				stack1 = stack1[0 : len(stack1)-1]
			}
		} else {
			stack1 = append(stack1, b)
		}
		fmt.Println(stack1)
	}

	return len(stack1) == 0
}

/**

包含min函数的栈

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

func Push(node int) {
	// write code here
}
func Pop() {
	// write code here
}
func Top() int {
	// write code here
	return -1
}
func Min() int {
	// write code here
	return -1
}
