package dp

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "empty coins",
			coins:  []int{},
			amount: 5,
			want:   -1,
		},
		{
			name:   "amount zero",
			coins:  []int{1, 2, 5},
			amount: 0,
			want:   0,
		},
		{
			name:   "single coin exact match",
			coins:  []int{5},
			amount: 5,
			want:   1,
		},
		{
			name:   "single coin no match",
			coins:  []int{3},
			amount: 5,
			want:   -1,
		},
		{
			name:   "classic example",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "impossible amount",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "multiple coins same value",
			coins:  []int{1, 1, 1, 5},
			amount: 7,
			want:   3,
		},
		{
			name:   "large amount with small coins",
			coins:  []int{1, 3, 4},
			amount: 10,
			want:   3,
		},
		{
			name:   "coins with zeros",
			coins:  []int{0, 1, 2},
			amount: 3,
			want:   2,
		},
		{
			name:   "all coins larger than amount",
			coins:  []int{10, 20, 30},
			amount: 5,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
