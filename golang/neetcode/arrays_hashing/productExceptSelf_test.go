package arrays_hashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{24, 12, 8, 6},
		},
		{
			[]int{-1, 1, 0, -3, 3},
			[]int{0, 0, 9, 0, 0},
		},
		{
			[]int{2, 3},
			[]int{3, 2},
		},
		{
			[]int{1, 1, 1, 1},
			[]int{1, 1, 1, 1},
		},
		{
			[]int{0, 0},
			[]int{0, 0},
		},
		{
			[]int{5},
			[]int{1},
		},
		{
			[]int{2, 0, 4},
			[]int{0, 8, 0},
		},
		{
			[]int{-2, -3, -4},
			[]int{12, 8, 6},
		},
		{
			[]int{10, 20, 30},
			[]int{600, 300, 200},
		},
		{
			[]int{1, 0, 2, 0, 3},
			[]int{0, 0, 0, 0, 0},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums=%v", tt.nums), func(t *testing.T) {
			got := productExceptSelf(tt.nums)
			assert.Equal(t, tt.want, got, "productExceptSelf(%v) should equal %v", tt.nums, tt.want)
		})
	}
}
