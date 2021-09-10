package heap

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	input := []int{3, 2, 4, 1, 6, 5, 1}
	target := "[1 1 2 3 4 5 6]"
	HeapSort(input)
	if fmt.Sprint(input) != target {
		t.Error(`TestHeapSort failed`, input)
	}
}
