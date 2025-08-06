package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_2_Put(t *testing.T) {
	type fields struct {
		Cap  int
		keyM map[int]*LinkNode
		head *LinkNode
		tail *LinkNode
	}
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := ConstructorLRU_2(2)
			cache.Put(1, 1)
			cache.Put(2, 2)
			assert.Equal(t, 1, cache.Get(1))  // 返回  1
			cache.Put(3, 3)                   // 该操作会使得密钥 2 作废
			assert.Equal(t, -1, cache.Get(2)) // 返回 -1 (未找到)
			cache.Put(4, 4)                   // 该操作会使得密钥 1 作废
			assert.Equal(t, -1, cache.Get(1)) // 返回 -1 (未找到)
			assert.Equal(t, 3, cache.Get(3))  // 返回  3
			assert.Equal(t, 4, cache.Get(4))  // 返回  4
		})
	}
}
