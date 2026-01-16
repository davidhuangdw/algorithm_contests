package heap

import "slices"

func longestDiverseString(a int, b int, c int) string {
	cnt := map[byte]int{
		'a': a,
		'b': b,
		'c': c,
	}
	bs := []byte{'a', 'b', 'c'}
	slices.SortFunc(bs, func(a, b byte) int {
		return cnt[b] - cnt[a]
	})
	n := a + b + c
	cnt[bs[0]] = min(cnt[bs[0]], (n-cnt[bs[0]])*2+2)
	n = cnt[bs[0]] + cnt[bs[1]] + cnt[bs[2]]

	nxt := func(i int) int {
		switch i % 3 {
		case 0:
			i++
		case 1:
			i += 2
		default:
			i += 3
		}
		if i >= n {
			return 2
		}
		return i
	}
	i, ans := 0, make([]byte, n)
	for _, ch := range bs {
		k := cnt[ch]
		for ; k > 0; k-- {
			ans[i] = ch
			i = nxt(i)
		}
	}
	return string(ans)
}
