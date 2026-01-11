package greedy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictPartyVictory(t *testing.T) {
	tests := []struct {
		name     string
		senate   string
		expected string
	}{
		{
			name:     "LeetCode example 1 - RD",
			senate:   "RD",
			expected: "Radiant",
		},
		{
			name:     "LeetCode example 2 - RDD",
			senate:   "RDD",
			expected: "Dire",
		},
		{
			name:     "single Radiant",
			senate:   "R",
			expected: "Radiant",
		},
		{
			name:     "single Dire",
			senate:   "D",
			expected: "Dire",
		},
		{
			name:     "all Radiant",
			senate:   "RRR",
			expected: "Radiant",
		},
		{
			name:     "all Dire",
			senate:   "DDD",
			expected: "Dire",
		},
		{
			name:     "Radiant wins with alternating pattern",
			senate:   "RDRDRD",
			expected: "Radiant",
		},
		{
			name:     "Dire wins with more members",
			senate:   "DDRRR",
			expected: "Dire",
		},
		{
			name:     "complex scenario 1",
			senate:   "RDRD",
			expected: "Radiant",
		},
		{
			name:     "complex scenario 2",
			senate:   "DRRD",
			expected: "Dire",
		},
		{
			name:     "long sequence Radiant win",
			senate:   "RRDDRD",
			expected: "Radiant",
		},
		{
			name:     "long sequence Dire win",
			senate:   "DDRRDR",
			expected: "Dire",
		},
		{
			name:     "balanced but Radiant wins",
			senate:   "RDR",
			expected: "Radiant",
		},
		{
			name:     "complex+1",
			senate:   "DRRDRDRDRDDRDRDR",
			expected: "Radiant",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := predictPartyVictory(tt.senate)
			assert.Equal(t, tt.expected, result, "predictPartyVictory(%q)", tt.senate)
		})
	}
}
