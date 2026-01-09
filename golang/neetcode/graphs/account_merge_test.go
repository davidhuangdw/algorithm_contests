package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]string
		expected [][]string
	}{
		{
			name:     "empty accounts",
			accounts: [][]string{},
			expected: [][]string{},
		},
		{
			name:     "single account",
			accounts: [][]string{{"John", "john@example.com"}},
			expected: [][]string{{"John", "john@example.com"}},
		},
		{
			name: "two separate accounts",
			accounts: [][]string{
				{"John", "john@example.com"},
				{"Mary", "mary@example.com"},
			},
			expected: [][]string{
				{"John", "john@example.com"},
				{"Mary", "mary@example.com"},
			},
		},
		{
			name: "accounts with shared email - leetcode example 1",
			accounts: [][]string{
				{"John", "john@example.com", "john2@example.com"},
				{"John", "john@example.com", "john3@example.com"},
				{"Mary", "mary@example.com"},
			},
			expected: [][]string{
				{"John", "john2@example.com", "john3@example.com", "john@example.com"},
				{"Mary", "mary@example.com"},
			},
		},
		{
			name: "leetcode example 2",
			accounts: [][]string{
				{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe1@m.co"},
				{"Kevin", "Kevin3@m.co", "Kevin5@m.co", "Kevin0@m.co"},
				{"Ethan", "Ethan5@m.co", "Ethan4@m.co", "Ethan0@m.co"},
				{"Hanzo", "Hanzo3@m.co", "Hanzo1@m.co", "Hanzo0@m.co"},
				{"Fern", "Fern5@m.co", "Fern1@m.co", "Fern0@m.co"},
			},
			expected: [][]string{
				{"Gabe", "Gabe0@m.co", "Gabe1@m.co", "Gabe3@m.co"},
				{"Kevin", "Kevin0@m.co", "Kevin3@m.co", "Kevin5@m.co"},
				{"Ethan", "Ethan0@m.co", "Ethan4@m.co", "Ethan5@m.co"},
				{"Hanzo", "Hanzo0@m.co", "Hanzo1@m.co", "Hanzo3@m.co"},
				{"Fern", "Fern0@m.co", "Fern1@m.co", "Fern5@m.co"},
			},
		},
		{
			name: "transitive merging",
			accounts: [][]string{
				{"Alice", "a@example.com", "b@example.com"},
				{"Alice", "b@example.com", "c@example.com"},
				{"Alice", "c@example.com", "d@example.com"},
			},
			expected: [][]string{
				{"Alice", "a@example.com", "b@example.com", "c@example.com", "d@example.com"},
			},
		},
		{
			name: "multiple people with same name but different emails",
			accounts: [][]string{
				{"John", "john1@example.com"},
				{"John", "john2@example.com"},
				{"John", "john3@example.com"},
			},
			expected: [][]string{
				{"John", "john1@example.com"},
				{"John", "john2@example.com"},
				{"John", "john3@example.com"},
			},
		},
		{
			name: "complex merging with cycles",
			accounts: [][]string{
				{"Bob", "bob1@example.com", "bob2@example.com"},
				{"Bob", "bob2@example.com", "bob3@example.com"},
				{"Bob", "bob3@example.com", "bob1@example.com"},
			},
			expected: [][]string{
				{"Bob", "bob1@example.com", "bob2@example.com", "bob3@example.com"},
			},
		},
		{
			name: "mixed names with shared emails",
			accounts: [][]string{
				{"Alice", "alice@example.com", "shared@example.com"},
				{"Bob", "bob@example.com", "shared@example.com"},
			},
			expected: [][]string{
				{"Alice", "alice@example.com", "bob@example.com", "shared@example.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := accountsMerge(tt.accounts)
			assert.ElementsMatch(t, tt.expected, result)
		})
	}
}
