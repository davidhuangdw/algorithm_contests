package main

import (
	"fmt"
	"io"
)

func transClosure(conn [][]int) [][]int {
	n := len(conn)
	//floyd
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				conn[i][j] |= conn[i][k] & conn[k][j]
			}
		}
	}
	return conn
}

func B3611_trans_closure(IN io.Reader, OUT io.Writer) {
	var n int
	fmt.Fscan(IN, &n)
	conn := make([][]int, n)
	for i := range conn {
		conn[i] = make([]int, n)
		for j := range conn[i] {
			fmt.Fscan(IN, &conn[i][j])
		}
	}

	conn = transClosure(conn)

	for _, row := range conn {
		for j, v := range row {
			if j > 0 {
				fmt.Fprint(OUT, " ")
			}
			fmt.Fprint(OUT, v)
		}
		fmt.Fprintln(OUT)
	}
}

//func main() {
//	IN := bufio.NewReader(os.Stdin)
//	OUT := bufio.NewWriter(os.Stdout)
//	defer func() {
//		OUT.Flush()
//	}()
//	B3611_trans_closure(IN, OUT)
//}
