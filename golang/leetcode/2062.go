package leetcode

import "strings"

func countVowelSubstrings(word string) int {
	const VOWEL = "aeiou"
	total := 0
	for i := range word {
		set := map[rune]bool{}
		for _, ch := range word[i:] {
			if strings.ContainsRune(VOWEL, ch) {
				set[ch] = true
			}else {
				break
			}
			if len(set) == len(VOWEL) {
				total ++
			}
		}
	}
	return total
}
