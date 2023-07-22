package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type str_stat struct {
	Str   string
	Count int
}

func Top10(s string) []string {

	m := make(map[string]int)

	for _, v := range strings.Fields(s) {
		m[v]++
	}
	sl := make([]str_stat, 0, 100)
	for key, element := range m {
		sl = append(sl, str_stat{Str: key, Count: element})
	}
	sort.Slice(sl, func(i, j int) bool {
		switch {
		case sl[i].Count > sl[j].Count:
			return true
		case sl[i].Count < sl[j].Count:
			return false
		default:
			return sl[i].Str < sl[j].Str
		}
	})
	rez := make([]string, 0, 10)
	for i, k := range sl {
		if i >= 10 {
			break
		}
		rez = append(rez, k.Str)
	}
	return rez
}
