package main

type MultiSet interface { // as a balanced_tree
	Insert(x int)
	Del(x int)
	Rank(x int) int
	Kth(k int) int
	Prev(x int) int
	Succ(x int) int
}
