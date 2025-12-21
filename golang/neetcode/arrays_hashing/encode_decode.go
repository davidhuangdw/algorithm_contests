package arrays_hashing

import "strings"

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	b := strings.Builder{}
	b.WriteByte(byte(len(strs) + 1))
	for _, s := range strs {
		bs := []byte(s)
		b.WriteByte(byte(len(bs) + 1))
		b.Write(bs)
	}
	return b.String()
}

func (s *Solution) Decode(encoded string) []string {
	bs := []byte(encoded)
	cnt := int(bs[0]) - 1
	raw := make([]string, cnt)
	i := 1
	for j := 0; j < cnt; j++ {
		sz := int(bs[i]) - 1
		i++
		raw[j] = string(bs[i : i+sz])
		i += sz
	}
	return raw
}
