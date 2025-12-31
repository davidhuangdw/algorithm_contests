package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{
			name: "leetcode example 1",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "leetcode example 2",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			name: "single station sufficient",
			gas:  []int{5},
			cost: []int{3},
			want: 0,
		},
		{
			name: "single station insufficient",
			gas:  []int{2},
			cost: []int{3},
			want: -1,
		},
		{
			name: "start from index 0",
			gas:  []int{3, 1, 1},
			cost: []int{1, 2, 2},
			want: 0,
		},
		{
			name: "start from middle",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "exact balance",
			gas:  []int{2, 2, 2},
			cost: []int{2, 2, 2},
			want: 0,
		},
		{
			name: "large deficit at start",
			gas:  []int{1, 10, 1},
			cost: []int{10, 1, 1},
			want: 1,
		},
		{
			name: "circular path possible",
			gas:  []int{5, 1, 2, 3, 4},
			cost: []int{4, 4, 1, 5, 1},
			want: 4,
		},
		{
			name: "all stations insufficient",
			gas:  []int{1, 1, 1},
			cost: []int{2, 2, 2},
			want: -1,
		},
		{
			name: "alternating pattern",
			gas:  []int{4, 1, 4, 1},
			cost: []int{1, 4, 1, 4},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canCompleteCircuit(tt.gas, tt.cost)
			assert.Equal(t, tt.want, got, "canCompleteCircuit(%v, %v) = %d, want %d", tt.gas, tt.cost, got, tt.want)
		})
	}
}