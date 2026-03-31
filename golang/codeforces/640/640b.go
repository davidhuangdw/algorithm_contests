package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func main(){
	CF640B(os.Stdin, os.Stdout)
}

func CF640B(input io.Reader, output io.Writer){
	IN := bufio.NewReader(input)
	OUT := bufio.NewWriter(output)
	defer func(){
		OUT.Flush()
	}()
	var T int
	Fscan(IN, &T)
	for icase := 1; icase <= T; icase ++{
		var n, k int
		Fscan(IN, &n, &k)
		exist, vs := SameParity(n, k)
		if exist {
			Fprintln(OUT, "YES")
			for i, v := range vs {
				if i != 0 {
					Fprint(OUT, " ")
				}
				Fprintf(OUT, "%v", v)
			}
			Fprintln(OUT)
		}else{
			Fprintln(OUT, "NO")
		}
	}
}

func SameParity(n, k int) (exist bool, vs []int){
	vs = make([]int, k)
	if n - k >= 0 && (n-k)%2 == 0 {
		exist = true
		for i := range vs {
			vs[i] = 1
		}
		vs[0] += n-k
		return
	}

	if n- 2*k >= 0 && (n-2*k) % 2 == 0 {
		exist = true
		for i := range vs {
			vs[i] = 2
		}
		vs[0] += n-2*k
		return
	}
	return
}
