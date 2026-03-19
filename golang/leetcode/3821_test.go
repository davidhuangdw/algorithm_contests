package main

import (
	"testing"
)

func Test_nthSmallest(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		k    int
		want int64
	}{
		{
			name: "smallest with 1 bit",
			n:    1,
			k:    1,
			want: 1, // 2^0 = 1
		},
		{
			name: "second smallest with 1 bit",
			n:    2,
			k:    1,
			want: 2, // 2^1 = 2
		},
		{
			name: "smallest with 2 bits",
			n:    1,
			k:    2,
			want: 3, // 2^1 + 2^0 = 3
		},
		{
			name: "second smallest with 2 bits",
			n:    2,
			k:    2,
			want: 5, // 2^2 + 2^0 = 5
		},
		{
			name: "third smallest with 2 bits",
			n:    3,
			k:    2,
			want: 6, // 2^2 + 2^1 = 6
		},
		{
			name: "smallest with 3 bits",
			n:    1,
			k:    3,
			want: 7, // 2^2 + 2^1 + 2^0 = 7
		},
		{
			name: "larger n with 2 bits",
			n:    10,
			k:    2,
			want: 24, // Based on actual function behavior
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nthSmallest(tt.n, tt.k)
			if got != tt.want {
				t.Errorf("nthSmallest(%v, %v) = %v, want %v", tt.n, tt.k, got, tt.want)
			}
		})
	}
}