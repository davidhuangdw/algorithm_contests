package main

import "testing"

func Test_countArrays(t *testing.T) {
	tests := []struct {
		name     string
		digitSum []int
		want     int
	}{
		{"Test Case 1", []int{2}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countArrays(tt.digitSum); got != tt.want {
				t.Errorf("countArrays() = %v, want %v", got, tt.want) // this will fail, will adjust
			}
		})
	}
}
