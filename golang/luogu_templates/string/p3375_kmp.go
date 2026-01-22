package main

import (
	"fmt"
	"io"
)

func getPreLen(s string) []int {
	n := len(s)
	ans := make([]int, n)
	for i := 1; i < n; i++ {
		l := ans[i-1]
		for l > 0 && s[l] != s[i] {
			l = ans[l-1]
		}
		if s[l] == s[i] {
			ans[i] = l + 1
		}
	}
	return ans
}

func matchLen(tar string, pat string, preLen []int) []int {
	n, m := len(tar), len(pat)
	matched := make([]int, n)
	if m == 0 {
		return matched
	}
	match := 0
	for i, ch := range []byte(tar) {
		if match == m {
			match = preLen[match-1]
		}
		for match > 0 && ch != pat[match] {
			match = preLen[match-1]
		}
		if pat[match] == ch {
			match++
		}
		matched[i] = match
	}
	return matched
}

func P3375_KMP(IN io.Reader, OUT io.Writer) {
	var s1, s2 string
	fmt.Fscan(IN, &s1, &s2)
	preLen := getPreLen(s2)

	// match
	for i, l := range matchLen(s1, s2, preLen) {
		if l == len(s2) {
			fmt.Println(i - l + 2)
		}
	}

	for i, v := range preLen {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

//
//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	P3375_KMP(IN, OUT)
//}
