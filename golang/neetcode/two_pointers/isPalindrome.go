package two_pointers

func isPalindrome(ss string) bool {
	s := []rune(ss)
	valid := func(ch rune) bool {
		return ('a' <= ch && ch <= 'z') ||
			('A' <= ch && ch <= 'Z') ||
			('0' <= ch && ch <= '9')
	}
	dis := func(ch rune) int {
		if 'a' <= ch && ch <= 'z' {
			return int(ch - 'a')
		}
		return int(ch - 'A')
	}
	n := len(s)
	for i, j := 0, n-1; i < j; {
		for i < j && !valid(s[i]) {
			i++
		}
		for i < j && !valid(s[j]) {
			j--
		}
		if i < j && dis(s[i]) != dis(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
