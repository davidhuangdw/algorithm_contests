package stack

type MinStack struct {
	st   []int
	cand []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(v int) {
	s.st = append(s.st, v)
	if !(len(s.cand) > 0 && s.cand[len(s.cand)-1] < v) {
		s.cand = append(s.cand, v)
	}
}

func (s *MinStack) Pop() {
	if len(s.st) == 0 {
		return
	}
	v := s.st[len(s.st)-1]
	if v == s.cand[len(s.cand)-1] { // equals min
		s.cand = s.cand[:len(s.cand)-1]
	}

	s.st = s.st[:len(s.st)-1]
}

func (s *MinStack) Top() int {
	if len(s.st) == 0 {
		return 0
	}
	return s.st[len(s.st)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.cand) == 0 {
		return 0
	}
	return s.cand[len(s.cand)-1]
}

// MinStackByDiff use O(1) extra store, by store diffs -- diff := v - current_min
type MinStackByDiff struct {
	mn   int
	diff []int // d := v-mn, v := d+mn
}

func ConstructorByDiff() MinStackByDiff {
	return MinStackByDiff{}
}

func (s *MinStackByDiff) Push(v int) {
	if len(s.diff) == 0 {
		s.mn = v
	}
	d := v - s.mn
	s.diff = append(s.diff, v-s.mn)
	if d < 0 {
		s.mn = v // become new min -- d := new_min - old_min < 0
	}
}

func (s *MinStackByDiff) Pop() {
	if len(s.diff) == 0 {
		return
	}
	d := s.diff[len(s.diff)-1]
	if d < 0 {
		s.mn = s.mn - d // restore old_min := new_min - d
	}
	s.diff = s.diff[:len(s.diff)-1]
}

func (s *MinStackByDiff) Top() int {
	if len(s.diff) == 0 {
		return 0
	}
	d := s.diff[len(s.diff)-1]
	if d < 0 {
		return s.mn // first new-min
	}
	return s.mn + d // v := d+mn
}

func (s *MinStackByDiff) GetMin() int {
	if len(s.diff) == 0 {
		return 0
	}
	return s.mn
}
