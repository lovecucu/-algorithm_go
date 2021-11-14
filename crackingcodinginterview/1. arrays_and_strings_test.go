package crackingcodinginterview

import "testing"

func TestIsUnique(t *testing.T) {
	if !isUnique("abc") || isUnique("aabc") || isUnique("leetcode") {
		t.Error(`TestIsUnique is failed`)
	}
}

func TestCheckPermutation(t *testing.T) {
	if !CheckPermutation("abc", "bca") || CheckPermutation("abc", "bad") || CheckPermutation("abb", "aab") {
		t.Error(`TestCheckPermutation is failed`)
	}
}
