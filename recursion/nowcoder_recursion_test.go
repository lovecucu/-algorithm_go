package recursion

import "testing"

func TestNcRecursionFibonacci(t *testing.T) {
	if Fibonacci(0) != 0 || Fibonacci(1) != 1 || Fibonacci(2) != 1 {
		t.Error(`TestNcRecursionFibonacci failed`)
	}

	if Fibonacci(6) != 8 {
		t.Error(`TestNcRecursionFibonacci failed`)
	}
}

func TestNcRecursionJumpFloor(t *testing.T) {
	if jumpFloor(1) != 1 || jumpFloor(2) != 2 || jumpFloor(3) != 3 {
		t.Error(`TestNcRecursionJumpFloor failed`)
	}

	if jumpFloor(5) != 8 {
		t.Error(`TestNcRecursionJumpFloor failed`)
	}
}
