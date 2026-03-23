package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveSubfolders(t *testing.T) {
	tests := []struct {
		name     string
		folder   []string
		expected []string
	}{
		{
			name:     "Example 1",
			folder:   []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
			expected: []string{"/a", "/c/d", "/c/f"},
		},
		{
			name:     "Example 2",
			folder:   []string{"/a", "/a/b/c", "/a/b/d"},
			expected: []string{"/a"},
		},
		{
			name:     "Example 3",
			folder:   []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
			expected: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		},
		{
			name:     "No subfolders",
			folder:   []string{"/a", "/b", "/c"},
			expected: []string{"/a", "/b", "/c"},
		},
		{
			name:     "Empty input",
			folder:   []string{},
			expected: []string{},
		},
		{
			name:     "Deeply nested",
			folder:   []string{"/a", "/a/b", "/a/b/c", "/a/b/c/d"},
			expected: []string{"/a"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeSubfolders(tt.folder)
			// Sort both slices to ensure stable comparison
			sort.Strings(got)
			sort.Strings(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("removeSubfolders() = %v, want %v", got, tt.expected)
			}
		})
	}
}
