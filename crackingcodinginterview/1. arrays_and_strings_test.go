package crackingcodinginterview

import (
	"fmt"
	"testing"
)

func TestIsUnique(t *testing.T) {
	if !isUnique("abc") || isUnique("aabc") || isUnique("leetcode") {
		t.Error(`TestIsUnique failed`)
	}
}

func TestCheckPermutation(t *testing.T) {
	if !CheckPermutation("abc", "bca") || CheckPermutation("abc", "bad") || CheckPermutation("abb", "aab") {
		t.Error(`TestCheckPermutation failed`)
	}
}

func TestReplaceSpaces(t *testing.T) {
	if replaceSpaces("Mr John Smith    ", 13) != "Mr%20John%20Smith" || replaceSpaces("                ", 5) != "%20%20%20%20%20" {
		t.Error(`TestReplaceSpaces failed`)
	}
}

func TestCanPermutePalindrome(t *testing.T) {
	if !canPermutePalindrome("tactcoa") || canPermutePalindrome("aabc") {
		t.Error(`TestCanPermutePalindrome failed`)
	}
}

func TestOneEditAway(t *testing.T) {
	if !oneEditAway("pale", "ple") || oneEditAway("pales", "ple") {
		t.Error(`TestOneEditAway failed`)
	}
}

func TestCompressString(t *testing.T) {
	if compressString("aabcccccaaa") != "a2b1c5a3" || compressString("abbccd") != "abbccd" {
		t.Error(`TestcompressString failed`)
	}
}

func TestSetZeroes(t *testing.T) {
	if fmt.Sprint(setZeroes([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})) != "[[1 0 1] [0 0 0] [1 0 1]]" || fmt.Sprint(setZeroes([][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}})) != "[[0 0 0 0] [0 4 5 0] [0 3 1 0]]" {
		t.Error(`TestSetZeroes failed`)
	}
}

func TestIsFlippedString(t *testing.T) {
	if !isFlipedString("erbottlewat", "bottlewater") || isFlipedString("aa", "aba") {
		t.Error(`TestIsFlippedString`)
	}
}

func TestRotate(t *testing.T) {
	if fmt.Sprint(rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) != "[[7 4 1] [8 5 2] [9 6 3]]" || fmt.Sprint(rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})) != "[[15 13 2 5] [14 3 4 1] [12 6 8 9] [16 7 10 11]]" {
		t.Error(`TestRotate failed`)
	}
}
