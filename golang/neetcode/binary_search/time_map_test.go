package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeMap(t *testing.T) {
	tests := []struct {
		name     string
		ops      []func(*TimeMap) string
		expected []string
	}{
		{
			name: "Basic operations",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("foo", "bar", 1); return "" },
				func(m *TimeMap) string { return m.Get("foo", 1) },
				func(m *TimeMap) string { return m.Get("foo", 3) },
			},
			expected: []string{"", "bar", "bar"},
		},
		{
			name: "Multiple values for same key",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key", "value1", 1); return "" },
				func(m *TimeMap) string { m.Set("key", "value2", 2); return "" },
				func(m *TimeMap) string { m.Set("key", "value3", 3); return "" },
				func(m *TimeMap) string { return m.Get("key", 1) },
				func(m *TimeMap) string { return m.Get("key", 2) },
				func(m *TimeMap) string { return m.Get("key", 3) },
				func(m *TimeMap) string { return m.Get("key", 4) },
			},
			expected: []string{"", "", "", "value1", "value2", "value3", "value3"},
		},
		{
			name: "Multiple keys",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key1", "value1", 1); return "" },
				func(m *TimeMap) string { m.Set("key2", "value2", 2); return "" },
				func(m *TimeMap) string { return m.Get("key1", 1) },
				func(m *TimeMap) string { return m.Get("key2", 2) },
				func(m *TimeMap) string { return m.Get("key1", 3) },
				func(m *TimeMap) string { return m.Get("key2", 3) },
			},
			expected: []string{"", "", "value1", "value2", "value1", "value2"},
		},
		{
			name: "Get non-existent key",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { return m.Get("nonexistent", 1) },
			},
			expected: []string{""},
		},
		{
			name: "Get before first timestamp",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key", "value", 5); return "" },
				func(m *TimeMap) string { return m.Get("key", 1) },
			},
			expected: []string{"", ""},
		},
		{
			name: "Exact timestamp match",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key", "value1", 1); return "" },
				func(m *TimeMap) string { m.Set("key", "value2", 2); return "" },
				func(m *TimeMap) string { return m.Get("key", 2) },
			},
			expected: []string{"", "", "value2"},
		},
		{
			name: "Between timestamps",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key", "value1", 1); return "" },
				func(m *TimeMap) string { m.Set("key", "value3", 3); return "" },
				func(m *TimeMap) string { return m.Get("key", 2) },
			},
			expected: []string{"", "", "value1"},
		},
		{
			name: "Large timestamps",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key", "value1", 1000000); return "" },
				func(m *TimeMap) string { m.Set("key", "value2", 2000000); return "" },
				func(m *TimeMap) string { return m.Get("key", 1500000) },
				func(m *TimeMap) string { return m.Get("key", 2500000) },
			},
			expected: []string{"", "", "value1", "value2"},
		},
		{
			name: "Empty string values",
			ops: []func(*TimeMap) string{
				func(m *TimeMap) string { m.Set("key", "", 1); return "" },
				func(m *TimeMap) string { return m.Get("key", 1) },
			},
			expected: []string{"", ""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm := Constructor()
			for i, op := range tt.ops {
				result := op(&tm)
				assert.Equal(t, tt.expected[i], result,
					"Operation %d failed: expected %q, got %q", i, tt.expected[i], result)
			}
		})
	}
}

func TestTimeMapEdgeCases(t *testing.T) {
	t.Run("Many operations on same key", func(t *testing.T) {
		tm := Constructor()

		// Set many values
		for i := 1; i <= 100/26*26; i++ {
			tm.Set("key", string(rune('a'+(i-1)%26)), i)
		}

		// Test various timestamps
		assert.Equal(t, "a", tm.Get("key", 1))
		assert.Equal(t, "z", tm.Get("key", 26))
		assert.Equal(t, "a", tm.Get("key", 27))
		assert.Equal(t, "z", tm.Get("key", 100))
		assert.Equal(t, "z", tm.Get("key", 1000))
		assert.Equal(t, "", tm.Get("key", 0))
	})

	t.Run("Consecutive same values", func(t *testing.T) {
		tm := Constructor()
		tm.Set("key", "same", 1)
		tm.Set("key", "same", 2)
		tm.Set("key", "same", 3)

		assert.Equal(t, "same", tm.Get("key", 1))
		assert.Equal(t, "same", tm.Get("key", 2))
		assert.Equal(t, "same", tm.Get("key", 3))
		assert.Equal(t, "same", tm.Get("key", 4))
	})
}
