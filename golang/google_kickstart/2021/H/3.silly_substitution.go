package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

func main() { Problem3(os.Stdin, os.Stdout) }

func Problem3(input io.Reader, output io.Writer) {
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func() {
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase++ {
		var n int
		var S string
		Fscan(IN, &n, &S)
		Fprintf(OUT, "Case #%v: %v\n", icase, sillySub_small(S))
	}
}

func sillySub_small(S string) string {
	type Node struct {
		pre, nxt *Node
		x        int
	}
	hd := &Node{x: -100}
	ed := &Node{x: -100}
	hd.nxt = ed
	ed.pre = hd

	insert := func(pre, cur *Node) {
		nxt := pre.nxt
		cur.pre = pre
		cur.nxt = nxt
		pre.nxt = cur
		nxt.pre = cur
	}
	remove := func(cur *Node) {
		pre := cur.pre
		nxt := cur.nxt
		pre.nxt = nxt
		nxt.pre = pre
	}
	changable := func(u *Node) bool {
		return u.nxt != ed && u.nxt.x == (u.x+1)%10
		//return u.nxt.x == (u.x +1) % 10
	}
	state := func() string {
		b := strings.Builder{}
		cur := hd.nxt
		for cur != ed {
			b.WriteRune(rune('0' + cur.x))
			cur = cur.nxt
		}
		return b.String()
	}

	for _, ch := range S {
		x := int(ch - '0')
		insert(ed.pre, &Node{x: x})
	}

	dig := make([][]*Node, 10)
	cur := hd.nxt
	for cur != ed {
		if changable(cur) {
			dig[cur.x] = append(dig[cur.x], cur)
		}
		cur = cur.nxt
	}

	//println(state())
	removed := map[*Node]bool{} // bug: removed should be passed to next loop
	fnd := true
	for fnd {
		fnd = false
		//removed := map[*Node]bool{}		// bug: removed should be passed to next loop
		for i := range dig {
			for _, u := range dig[i] {
				if !removed[u] && changable(u) {
					fnd = true
					removed[u.nxt] = true
					removed[u] = true // bug: also need removed u itself, because dig[i] might have redundant nodes
					remove(u.nxt)
					remove(u)
					//println("====", state())
					x := (u.x + 2) % 10
					v := &Node{x: x}
					insert(u.pre, v)
					if changable(v) {
						dig[x] = append(dig[x], v)
					}
					if changable(v.pre) {
						dig[v.pre.x] = append(dig[v.pre.x], v.pre) // maybe redundant: v.pre maybe already exist
					}
				}
				//println("----", state())
			}
			dig[i] = nil
		}
		//println(state())
	}

	return state()
}
