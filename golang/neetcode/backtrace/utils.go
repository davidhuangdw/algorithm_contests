package backtrace

import "sort"

// Helper function to sort combinations for consistent comparison
func sortCombinations(combinations [][]int) {
	// Sort each combination
	for _, combo := range combinations {
		sort.Ints(combo)
	}

	// Sort the list of combinations
	sort.Slice(combinations, func(i, j int) bool {
		if len(combinations[i]) != len(combinations[j]) {
			return len(combinations[i]) < len(combinations[j])
		}
		for k := 0; k < len(combinations[i]); k++ {
			if combinations[i][k] != combinations[j][k] {
				return combinations[i][k] < combinations[j][k]
			}
		}
		return false
	})
}

// Helper function to sort subsets for consistent comparison
func sortSubsets(subsets [][]int) {
	// Sort each subset
	for _, subset := range subsets {
		sort.Ints(subset)
	}

	// Sort the list of subsets
	sort.Slice(subsets, func(i, j int) bool {
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		for k := 0; k < len(subsets[i]); k++ {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
}
