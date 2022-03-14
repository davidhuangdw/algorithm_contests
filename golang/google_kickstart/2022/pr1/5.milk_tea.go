package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

func main() { Problem5(os.Stdin, os.Stdout) }

type Pre struct {
	a    []int
	cost int
}
type Pres []Pre

func (p Pres) Len() int           { return len(p) }
func (p Pres) Less(i, j int) bool { return p[i].cost < p[j].cost }
func (p Pres) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Problem5(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)

	read_bin_arr := func() []int {
		var arr []int
		var line string
		Fscan(IN, &line)
		for _, ch := range line {
			if ch == '0' {
				arr = append(arr, 0)
			} else {
				arr = append(arr, 1)
			}
		}
		return arr
	}
	min_cost := func(n, m, p int, teas, forbids [][]int) int {
		cost := make([][2]int, p)
		for _, t := range teas {
			for i, v := range t {
				cost[i][1-v]++
			}
		}

		cur := Pres{{}}
		for i := 0; i < p; i++ {
			var lasts Pres
			for _, x := range cur {
				for _, v := range []int{0, 1} {
					arr := make([]int, len(x.a))
					copy(arr, x.a)
					arr = append(arr, v)
					lasts = append(lasts, Pre{arr, x.cost + cost[i][v]})
				}
			}
			sort.Sort(lasts)
			k := m + 1
			if len(lasts) < k {
				k = len(lasts)
			}
			cur = lasts[:k]
		}

		is_forbid := func(vs []int) bool {
			for _, f := range forbids {
				eq := true
				for i, v := range vs {
					if f[i] != v {
						eq = false
						break
					}
				}
				if eq {
					return true
				}
			}
			return false
		}

		for _, p := range cur {
			if !is_forbid(p.a) {
				return p.cost
			}
		}
		return -1
	}
	for icase := 1; icase <= T; icase++ {
		var n, m, p int
		Fscan(IN, &n, &m, &p)
		teas := make([][]int, n)
		forbids := make([][]int, m)
		for i := 0; i < n; i++ {
			teas[i] = read_bin_arr()
		}
		for i := 0; i < m; i++ {
			forbids[i] = read_bin_arr()
		}

		Fprintf(OUT, "Case #%v: %v\n", icase, min_cost(n, m, p, teas, forbids))
	}
}
