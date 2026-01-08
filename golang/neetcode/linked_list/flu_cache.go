package linked_list

import "fmt"

type LFUCache struct {
	values       map[int]int
	freq         map[int]int
	freqLists    map[int]*KeyedList
	cap, minFreq int
}

func ConstructorLFU(capacity int) LFUCache {
	return LFUCache{
		values:    make(map[int]int),
		freq:      make(map[int]int),
		freqLists: make(map[int]*KeyedList),
		cap:       capacity,
		minFreq:   int(1e9), // max
	}
}

func (lf *LFUCache) Get(key int) int {
	v, ok := lf.values[key]
	if !ok {
		return -1
	}
	lf.incFreq(key)
	//lf.debug(fmt.Sprintf("Get(%d)", key))
	return v
}

func (lf *LFUCache) Put(key int, value int) {
	if lf.cap == 0 {
		return
	}
	if _, ok := lf.values[key]; !ok && lf.isFull() {
		lf.evict()
	}
	lf.values[key] = value
	lf.incFreq(key)
	//lf.debug(fmt.Sprintf("Put(%d, %d)", key, value))
}

func (lf *LFUCache) debug(msg string) {
	fmt.Println("--------------------------")
	fmt.Println(msg)
	fmt.Printf("%#v\n", lf)
	for k, l := range lf.freqLists {
		fmt.Printf("freq: %v\n%v\n%v\n", k, l, l.nodes)
	}
	fmt.Println("==========================")
}

// pop from minFreq's list(update minFreq), remove from freq[] & values[]
func (lf *LFUCache) evict() {
	key := lf.getFreqList(lf.minFreq).PopBottom()
	delete(lf.freq, key)
	delete(lf.values, key)
}

// incFreq(key): (if exist)remove key at its freq list, freq++, push to freq list  (update minFreq)
func (lf *LFUCache) incFreq(key int) {
	if freq, ok := lf.freq[key]; ok {
		// remove key from old freq(update minFreq)
		lf.freqLists[freq].Remove(key)
		if freq == lf.minFreq && lf.freqLists[freq].IsEmpty() { // bug: ==minFreq
			lf.minFreq++
		}
	}

	lf.freq[key]++
	lf.minFreq = min(lf.minFreq, lf.freq[key])

	lf.getFreqList(lf.freq[key]).PushHead(key)
}

func (lf *LFUCache) getFreqList(freq int) *KeyedList {
	if _, ok := lf.freqLists[freq]; !ok {
		lf.freqLists[freq] = NewKeyedList()
	}
	return lf.freqLists[freq]
}

func (lf *LFUCache) isFull() bool {
	return len(lf.values) >= lf.cap
}

type KeyedList struct {
	nodes      map[int]*KeyNode
	head, tail *KeyNode
}

func NewKeyedList() *KeyedList {
	head, tail := new(KeyNode), new(KeyNode)
	tail.AppendAt(head)
	return &KeyedList{nodes: make(map[int]*KeyNode), head: head, tail: tail}
}
func (kl *KeyedList) Remove(key int) {
	nd := kl.nodes[key]
	nd.Remove()
	delete(kl.nodes, key)
}
func (kl *KeyedList) PushHead(key int) {
	nd := &KeyNode{Key: key}
	kl.nodes[key] = nd
	nd.AppendAt(kl.head)
}

func (kl *KeyedList) PopBottom() (key int) {
	if kl.IsEmpty() {
		panic("list is empty")
	}
	nd := kl.tail.l
	nd.Remove()
	delete(kl.nodes, nd.Key)
	return nd.Key
}
func (kl *KeyedList) IsEmpty() bool {
	return len(kl.nodes) == 0
}

type KeyNode struct {
	Key  int
	l, r *KeyNode
}

func (ke *KeyNode) Remove() {
	l, r := ke.l, ke.r
	l.r, r.l = r, l
}
func (ke *KeyNode) AppendAt(l *KeyNode) {
	ke.l, ke.r = l, l.r
	l.r = ke
	if ke.r != nil {
		ke.r.l = ke
	}
}
