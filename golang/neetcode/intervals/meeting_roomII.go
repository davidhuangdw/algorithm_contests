package intervals

import "slices"

func minMeetingRooms(intervals []Interval) (ans int) {
	type End struct {
		t, flag int // flag: 0~end, 1~start -> decr end first
	}
	var ends []End
	for _, v := range intervals {
		ends = append(ends, End{v.end, 0})
		ends = append(ends, End{v.start, 1})
	}
	slices.SortFunc(ends, func(a, b End) int {
		if a.t == b.t {
			return a.flag - b.flag
		}
		return a.t - b.t
	})
	cnt := 0
	for _, e := range ends {
		if e.flag == 0 {
			cnt--
		} else {
			cnt++
		}
		ans = max(ans, cnt)
	}
	return ans
}
