package math_geo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		name         string
		columnNumber int
		expected     string
	}{
		{
			name:         "column 1",
			columnNumber: 1,
			expected:     "A",
		},
		{
			name:         "column 2",
			columnNumber: 2,
			expected:     "B",
		},
		{
			name:         "column 26",
			columnNumber: 26,
			expected:     "Z",
		},
		{
			name:         "column 27",
			columnNumber: 27,
			expected:     "AA",
		},
		{
			name:         "column 28",
			columnNumber: 28,
			expected:     "AB",
		},
		{
			name:         "column 52",
			columnNumber: 52,
			expected:     "AZ",
		},
		{
			name:         "column 53",
			columnNumber: 53,
			expected:     "BA",
		},
		{
			name:         "column 701",
			columnNumber: 701,
			expected:     "ZY",
		},
		{
			name:         "column 702",
			columnNumber: 702,
			expected:     "ZZ",
		},
		{
			name:         "column 703",
			columnNumber: 703,
			expected:     "AAA",
		},
		{
			name:         "column 18278",
			columnNumber: 18278,
			expected:     "ZZZ",
		},
		{
			name:         "column 18279",
			columnNumber: 18279,
			expected:     "AAAA",
		},
		{
			name:         "column 1000",
			columnNumber: 1000,
			expected:     "ALL",
		},
		{
			name:         "column 500",
			columnNumber: 500,
			expected:     "SF",
		},
		{
			name:         "column 10000",
			columnNumber: 10000,
			expected:     "NTP",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToTitle(tt.columnNumber)
			assert.Equal(t, tt.expected, result, "convertToTitle(%d)", tt.columnNumber)
		})
	}
}
