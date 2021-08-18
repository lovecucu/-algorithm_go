package stack

import (
	"testing"
)

func TestIsValidBrackets(t *testing.T) {
	if !isValidBrackets("()") {
		t.Error(`TestIsValidBrackets failed`)
	}

	if !isValidBrackets("()[]{}") {
		t.Error(`TestIsValidBrackets failed`)
	}

	if isValidBrackets("(]") {
		t.Error(`TestIsValidBrackets failed`)
	}

	if isValidBrackets("([)]") {
		t.Error(`TestIsValidBrackets failed`)
	}

	if !isValidBrackets("{[]}") {
		t.Error(`TestIsValidBrackets failed`)
	}
}
