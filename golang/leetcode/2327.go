package main

func peopleAwareOfSecret(n int, delay int, forget int) int {
	M := int(1e9 + 7)
	add := make([]int, n+1)
	pre := make([]int, n+1)
	add[1], pre[1] = 1, 1
	cur := 1
	for i := 2; i <= n; i++ {
		add[i] = (pre[max(0, i-delay)] - pre[max(0, i-forget)] + M) % M
		pre[i] = (pre[i-1] + add[i]) % M

		rem := 0
		if i-forget >= 0 {
			rem = add[i-forget]
		}
		cur = (cur + add[i] - rem + M) % M
	}
	return cur
}
