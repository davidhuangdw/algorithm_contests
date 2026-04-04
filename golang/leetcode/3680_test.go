package main

import (
	"reflect"
	"testing"
)

func isValidSchedule(n int, schedule [][]int) bool {
	if n <= 4 {
		return schedule == nil
	}
	if len(schedule) != n*(n-1) {
		return false
	}

	seen := make(map[[2]int]bool)
	for i, match := range schedule {
		if len(match) != 2 {
			return false
		}
		u, v := match[0], match[1]
		if u < 0 || u >= n || v < 0 || v >= n || u == v {
			return false
		}
		if seen[[2]int{u, v}] {
			return false
		}
		seen[[2]int{u, v}] = true

		if i > 0 {
			prev := schedule[i-1]
			if u == prev[0] || u == prev[1] || v == prev[0] || v == prev[1] {
				return false
			}
		}
	}
	return len(seen) == n*(n-1)
}

func Test_generateSchedule(t *testing.T) {
	tests := []struct {
		n    int
		want bool // whether it should be a valid schedule or nil
	}{
		{2, false},
		{3, false},
		{4, false},
		{5, true},
		{6, true},
		{10, true},
	}
	for _, tt := range tests {
		t.Run(reflect.TypeOf(tt.n).Name(), func(t *testing.T) {
			got := generateSchedule(tt.n)
			if tt.want {
				if !isValidSchedule(tt.n, got) {
					t.Errorf("generateSchedule(%d) returned an invalid schedule", tt.n)
				}
			} else {
				if got != nil {
					t.Errorf("generateSchedule(%d) = %v, want nil", tt.n, got)
				}
			}
		})
	}
}
