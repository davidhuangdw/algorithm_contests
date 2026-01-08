package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFUCache(t *testing.T) {
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
			name:     "Basic functionality",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{{
				method: "put",
				key:    1,
				value:  1,
				expect: -1,
			}, {
				method: "put",
				key:    2,
				value:  2,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: 1,
			}, {
				method: "put",
				key:    3,
				value:  3,
				expect: -1,
			}, {
				method: "get",
				key:    2,
				expect: -1,
			}, {
				method: "get",
				key:    3,
				expect: 3,
			}},
		},
		{
			name:     "Eviction with same frequency",
			capacity: 3,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{{
				method: "put",
				key:    1,
				value:  1,
				expect: -1,
			}, {
				method: "put",
				key:    2,
				value:  2,
				expect: -1,
			}, {
				method: "put",
				key:    3,
				value:  3,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: 1,
			}, {
				method: "get",
				key:    2,
				expect: 2,
			}, {
				method: "put",
				key:    4,
				value:  4,
				expect: -1,
			}, {
				method: "get",
				key:    3,
				expect: -1,
			}, {
				method: "get",
				key:    2,
				expect: 2,
			}, {
				method: "put",
				key:    5,
				value:  5,
				expect: -1,
			}, {
				method: "get",
				key:    4,
				expect: -1,
			}},
		},
		{
			name:     "Update existing key",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{{
				method: "put",
				key:    1,
				value:  1,
				expect: -1,
			}, {
				method: "put",
				key:    2,
				value:  2,
				expect: -1,
			}, {
				method: "put",
				key:    1,
				value:  10,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: 10,
			}, {
				method: "put",
				key:    3,
				value:  3,
				expect: -1,
			}, {
				method: "get",
				key:    2,
				expect: -1,
			}},
		},
		{
			name:     "Capacity 0",
			capacity: 0,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{{
				method: "put",
				key:    1,
				value:  1,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: -1,
			}},
		},
		{
			name:     "Capacity 1",
			capacity: 1,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{{
				method: "put",
				key:    1,
				value:  1,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: 1,
			}, {
				method: "put",
				key:    2,
				value:  2,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: -1,
			}, {
				method: "get",
				key:    2,
				expect: 2,
			}},
		},
		{
			name:     "Complex scenario from LeetCode",
			capacity: 2,
			ops: []struct {
				method string
				key    int
				value  int
				expect int
			}{{
				method: "put",
				key:    1,
				value:  1,
				expect: -1,
			}, {
				method: "put",
				key:    2,
				value:  2,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: 1,
			}, {
				method: "put",
				key:    3,
				value:  3,
				expect: -1,
			}, {
				method: "get",
				key:    2,
				expect: -1,
			}, {
				method: "get",
				key:    3,
				expect: 3,
			}, {
				method: "put",
				key:    4,
				value:  4,
				expect: -1,
			}, {
				method: "get",
				key:    1,
				expect: -1,
			}, {
				method: "get",
				key:    3,
				expect: 3,
			}, {
				method: "get",
				key:    4,
				expect: 4,
			}},
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := ConstructorLFU(tt.capacity)

			for _, op := range tt.ops {
				switch op.method {
				case "put":
					cache.Put(op.key, op.value)
				case "get":
					result := cache.Get(op.key)
					assert.Equal(t, op.expect, result, "Get(%d) should return %d, got %d", op.key, op.expect, result)
				}
			}
		})
	}
}
