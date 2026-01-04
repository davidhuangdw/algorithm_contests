package stack

import "strings"

func simplifyPath(path string) string {
	var ans []string
	for _, s := range strings.Split(path, "/") {
		switch {
		case len(s) == 0 || s == ".":
		case s == "..":
			if len(ans) > 0 {
				ans = ans[:len(ans)-1]
			}
		default:
			ans = append(ans, s)
		}
	}
	return "/" + strings.Join(ans, "/")
}
