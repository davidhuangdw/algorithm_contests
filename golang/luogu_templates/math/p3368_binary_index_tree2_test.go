package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryIndexTree2(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		init    []int
		actions []struct {
			typeAction string // "add" or "get"
			l          int    // l for add, index for get
			r          int    // r for add, ignored for get
			v          int    // value for add, ignored for get
			want       int    // expected get result, ignored for add
		}
	}{
		{
			name: "empty initialization",
			n:    5,
			init: []int{},
			actions: []struct {
				typeAction string
				l          int
				r          int
				v          int
				want       int
			}{
				{"add", 1, 3, 10, 0},
				{"add", 2, 5, 20, 0},
				{"get", 1, 0, 0, 10},
				{"get", 2, 0, 0, 30}, // 10 + 20
				{"get", 3, 0, 0, 30}, // 10 + 20
				{"get", 4, 0, 0, 20},
				{"get", 5, 0, 0, 20},
				{"add", 3, 3, -15, 0},
				{"get", 3, 0, 0, 15}, // 10 + 20 -15
			},
		},
		{
			name: "initial array",
			n:    4,
			init: []int{1, 2, 3, 4},
			actions: []struct {
				typeAction string
				l          int
				r          int
				v          int
				want       int
			}{
				{"get", 1, 0, 0, 1},
				{"get", 4, 0, 0, 4},
				{"add", 1, 2, 5, 0},
				{"get", 1, 0, 0, 6}, // 1 + 5
				{"get", 2, 0, 0, 7}, // 2 + 5
				{"get", 3, 0, 0, 3},
				{"add", 2, 4, 3, 0},
				{"get", 2, 0, 0, 10}, // 2 + 5 + 3
				{"get", 3, 0, 0, 6},  // 3 + 3
				{"get", 4, 0, 0, 7},  // 4 + 3
			},
		},
		{
			name: "single element",
			n:    1,
			init: []int{10},
			actions: []struct {
				typeAction string
				l          int
				r          int
				v          int
				want       int
			}{
				{"get", 1, 0, 0, 10},
				{"add", 1, 1, -5, 0},
				{"get", 1, 0, 0, 5},
				{"add", 1, 1, 7, 0},
				{"get", 1, 0, 0, 12},
			},
		},
		{
			name: "multiple overlapping ranges",
			n:    5,
			init: []int{0, 0, 0, 0, 0},
			actions: []struct {
				typeAction string
				l          int
				r          int
				v          int
				want       int
			}{
				{"add", 1, 5, 2, 0},
				{"add", 2, 4, 3, 0},
				{"add", 3, 3, 5, 0},
				{"get", 1, 0, 0, 2},  // 2
				{"get", 2, 0, 0, 5},  // 2 + 3
				{"get", 3, 0, 0, 10}, // 2 + 3 + 5
				{"get", 4, 0, 0, 5},  // 2 + 3
				{"get", 5, 0, 0, 2},  // 2
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bit := NewBinaryIndexTree2(tt.n, tt.init)

			for _, action := range tt.actions {
				switch action.typeAction {
				case "add":
					bit.Add(action.l, action.r, action.v)
				case "get":
					got := bit.Get(action.l)
					assert.Equal(t, action.want, got, "get at index %d should be %d, got %d", action.l, action.want, got)
				}
			}
		})
	}
}
