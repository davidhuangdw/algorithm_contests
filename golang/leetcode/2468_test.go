package main

import (
	"reflect"
	"testing"
)

func TestSplitMessage(t *testing.T) {
	tests := []struct {
		name    string
		message string
		limit   int
		want    []string
	}{
		{
			name:    "Example 1",
			message: "this is really a very awesome message",
			limit:   9,
			want: []string{
				"thi<1/14>", "s i<2/14>", "s r<3/14>", "eal<4/14>", "ly <5/14>",
				"a v<6/14>", "ery<7/14>", " aw<8/14>", "eso<9/14>", "me<10/14>",
				" m<11/14>", "es<12/14>", "sa<13/14>", "ge<14/14>",
			},
		},
		{
			name:    "Example 2",
			message: "short message",
			limit:   15,
			want:    []string{"short mess<1/2>", "age<2/2>"},
		},
		{
			name:    "Impossible case - limit too small",
			message: "abc",
			limit:   3,
			want:    nil,
		},
		{
			name:    "Single character, large limit",
			message: "a",
			limit:   10,
			want:    []string{"a<1/1>"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitMessage(tt.message, tt.limit)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
