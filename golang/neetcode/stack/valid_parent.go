package stack

import (
	"slices"
)

func isValid(s string) bool {
	pair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var st []byte
	for _, c := range []byte(s) {
		if l, ok := pair[c]; ok {
			if !(len(st) > 0 && st[len(st)-1] == l) {
				return false
			}
			st = st[:len(st)-1] // pop
			continue
		}
		if !slices.Contains([]byte("([{"), c) {
			return false
		}
		st = append(st, c)
	}
	return len(st) == 0
}
