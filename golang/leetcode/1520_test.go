package main

import (
	"reflect"
	"testing"
)

func TestMaxNumOfSubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"adefaddaccc", []string{"e", "f", "ccc"}},
		{"abbaccd", []string{"bb", "cc", "d"}},
	}
	for i, tt := range tests {
		got := maxNumOfSubstrings(tt.s)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("case %d: maxNumOfSubstrings(%v) = %v; want %v", i, tt.s, got, tt.want)
		}
	}
}
