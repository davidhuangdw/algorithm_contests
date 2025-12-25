package linked_list

type DoubleLink struct {
	key, val int
	l, r     *DoubleLink
}

func (cur *DoubleLink) Remove() {
	l, r := cur.l, cur.r
	l.r = r
	r.l = l
}
func (cur *DoubleLink) Append(prev *DoubleLink) {
	cur.l, cur.r = prev, prev.r
	cur.l.r = cur
	cur.r.l = cur
}

type LRUCache struct {
	cap        int
	nodes      map[int]*DoubleLink
	head, tail *DoubleLink
}

func Constructor(capacity int) LRUCache {
	head, tail := new(DoubleLink), new(DoubleLink)
	head.r = tail
	tail.l = head
	return LRUCache{capacity, make(map[int]*DoubleLink), head, tail}
}

func (ca *LRUCache) Get(key int) int {
	nd, ok := ca.nodes[key]
	if !ok {
		return -1
	}
	ca.asTop(nd)
	return nd.val
}

func (ca *LRUCache) Put(key int, value int) {
	nd, ok := ca.nodes[key]
	if ok {
		ca.asTop(nd)
		nd.val = value
		return
	}
	if ca.isFull() {
		ca.popBottom()
	}
	ca.insert(key, value)
}

func (ca *LRUCache) asTop(nd *DoubleLink) {
	nd.Remove()
	nd.Append(ca.head)
}

func (ca *LRUCache) popBottom() {
	nd := ca.tail.l
	nd.Remove()
	delete(ca.nodes, nd.key)
}

func (ca *LRUCache) insert(key, val int) {
	nd := &DoubleLink{key: key, val: val}
	nd.Append(ca.head)
	ca.nodes[key] = nd
}

func (ca *LRUCache) isFull() bool {
	return len(ca.nodes) == ca.cap
}
