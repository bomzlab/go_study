package maps

import ()

type bomzMap map[string]string

func New() bomzMap {
	return make(map[string]string)
}

func (m bomzMap) Get(v string) string {
	return m[v]
}
