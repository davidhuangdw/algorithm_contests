package main

import "testing"

func Test_longestBalanced(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "010101",
			want: 6,
		},
		{
			name: "Example 2",
			s:    "000111",
			want: 6,
		},
		{
			name: "Example 3",
			s:    "11100",
			want: 4, // "1100" after swap? or "1100" is sum 0? 1100 has sum 0.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestBalanced(tt.s); got != tt.want {
				t.Errorf("longestBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
