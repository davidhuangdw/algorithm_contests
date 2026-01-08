package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoneGameIII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   string
	}{
		{
			name:   "Alice wins with simple 3 stones",
			stones: []int{1, 2, 3},
			want:   "Alice",
		},
		{
			name:   "Bob wins with negative stones",
			stones: []int{-1, -2, -3},
			want:   "Tie",
		},
		{
			name:   "Tie game",
			stones: []int{1, -1, 1, -1},
			want:   "Alice",
		},
		{
			name:   "Single stone",
			stones: []int{100},
			want:   "Alice",
		},
		{
			name:   "Two stones",
			stones: []int{10, 20},
			want:   "Alice",
		},
		{
			name:   "Zero stones",
			stones: []int{},
			want:   "Tie",
		},
		{
			name:   "Alice chooses best path",
			stones: []int{1, 2, 3, 7},
			want:   "Bob",
		},
		{
			name:   "Complex Alice win",
			stones: []int{1, 2, 3, 6},
			want:   "Tie",
		},
		{
			name:   "Large numbers",
			stones: []int{10000, 20000, 30000, 40000},
			want:   "Alice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := stoneGameIII(tt.stones)
			assert.Equal(t, tt.want, got, "stoneGameIII(%v) = %s, want %s", tt.stones, got, tt.want)
		})
	}
}
