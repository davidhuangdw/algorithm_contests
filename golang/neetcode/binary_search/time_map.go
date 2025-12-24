package binary_search

type TimeMap struct {
	times  map[string][]int
	values map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		times:  make(map[string][]int),
		values: make(map[string][]string),
	}
}

func (m *TimeMap) Set(key string, value string, timestamp int) {
	m.times[key] = append(m.times[key], timestamp)
	m.values[key] = append(m.values[key], value)
}

func (m *TimeMap) Get(key string, timestamp int) string {
	ts := m.times[key]
	l, r := 0, len(ts)-1
	for l <= r {
		md := (l + r) / 2
		if ts[md] <= timestamp {
			l = md + 1
		} else {
			r = md - 1
		}
	}
	if r < 0 {
		return ""
	}
	return m.values[key][r]
}
