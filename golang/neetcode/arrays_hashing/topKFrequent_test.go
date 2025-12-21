package arrays_hashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 2, 2, 3, 3, 3}, 2, []int{2, 3}},
		{[]int{7, 7}, 1, []int{7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums=%v k=%d", tt.nums, tt.k), func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
