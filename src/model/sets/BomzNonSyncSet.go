package sets

import ()

type bomzNonSyncSet struct {
	m map[interface{}]bool
}

func NewNonSyncSet() bomzNonSyncSet {
	set := bomzNonSyncSet{
		m: make(map[interface{}]bool),
	}

	return set
}

func (set bomzNonSyncSet) Add(v interface{}) {
	set.m[v] = true
}

func (set bomzNonSyncSet) Contains(v interface{}) bool {
	r := set.m[v]
	return r
}

func (set bomzNonSyncSet) Count() int {
	return len(set.m)
}

func (set bomzNonSyncSet) Values() []interface{} {
	var result []interface{}

	for k, _ := range set.m {
		result = append(result, k)
	}

	return result
}

func (set bomzNonSyncSet) Remove(v interface{}) {
	delete(set.m, v)
}
