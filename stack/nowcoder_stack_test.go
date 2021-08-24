package stack

import (
	"testing"
)

func TestNcStackQueuePush(t *testing.T) {
	QueuePush(1)
	QueuePush(2)
	QueuePush(3)
	if len(stack1) != 3 || stack1[0] != 1 || stack1[1] != 2 || stack1[2] != 3 {
		t.Error(`TestNcStackQueuePush failed`)
	}
}

func TestNcStackQueuePop(t *testing.T) {
	QueuePush(1)
	QueuePush(2)
	QueuePush(3)
	pop1 := QueuePop()
	pop2 := QueuePop()
	pop3 := QueuePop()
	pop4 := QueuePop()
	if pop1 != 3 || pop2 != 2 || pop3 != 1 || pop4 != -1 {
		t.Error(`TestNcStackQueuePop failed`, pop1, pop2, pop3, pop4)
	}
}

func TestNcStackIsValid(t *testing.T) {
	if isValid("[") {
		t.Error(`TestNcStackIsValid failed`, isValid("["))
	}

	if !isValid("[]") {
		t.Error(`TestNcStackIsValid failed`, isValid("[]"))
	}

	if !isValid("()[]{}") {
		t.Error(`TestNcStackIsValid failed`, isValid("()[]{}"))
	}

	if isValid("([])") {
		t.Error(`TestNcStackIsValid failed`, isValid("([])"))
	}
}
