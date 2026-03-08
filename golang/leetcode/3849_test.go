package main

import "testing"

func Test_maximumXor(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "simple case all 0s and 1s",
			s:    "10",
			t:    "01",
			want: "11",
		},
		{
			name: "maximize with available bits",
			s:    "1010",
			t:    "1100",
			want: "1111",
		},
		{
			name: "all same bits",
			s:    "1111",
			t:    "1111",
			want: "0000",
		},
		{
			name: "single bit each",
			s:    "1",
			t:    "0",
			want: "1",
		},
		{
			name: "single bit same",
			s:    "0",
			t:    "0",
			want: "0",
		},
		{
			name: "longer string s",
			s:    "11001",
			t:    "001",
			want: "11100",
		},
		{
			name: "greedy selection",
			s:    "000",
			t:    "111",
			want: "111",
		},
		{
			name: "mixed availability",
			s:    "101010",
			t:    "001111",
			want: "111101",
		},
		{
			name: "t has more 0s",
			s:    "111",
			t:    "0000",
			want: "111",
		},
		{
			name: "all zeros",
			s:    "0000",
			t:    "0000",
			want: "0000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumXor(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("maximumXor(%q, %q) = %q, want %q", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
