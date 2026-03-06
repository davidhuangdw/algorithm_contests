package main

import "testing"

func TestFhqTreap_RepeatInsert(t *testing.T) {
	tests := []struct {
		name string
		ops  []struct {
			op   string
			val  int
			want int
		}
	}{
		{
			name: "insert duplicates and check size",
			ops: []struct {
				op   string
				val  int
				want int
			}{
				{"insert", 5, 0},
				{"insert", 5, 0},
				{"insert", 5, 0},
				{"size", 0, 3}, // should be 3, not 1
				{"rank", 5, 1}, // rank of 5 is 1 (no elements < 5)
				{"rank", 6, 4}, // rank of 6 is 4 (3 elements of value 5 < 6)
			},
		},
		{
			name: "insert duplicates with different values",
			ops: []struct {
				op   string
				val  int
				want int
			}{
				{"insert", 3, 0},
				{"insert", 3, 0},
				{"insert", 5, 0},
				{"insert", 5, 0},
				{"insert", 7, 0},
				{"size", 0, 5},
				{"rank", 4, 3}, // 2 elements of value 3 < 4
				{"rank", 6, 5}, // 2+2 elements < 6
				{"kth", 1, 3},  // 1st smallest is 3
				{"kth", 3, 5},  // 3rd smallest is 5
				{"kth", 5, 7},  // 5th smallest is 7
			},
		},
		{
			name: "delete duplicates",
			ops: []struct {
				op   string
				val  int
				want int
			}{
				{"insert", 5, 0},
				{"insert", 5, 0},
				{"insert", 5, 0},
				{"size", 0, 3},
				{"del", 5, 0},
				{"size", 0, 2},
				{"del", 5, 0},
				{"size", 0, 1},
				{"del", 5, 0},
				{"size", 0, 0},
			},
		},
		{
			name: "prev and succ with duplicates",
			ops: []struct {
				op   string
				val  int
				want int
			}{
				{"insert", 5, 0},
				{"insert", 5, 0},
				{"insert", 10, 0},
				{"insert", 10, 0},
				{"prev", 6, 5},   // predecessor of 6 is 5
				{"succ", 6, 10},  // successor of 6 is 10
				{"prev", 11, 10}, // predecessor of 11 is 10
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &FhqTreap{}
			for i, op := range tt.ops {
				switch op.op {
				case "insert":
					tr.Insert(op.val)
				case "del":
					tr.Del(op.val)
				case "size":
					got := tr.root.Size()
					if got != op.want {
						t.Errorf("step %d: size = %v, want %v", i, got, op.want)
					}
				case "rank":
					got := tr.Rank(op.val)
					if got != op.want {
						t.Errorf("step %d: rank(%v) = %v, want %v", i, op.val, got, op.want)
					}
				case "kth":
					got := tr.SubKth(tr.root, op.val)
					if got != op.want {
						t.Errorf("step %d: kth(%v) = %v, want %v", i, op.val, got, op.want)
					}
				case "prev":
					got := tr.Prev(op.val)
					if got != op.want {
						t.Errorf("step %d: prev(%v) = %v, want %v", i, op.val, got, op.want)
					}
				case "succ":
					got := tr.Succ(op.val)
					if got != op.want {
						t.Errorf("step %d: succ(%v) = %v, want %v", i, op.val, got, op.want)
					}
				}
			}
		})
	}
}
