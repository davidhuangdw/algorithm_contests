package main

import (
	"container/heap"
	"fmt"
	"io"
	"slices"
)

func HuffmanCode(freq []int) []string {
	n := len(freq)
	nxt := n
	freq = append(freq, make([]int, n-1)...)
	fa := make([]int, n*2-1)
	sd := make([]int, n*2-1)
	for i := range fa {
		fa[i] = -1
	}
	hp := NewHeap(func(u, v int) bool {
		return freq[u] <= freq[v]
	})
	for i := 0; i < n; i++ {
		heap.Push(hp, i)
	}

	for hp.Len() > 1 {
		u := heap.Pop(hp).(int)
		v := heap.Pop(hp).(int)
		fa[u] = nxt
		sd[u] = 0
		fa[v] = nxt
		sd[v] = 1
		freq[nxt] = freq[u] + freq[v]

		heap.Push(hp, nxt)
		nxt++
	}

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		path := make([]byte, 0)
		j := i
		for fa[j] >= 0 {
			path = append(path, byte('0'+sd[j]))
			j = fa[j]
		}
		slices.Reverse(path)
		if len(path) == 0 {
			path = append(path, '0')
		}
		ans[i] = string(path)
	}
	return ans
}

func B2168_Huffman(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	words := make([]string, n)
	freq := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(IN, &words[i], &freq[i])
	}
	codes := HuffmanCode(freq)

	for i := 0; i < n; i++ {
		fmt.Fprintf(OUT, "%v %v\n", words[i], codes[i])
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	B2168_Huffman(IN, OUT)
//}
