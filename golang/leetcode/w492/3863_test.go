package main

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "already sorted",
			s:    "abc",
			want: 0,
		},
		{
			name: "impossible case",
			s:    "ba",
			want: -1,
		},
		{
			name: "min at start",
			s:    "acb",
			want: 1,
		},
		{
			name: "max at end",
			s:    "bac",
			want: 1,
		},
		{
			name: "special case",
			s:    "cba",
			want: 3,
		},
		{
			name: "complex case",
			s:    "oool",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.s); got != tt.want {
				t.Errorf("minOperations(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
