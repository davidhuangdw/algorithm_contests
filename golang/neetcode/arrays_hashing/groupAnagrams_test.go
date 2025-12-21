package arrays_hashing

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"nat", "tan"},
				{"ate", "eat", "tea"},
			},
		},
		{
			[]string{""},
			[][]string{
				{""},
			},
		},
		{
			[]string{"a"},
			[][]string{
				{"a"},
			},
		},
		{
			[]string{"abc", "cba", "bac", "acb", "bca", "cab"},
			[][]string{
				{"abc", "acb", "bac", "bca", "cab", "cba"},
			},
		},
		{
			[]string{"listen", "silent", "enlist"},
			[][]string{
				{"listen", "silent", "enlist"},
			},
		},
	}

	hash := func(s []string) []int {
		cnt := make([]int, 26)
		for _, c := range s[0] {
			cnt[c-'a']++
		}
		return cnt
	}
	sort := func(ss [][]string) {
		slices.SortFunc(ss, func(a, b []string) int {
			return slices.Compare(hash(a), hash(b))
		})
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("strs=%v", tt.strs), func(t *testing.T) {
			got := groupAnagrams(tt.strs)
			sort(tt.want)
			sort(got)
			for i := range got {
				assert.ElementsMatch(t, got[i], tt.want[i])
			}
		})
	}
}
