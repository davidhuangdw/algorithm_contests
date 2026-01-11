package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{
			name:      "empty equations and queries",
			equations: [][]string{},
			values:    []float64{},
			queries:   [][]string{},
			expected:  []float64{},
		},
		{
			name:      "single equation - direct division",
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}},
			expected:  []float64{0.5, 2.0},
		},
		{
			name:      "LeetCode example 1",
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			expected:  []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name:      "LeetCode example 2",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 0.2},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			expected:  []float64{3.75, 0.4, 0.2, 5.0},
		},
		{
			name:      "disconnected graph",
			equations: [][]string{{"a", "b"}, {"c", "d"}},
			values:    []float64{0.5, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "d"}, {"a", "b"}, {"c", "d"}},
			expected:  []float64{-1.0, -1.0, 0.5, 3.0},
		},
		{
			name:      "self-referential and unknown variables",
			equations: [][]string{{"x", "y"}},
			values:    []float64{2.0},
			queries:   [][]string{{"x", "x"}, {"y", "y"}, {"z", "z"}, {"x", "z"}},
			expected:  []float64{1.0, 1.0, -1.0, -1.0},
		},
		{
			name:      "complex chain with multiple paths",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}},
			values:    []float64{2.0, 3.0, 4.0, 8.0},
			queries:   [][]string{{"a", "d"}, {"d", "a"}, {"d", "b"}, {"b", "d"}},
			expected:  []float64{24, 0.041666666666666664, 0.08333333333333333, 12.0},
		},
		{
			name:      "single node graph",
			equations: [][]string{{"a", "a"}},
			values:    []float64{1.0},
			queries:   [][]string{{"a", "a"}},
			expected:  []float64{1.0},
		},
		{
			name:      "multiple equations with same variables",
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"a", "c"}},
			values:    []float64{2.0, 3.0, 6.0},
			queries:   [][]string{{"c", "a"}, {"a", "c"}, {"c", "b"}},
			expected:  []float64{0.16666666666666666, 6.0, 0.3333333333333333},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calcEquation(tt.equations, tt.values, tt.queries)
			assert.InDeltaSlice(t, tt.expected, result, 0.0001, "calcEquation(%v, %v, %v)", tt.equations, tt.values, tt.queries)
		})
	}
}
