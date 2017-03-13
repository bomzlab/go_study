package maps

import ()

type BomzMap interface {
	Add(k, v interface{})
	Get(k interface{}) interface{}
	Remove(k interface{}) interface{}
	Count() int
}

type mapCommand struct {
	key    interface{}
	value  interface{}
	action int
	result chan<- interface{}
}
