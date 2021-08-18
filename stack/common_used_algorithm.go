package stack

// ------------------------------ 有效的括号 -------------------------------

/*

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。

输入：s = "()"
输出：true

输入：s = "()[]{}"
输出：true

输入：s = "(]"
输出：false

输入：s = "([)]"
输出：false

输入：s = "{[]}"
输出：true

提示：
1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

/* func isValidBrackets(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	leftStack := NewListStack()
	rightStack := NewListStack()
	for i := 0; i < len(s); i++ {
		char := s[i]
		switch char {
		case '(':
			leftStack.Push(char)
		case '{':
			leftStack.Push(char)
		case '[':
			leftStack.Push(char)
		case ')':
			rightStack.Push(char)
		case '}':
			rightStack.Push(char)
		case ']':
			rightStack.Push(char)
		}
	}
	fmt.Printf("s = %s\n", s)
	if leftStack.length != rightStack.length {
		return false
	}
	for !leftStack.Empty() {
		left := leftStack.Pop().(byte)
		right := rightStack.Pop().(byte)
		fmt.Printf("%v ", left)
		fmt.Printf("%v \n", right)
		if !(left == '(' && right == ')' || left == '{' && right == '}' || left == '[' && right == ']') {
			return false
		}
	}
	return true
} */

func isValidBrackets(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { // 如果是右括号
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] { // 没有左括号可用，或两者不匹配，直接返回false
				return false
			}
			stack = stack[:len(stack)-1] // 移除栈顶的左括号
		} else {
			stack = append(stack, s[i]) // 左括号直接入栈
		}
	}
	return true
}
