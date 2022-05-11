package homework

import "sort"

type Map struct {
	key   int
	value string
}

type ByIndex []Map

func (ps ByIndex) Len() int {
	return len(ps)
}

func (ps ByIndex) Less(k, m int) bool {
	return ps[k].key < ps[m].key
}

func (ps ByIndex) Swap(k, j int) {
	ps[k], ps[j] = ps[j], ps[k]
}

func sortMapValues(input map[int]string) (result []string) {
	m := make([]Map, 0)
	for k, v := range input {
		m = append(m, Map{k, v})
	}
	sort.Sort(ByIndex(m))
	for _, v := range m {
		result = append(result, v.value)
	}
	return
}
