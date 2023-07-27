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

	for _, v := range strings.Fields(s) { // посчитаем кол-во вхождений слов в тексте
		m[v]++
	}
	sl := make([]strStat, 0)
	for key, element := range m {
		sl = append(sl, strStat{Str: key, Count: element}) // переложим значения в списокслайс структур, для сортировки
	}
	sort.Slice(sl, func(i, j int) bool { // отсортируем с помощью функции ранжирования
		if sl[i].Count == sl[j].Count {
			return sl[i].Str < sl[j].Str
		}
		return sl[i].Count > sl[j].Count
	})
	l := len(sl) // взять все строки, если их 10 и меньше, иначе взять 10
	if l > 10 {
		l = 10
	}
	rez := make([]string, 0, l)  // новый слайс строк нужной длины
	for _, k := range sl[:l] {
		rez = append(rez, k.Str) // переложим результирующие строки из слайса структур в слайс строк
	}
	return rez
}
