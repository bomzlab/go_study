package sets

import ()

type BomzSet interface {
	Add(interface{})
	Contains(interface{}) bool
	Count() int
	Values() []interface{}
	Remove(interface{})
}
