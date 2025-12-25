package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	tests := []struct {
		name     string
		capacity int
		ops      []struct {
			method string
			key    int
			value  int
			expect int
		}
	}{
		{
			name:     "Basic put and get",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{
				{"put", 1, 1, -1},
				{"get", 1, -1, 1},
				{"put", 2, 2, -1},
				{"get", 1, -1, 1},
				{"get", 2, -1, 2},
			},
		},
		{
			name:     "Capacity eviction",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{
				{"put", 1, 1, -1},
				{"put", 2, 2, -1},
				{"put", 3, 3, -1}, // Should evict key 1
				{"get", 1, -1, -1}, // Key 1 should be evicted
				{"get", 2, -1, 2},
				{"get", 3, -1, 3},
			},
		},
		{
			name:     "LRU order maintenance",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{
				{"put", 1, 1, -1},
				{"put", 2, 2, -1},
				{"get", 1, -1, 1}, // Access key 1, making it most recently used
				{"put", 3, 3, -1}, // Should evict key 2 (least recently used)
				{"get", 2, -1, -1}, // Key 2 should be evicted
				{"get", 1, -1, 1},
				{"get", 3, -1, 3},
			},
		},
		{
			name:     "Update existing key",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{
				{"put", 1, 1, -1},
				{"put", 2, 2, -1},
				{"put", 1, 10, -1}, // Update key 1
				{"get", 1, -1, 10}, // Should get updated value
				{"put", 3, 3, -1},  // Should evict key 2
				{"get", 2, -1, -1},
				{"get", 1, -1, 10},
			},
		},
		{
			name:     "Single capacity",
			capacity: 1,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{
				{"put", 1, 1, -1},
				{"get", 1, -1, 1},
				{"put", 2, 2, -1}, // Should evict key 1
				{"get", 1, -1, -1},
				{"get", 2, -1, 2},
			},
		},
		{
			name:     "Large capacity",
			capacity: 3,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{
				{"put", 1, 1, -1},
				{"put", 2, 2, -1},
				{"put", 3, 3, -1},
				{"get", 1, -1, 1},
				{"get", 2, -1, 2},
				{"get", 3, -1, 3},
				{"put", 4, 4, -1}, // Should evict key 1 (least recently used)
				{"get", 1, -1, -1},
				{"get", 4, -1, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := Constructor(tt.capacity)
			
			for _, op := range tt.ops {
				switch op.method {
				case "put":
					cache.Put(op.key, op.value)
				case "get":
					result := cache.Get(op.key)
					assert.Equal(t, op.expect, result, 
						"Get(%d) should return %d, got %d", op.key, op.expect, result)
				}
			}
		})
	}
}