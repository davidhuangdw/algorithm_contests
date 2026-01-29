package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrimes(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []int
	}{
		{
			name: "n=0",
			n:    0,
			want: []int{},
		},
		{
			name: "n=1",
			n:    1,
			want: []int{},
		},
		{
			name: "n=2",
			n:    2,
			want: []int{2},
		},
		{
			name: "n=3",
			n:    3,
			want: []int{2, 3},
		},
		{
			name: "n=10",
			n:    10,
			want: []int{2, 3, 5, 7},
		},
		{
			name: "n=13",
			n:    13,
			want: []int{2, 3, 5, 7, 11, 13},
		},
		{
			name: "n=20",
			n:    20,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
		{
			name: "n=30",
			n:    30,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
		{
			name: "n=50",
			n:    50,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPrimes(tt.n)
			assert.Equal(t, tt.want, got)
		})
	}
}
