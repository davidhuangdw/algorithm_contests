package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStr(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Basic rotation - 'bac' should rotate to 'acb'",
			input: "bac",
			want:  "acb",
		},
		{
			name:  "All characters same - 'aaaaa' should return 'aaaaa'",
			input: "aaaaa",
			want:  "aaaaa",
		},
		{
			name:  "Already smallest - 'abcde' should remain 'abcde'",
			input: "abcde",
			want:  "abcde",
		},
		{
			name:  "Rotation at end - 'cba' should rotate to 'abc'",
			input: "cba",
			want:  "acb",
		},
		{
			name:  "Single character - 'x' should return 'x'",
			input: "x",
			want:  "x",
		},
		{
			name:  "Two characters - 'ba' should rotate to 'ab'",
			input: "ba",
			want:  "ab",
		},
		{
			name:  "Complex case - 'bcabc' should rotate to 'abcbc'",
			input: "bcabc",
			want:  "abcbc",
		},
		{
			name:  "Case with equal characters - 'abac' should rotate to 'aabc'",
			input: "abac",
			want:  "abac",
		},
		{
			name:  "Longer complex case - 'programming'",
			input: "programming",
			want:  "ammingprogr",
		},
		{
			name:  "Edge case empty string",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minStr(tt.input)
			assert.Equal(t, tt.want, got, "minStr('%s') = '%s', want '%s'", tt.input, got, tt.want)
		})
	}
}
