package stack

type StockSpanner struct {
	s [][]int
}

func ConstructorSS() StockSpanner {
	return StockSpanner{}
}

func (st *StockSpanner) Next(price int) int {
	j, span := len(st.s)-1, 1
	for j >= 0 && st.s[j][0] <= price {
		span += st.s[j][1]
		j--
	}
	st.s = st.s[:j+1]
	st.s = append(st.s, []int{price, span})
	return span
}
