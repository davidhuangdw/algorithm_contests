package main

import (
	"testing"
)

func TestCountSteppingNumbers(t *testing.T) {
	tests := []struct {
		name string
		low  string
		high string
		want int
	}{
		{
			name: "Example 1",
			low:  "1",
			high: "11",
			want: 10,
		},
		{
			name: "Example 2",
			low:  "90",
			high: "101",
			want: 2,
		},
		{
			name: "Single digit test",
			low:  "1",
			high: "9",
			want: 9,
		},
		{
			name: "Zero included",
			low:  "0",
			high: "9",
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSteppingNumbers(tt.low, tt.high); got != tt.want {
				t.Errorf("countSteppingNumbers(%q, %q) = %v, want %v", tt.low, tt.high, got, tt.want)
			}
		})
	}
}
