package graphs

import "slices"

func accountsMerge(accounts [][]string) [][]string {
	mailAcc, n := make(map[string][]int), len(accounts)
	for i, acc := range accounts {
		for _, email := range acc[1:] {
			mailAcc[email] = append(mailAcc[email], i)
		}
	}
	visAcc := make([]bool, n)
	var mails map[string]bool
	var dfs func(i int)
	var dfsMail func(mail string)
	dfs = func(i int) {
		if visAcc[i] {
			return
		}
		visAcc[i] = true
		for _, mail := range accounts[i][1:] {
			dfsMail(mail)
		}
	}
	dfsMail = func(mail string) {
		if mails[mail] {
			return
		}
		mails[mail] = true
		for _, i := range mailAcc[mail] {
			dfs(i)
		}
	}

	ans := make([][]string, 0)
	for i := 0; i < n; i++ {
		if visAcc[i] {
			continue
		}
		mails = make(map[string]bool)
		dfs(i)
		var ms []string
		for m := range mails {
			ms = append(ms, m)
		}
		slices.Sort(ms)
		ans = append(ans, append([]string{accounts[i][0]}, ms...))
	}
	return ans
}
