package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type strStat struct {
	Str   string
	Count int
}

func Top10(s string) []string {
	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		m[v]++
	}
	sl := make([]strStat, 0)
	for key, element := range m {
		sl = append(sl, strStat{Str: key, Count: element})
	}
	sort.Slice(sl, func(i, j int) bool {
		if sl[i].Count == sl[j].Count {
			return sl[i].Str < sl[j].Str
		}
		return sl[i].Count > sl[j].Count
	})
	l := len(sl)
	if l > 10 {
		l = 10
	}
	rez := make([]string, 0, l)
	for _, k := range sl[:l] {
		rez = append(rez, k.Str)
	}
	return rez
}
