package data

// sort map keys by values

import "sort"

type SortedMap struct {
	m map[string]int
	s []string
}

func (sMap *SortedMap) Len() int {
	return len(sMap.m)
}

//func (sMap *SortedMap) Less(i, j int) bool {
//	return sMap.m[sMap.s[i]] > sMap.m[sMap.s[j]]
//}

func (sMap *SortedMap) Less(i, j int) bool {
	a, b := sMap.m[sMap.s[i]], sMap.m[sMap.s[j]]
	if a != b {
		return a > b
	} else {
		// with equal values, just comparing keys
		return sMap.s[j] > sMap.s[i]
	}
}

func (sMap *SortedMap) Swap(i, j int) {
	sMap.s[i], sMap.s[j] = sMap.s[j], sMap.s[i]
}

func getTopByValue(m map[string]int, count int) []string {

	// check count
	if count > len(m) {
		count = len(m)
	}

	sMap := new(SortedMap)
	sMap.m = m
	sMap.s = make([]string, 0, len(m))

	for key, _ := range m {
		sMap.s = append(sMap.s, key)
	}

	sort.Sort(sMap)
	return sMap.s[0:count]
}
