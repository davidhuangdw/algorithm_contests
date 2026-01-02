package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindItinerary(t *testing.T) {
	tests := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		{
			name: "LeetCode example 1 - simple itinerary",
			tickets: [][]string{
				{"MUC", "LHR"},
				{"JFK", "MUC"},
				{"SFO", "SJC"},
				{"LHR", "SFO"},
			},
			expected: []string{"JFK", "MUC", "LHR", "SFO", "SJC"},
		},
		{
			name: "LeetCode example 2 - multiple choices",
			tickets: [][]string{
				{"JFK", "SFO"},
				{"JFK", "ATL"},
				{"SFO", "ATL"},
				{"ATL", "JFK"},
				{"ATL", "SFO"},
			},
			expected: []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"},
		},
		{
			name: "Single ticket",
			tickets: [][]string{
				{"JFK", "LAX"},
			},
			expected: []string{"JFK", "LAX"},
		},
		{
			name: "Multiple tickets from same origin",
			tickets: [][]string{
				{"JFK", "LAX"},
				{"JFK", "SFO"},
				{"LAX", "JFK"},
				{"SFO", "JFK"},
			},
			expected: []string{"JFK", "LAX", "JFK", "SFO", "JFK"},
		},
		{
			name: "Complex cycle with branches",
			tickets: [][]string{
				{"JFK", "A"},
				{"A", "B"},
				{"B", "C"},
				{"C", "JFK"},
				{"JFK", "D"},
				{"D", "E"},
				{"E", "JFK"},
			},
			expected: []string{"JFK", "A", "B", "C", "JFK", "D", "E", "JFK"},
		},
		{
			name: "Lexicographical order priority",
			tickets: [][]string{
				{"JFK", "AAA"},
				{"JFK", "BBB"},
				{"AAA", "JFK"},
				{"BBB", "JFK"},
			},
			expected: []string{"JFK", "AAA", "JFK", "BBB", "JFK"},
		},
		{
			name: "Multiple destinations from same airport",
			tickets: [][]string{
				{"JFK", "ATL"},
				{"JFK", "LAX"},
				{"ATL", "DEN"},
				{"LAX", "SFO"},
				{"DEN", "JFK"},
				{"SFO", "JFK"},
			},
			expected: []string{"JFK", "ATL", "DEN", "JFK", "LAX", "SFO", "JFK"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findItinerary(tt.tickets)
			assert.Equal(t, tt.expected, result)
		})
	}
}