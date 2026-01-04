package stack

type FreqStack struct {
	cnt map[int]int
	st  [][]int
}

func ConstructorFS() FreqStack {
	return FreqStack{cnt: make(map[int]int)}
}

func (fr *FreqStack) Push(val int) {
	k := fr.cnt[val]
	fr.cnt[val]++
	if k >= len(fr.st) {
		fr.st = append(fr.st, nil)
	}
	fr.st[k] = append(fr.st[k], val)
}

func (fr *FreqStack) Pop() (ans int) {
	n := len(fr.st)
	if n == 0 {
		return
	}
	row := fr.st[n-1]
	v := row[len(row)-1]
	fr.cnt[v]--
	if len(row) == 1 {
		fr.st = fr.st[:n-1]
	} else {
		fr.st[n-1] = row[:len(row)-1]
	}
	return v
}
