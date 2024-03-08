package tree

import (
	"fmt"
	"testing"
)

func TestNewHeapMax(t *testing.T) {
	maxHeap := NewMaxHeap([]any{1, 6, 7, 4, 3, 9, 10, 5})
	fmt.Printf("Max heap:%v", maxHeap.data)
}
