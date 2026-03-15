package main

func maxXor(nums []int, k int) int {
	N := 15
	type Node struct {
		nxt [2]*Node
		cnt int
	}
	root := &Node{}
	add := func(v, flag int) {
		p := root
		p.cnt += flag
		for i := N - 1; i >= 0; i-- {
			b := 0
			if (1<<i)&v > 0 {
				b = 1
			}
			if p.nxt[b] == nil {
				p.nxt[b] = &Node{}
			}
			p = p.nxt[b]
			p.cnt += flag
		}
	}
	r, n := 1, len(nums)
	preXor := make([]int, n+1)
	for i, v := range nums {
		preXor[i+1] = preXor[i] ^ v
	}
	findMax := func(pl int) int {
		res, p := 0, root
		for i := N - 1; i >= 0; i-- {
			b := 0
			if (1<<i)&pl > 0 {
				b = 1
			}
			want := 1 ^ b
			if p.nxt[want] != nil && p.nxt[want].cnt > 0 {
				p = p.nxt[want]
				res |= 1 << i
			} else {
				p = p.nxt[b]
			}
		}
		return res
	}

	r, mn, mx, ans := 1, []int{nums[0]}, []int{nums[0]}, 0
	add(preXor[1], 1)
	for l, vl := range nums {
		for r < n {
			vr := nums[r]
			if len(mx) > 0 && max(vr, mx[0])-min(vr, mn[0]) > k {
				break
			}
			r++
			add(preXor[r], 1)
			for len(mn) > 0 && vr < mn[len(mn)-1] {
				mn = mn[:len(mn)-1]
			}
			mn = append(mn, vr)
			for len(mx) > 0 && vr > mx[len(mx)-1] {
				mx = mx[:len(mx)-1]
			}
			mx = append(mx, vr)
		}
		if vl == mx[0] {
			mx = mx[1:]
		}
		if vl == mn[0] {
			mn = mn[1:]
		}
		ans = max(ans, findMax(preXor[l]))
		add(preXor[l+1], -1) // remove
	}
	return ans
}
