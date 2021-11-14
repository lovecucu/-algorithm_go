package crackingcodinginterview

/**
面试题 01.01. 判定字符是否唯一
实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

示例 1：

输入: s = "leetcode"
输出: false
示例 2：

输入: s = "abc"
输出: true
限制：

0 <= len(s) <= 100
如果你不使用额外的数据结构，会很加分。
*/
func isUnique(astr string) bool {
	checker := 0
	// 核心思路：用bit位代替bool数组（确认字符范围，如果是ascii，则需要128位；如果是a-z只需要26位；）
	for i := 0; i < len(astr); i++ {
		val := astr[i] - 'a'
		if (checker & (1 << val)) > 0 { // val位对应的字母出现过，直接报错，相当于nums[val]是否存在
			return false
		}
		checker |= (1 << val) // val位置1，相当于nums[val] = 1
	}
	return true
}

/**
面试题 01.02. 判定是否互为字符重排
给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

示例 1：

输入: s1 = "abc", s2 = "bca"
输出: true
示例 2：

输入: s1 = "abc", s2 = "bad"
输出: false
说明：

0 <= len(s1) <= 100
0 <= len(s2) <= 100
*/
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	maps := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		if v, ok := maps[s1[i]]; ok {
			maps[s1[i]] = v + 1
		} else {
			maps[s1[i]] = 1
		}
	}

	for i := 0; i < len(s2); i++ {
		v, ok := maps[s2[i]]
		if !ok {
			return false
		}

		maps[s2[i]] = v - 1
		if maps[s2[i]] < 0 {
			return false
		}
	}

	return true
}

/**
面试题 01.03. URL化
URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）



示例 1：

输入："Mr John Smith    ", 13
输出："Mr%20John%20Smith"
示例 2：

输入："               ", 5
输出："%20%20%20%20%20"


提示：

字符串长度在 [0, 500000] 范围内。
*/
func replaceSpaces(S string, length int) string {
	spaceCount := 0 // 先找出空格数
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			spaceCount++
		}
	}

	trueLength := length + spaceCount*2 // 确定要返回的字符串真实长度
	bytes := []byte(S)[:trueLength]
	for i := length - 1; i >= 0; i-- { // 倒序填充字符串
		if bytes[i] == ' ' {
			bytes[trueLength-1] = '0'
			bytes[trueLength-2] = '2'
			bytes[trueLength-3] = '%'
			trueLength -= 3
		} else {
			bytes[trueLength-1] = bytes[i]
			trueLength--
		}
	}
	return string(bytes)
}

/**
面试题 01.04. 回文排列
给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。

回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。

回文串不一定是字典当中的单词。



示例1：

输入："tactcoa"
输出：true（排列有"tacocat"、"atcocta"，等等）
*/
// func canPermutePalindrome(s string) bool {

// }

/**
面试题 01.05. 一次编辑
字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。



示例 1:

输入:
first = "pale"
second = "ple"
输出: True


示例 2:

输入:
first = "pales"
second = "pal"
输出: False
*/
// func oneEditAway(first string, second string) bool {

// }

/**
面试题 01.06. 字符串压缩
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

 输入："aabcccccaaa"
 输出："a2b1c5a3"
示例2:

 输入："abbccd"
 输出："abbccd"
 解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。
提示：

字符串长度在[0, 50000]范围内。
*/
// func compressString(S string) string {

// }

/**
面试题 01.07. 旋转矩阵
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？



示例 1:

给定 matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
注意：本题与主站 48 题相同：https://leetcode-cn.com/problems/rotate-image/
*/
// func rotate(matrix [][]int) {

// }

/**
面试题 01.08. 零矩阵
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。



示例 1：

输入：
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出：
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2：

输入：
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出：
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
*/
// func setZeroes(matrix [][]int) {

// }

/**
面试题 01.09. 字符串轮转
字符串轮转。给定两个字符串s1和s2，请编写代码检查s2是否为s1旋转而成（比如，waterbottle是erbottlewat旋转后的字符串）。

示例1:

 输入：s1 = "waterbottle", s2 = "erbottlewat"
 输出：True
示例2:

 输入：s1 = "aa", s2 = "aba"
 输出：False
提示：

字符串长度在[0, 100000]范围内。
说明:

你能只调用一次检查子串的方法吗？
*/
// func isFlipedString(s1 string, s2 string) bool {

// }
