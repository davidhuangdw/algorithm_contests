package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryIndexTree1(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		init    []int
		actions []struct {
			typeAction string // "add" or "sum"
			x          int    // index for add, l for sum
			y          int    // value for add, r for sum
			want       int    // expected sum result, ignored for add
		}
	}{
		{
			name: "empty initialization",
			n:    5,
			init: []int{},
			actions: []struct {
				typeAction string
				x          int
				y          int
				want       int
			}{
				{"add", 1, 10, 0},
				{"add", 3, 20, 0},
				{"add", 5, 30, 0},
				{"sum", 1, 5, 60},
				{"sum", 2, 4, 20},
				{"sum", 3, 3, 20},
			},
		},
		{
			name: "initial array",
			n:    4,
			init: []int{1, 2, 3, 4},
			actions: []struct {
				typeAction string
				x          int
				y          int
				want       int
			}{
				{"sum", 1, 4, 10},
				{"sum", 2, 3, 5},
				{"add", 2, 5, 0},
				{"sum", 1, 4, 15},
				{"sum", 2, 3, 10},
			},
		},
		{
			name: "single element",
			n:    1,
			init: []int{5},
			actions: []struct {
				typeAction string
				x          int
				y          int
				want       int
			}{
				{"sum", 1, 1, 5},
				{"add", 1, -3, 0},
				{"sum", 1, 1, 2},
			},
		},
		{
			name: "large range operations",
			n:    10,
			init: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			actions: []struct {
				typeAction string
				x          int
				y          int
				want       int
			}{
				{"sum", 1, 10, 55},
				{"add", 5, 10, 0},
				{"sum", 1, 10, 65},
				{"sum", 3, 7, 35},
				{"add", 1, 5, 0},
				{"sum", 1, 3, 11},
				{"sum", 1, 3, 11},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bit := NewBinaryIndexTree1(tt.n, tt.init)

			for _, action := range tt.actions {
				switch action.typeAction {
				case "add":
					bit.Add(action.x, action.y)
				case "sum":
					got := bit.Sum(action.x, action.y)
					assert.Equal(t, action.want, got, "sum from %d to %d should be %d, got %d", action.x, action.y, action.want, got)
				}
			}
		})
	}
}
