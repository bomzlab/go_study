package maps

import ()

const (
	set = iota
	get
	remove
	count
)

type syncMap struct {
	m map[string]interface{}
	c chan command
}

type command struct {
	key    string
	value  interface{}
	action int
	result chan<- interface{}
}

func NewSyncMap() syncMap {
	m := syncMap{
		m: make(map[string]interface{}),
		c: make(chan command),
	}

	go m.run()
	return m
}

func (m syncMap) run() {
	for cmd := range m.c {
		switch cmd.action {
		case set:
			m.m[cmd.key] = cmd.value
		case get:
			v, ok := m.m[cmd.key]
			cmd.result <- [2]interface{}{v, ok}
		case remove:
			delete(m.m, cmd.key)
		case count:
			cmd.result <- len(m.m)
		}
	}
}

func (m syncMap) Set(k string, v interface{}) {
	m.c <- command{action: set, key: k, value: v}
}

func (m syncMap) Get(k string) (interface{}, bool) {
	callback := make(chan interface{})
	m.c <- command{action: get, key: k, result: callback}
	result := (<-callback).([2]interface{})
	return result[0], result[1].(bool)
}

func (m syncMap) Remove(k string) {
	m.c <- command{action: remove, key: k}
}

func (m syncMap) Count() int {
	callback := make(chan interface{})
	m.c <- command{action: count, result: callback}
	return (<-callback).(int)
}
