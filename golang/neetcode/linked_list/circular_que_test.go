package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularQueue(t *testing.T) {
	tests := []struct {
		name string
		k    int
		ops  []struct {
			op    string
			val   int
			want  bool
			front int
			rear  int
		}
	}{
		{
			name: "LeetCode Example 1",
			k:    3,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"enqueue", 1, true, 1, 1},
				{"enqueue", 2, true, 1, 2},
				{"enqueue", 3, true, 1, 3},
				{"enqueue", 4, false, 1, 3},
				{"rear", 0, false, 1, 3},
				{"isfull", 0, true, 1, 3},
				{"dequeue", 0, true, 2, 3},
				{"enqueue", 4, true, 2, 4},
				{"rear", 0, false, 2, 4},
			},
		},
		{
			name: "Empty queue operations",
			k:    2,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"isempty", 0, true, -1, -1},
				{"front", 0, false, -1, -1},
				{"rear", 0, false, -1, -1},
				{"dequeue", 0, false, -1, -1},
				{"enqueue", 1, true, 1, 1},
				{"isempty", 0, false, 1, 1},
			},
		},
		{
			name: "Single element queue",
			k:    1,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"enqueue", 5, true, 5, 5},
				{"isfull", 0, true, 5, 5},
				{"enqueue", 6, false, 5, 5},
				{"front", 0, false, 5, 5},
				{"rear", 0, false, 5, 5},
				{"dequeue", 0, true, -1, -1},
				{"dequeue", 0, false, -1, -1},
			},
		},
		{
			name: "Circular behavior",
			k:    3,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"enqueue", 1, true, 1, 1},
				{"enqueue", 2, true, 1, 2},
				{"enqueue", 3, true, 1, 3},
				{"dequeue", 0, true, 2, 3},
				{"enqueue", 4, true, 2, 4},
				{"dequeue", 0, true, 3, 4},
				{"enqueue", 5, true, 3, 5},
				{"dequeue", 0, true, 4, 5},
				{"enqueue", 6, true, 4, 6},
				{"isfull", 0, true, 4, 6},
			},
		},
		{
			name: "Zero capacity",
			k:    0,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"enqueue", 1, false, -1, -1},
				{"dequeue", 0, false, -1, -1},
				{"isempty", 0, true, -1, -1},
				{"isfull", 0, true, -1, -1},
			},
		},
		{
			name: "Large capacity",
			k:    5,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"enqueue", 1, true, 1, 1},
				{"enqueue", 2, true, 1, 2},
				{"enqueue", 3, true, 1, 3},
				{"enqueue", 4, true, 1, 4},
				{"enqueue", 5, true, 1, 5},
				{"isfull", 0, true, 1, 5},
				{"enqueue", 6, false, 1, 5},
				{"dequeue", 0, true, 2, 5},
				{"enqueue", 6, true, 2, 6},
				{"rear", 0, false, 2, 6},
			},
		},
		{
			name: "Mixed operations",
			k:    4,
			ops: []struct {
				op    string
				val   int
				want  bool
				front int
				rear  int
			}{
				{"enqueue", 10, true, 10, 10},
				{"enqueue", 20, true, 10, 20},
				{"dequeue", 0, true, 20, 20},
				{"enqueue", 30, true, 20, 30},
				{"enqueue", 40, true, 20, 40},
				{"enqueue", 50, true, 20, 50},
				{"isfull", 0, true, 20, 50},
				{"dequeue", 0, true, 30, 50},
				{"dequeue", 0, true, 40, 50},
				{"dequeue", 0, true, 50, 50},
				{"isempty", 0, false, 50, 50},
				{"dequeue", 0, true, -1, -1},
				{"isempty", 0, true, -1, -1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := ConstructorCQ(tt.k)

			for i, op := range tt.ops {
				switch op.op {
				case "enqueue":
					result := cq.EnQueue(op.val)
					assert.Equal(t, op.want, result, "EnQueue operation %d failed", i)
				case "dequeue":
					result := cq.DeQueue()
					assert.Equal(t, op.want, result, "DeQueue operation %d failed", i)
				case "front":
					result := cq.Front()
					assert.Equal(t, op.front, result, "Front operation %d failed", i)
				case "rear":
					result := cq.Rear()
					assert.Equal(t, op.rear, result, "Rear operation %d failed", i)
				case "isempty":
					result := cq.IsEmpty()
					assert.Equal(t, op.want, result, "IsEmpty operation %d failed", i)
				case "isfull":
					result := cq.IsFull()
					assert.Equal(t, op.want, result, "IsFull operation %d failed", i)
				}
			}
		})
	}
}
