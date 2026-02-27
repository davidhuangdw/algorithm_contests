package main

import (
	"testing"
)

func Test_diffConstraints(t *testing.T) {
	tests := []struct {
		name        string
		n           int
		constraints [][3]int
		wantNil     bool
	}{
		{
			name: "valid system",
			n:    3,
			constraints: [][3]int{
				{1, 2, 3},
				{2, 3, -1},
			},
			wantNil: false,
		},
		{
			name: "negative cycle",
			n:    3,
			constraints: [][3]int{
				{1, 2, -1},
				{2, 3, -1},
				{3, 1, -1},
			},
			wantNil: true,
		},
		{
			name:        "no constraints",
			n:           2,
			constraints: [][3]int{},
			wantNil:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffConstraints(tt.n, tt.constraints)
			if (got == nil) != tt.wantNil {
				t.Errorf("diffConstraints() = %v, wantNil %v", got, tt.wantNil)
			}
			if !tt.wantNil && len(got) != tt.n {
				t.Errorf("diffConstraints() length = %v, want %v", len(got), tt.n)
			}
			if !tt.wantNil && got != nil {
				// verify constraints are satisfied
				for _, c := range tt.constraints {
					u, v, y := c[0], c[1], c[2]
					if got[u-1]-got[v-1] > y {
						t.Errorf("constraint x[%d]-x[%d]<=%d violated: %d-%d=%d > %d",
							u, v, y, got[u-1], got[v-1], got[u-1]-got[v-1], y)
					}
				}
			}
		})
	}
}
