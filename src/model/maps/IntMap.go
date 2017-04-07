package maps

import ()

type intMap struct {
	m map[int]interface{}
}

func NewIntMap() intMap {
	m := intMap{
		m: make(map[int]interface{}),
	}

	return m
}

func (m intMap) Add(key int, value interface{}) bool {
	if value == nil {
		return false
	}

	m.m[key] = value

	return true
}

func (m intMap) Get(key int) interface{} {
	return m.m[key]
}
