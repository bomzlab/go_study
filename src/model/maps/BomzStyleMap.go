package maps

import ()

type mapInfo struct {
	m   map[interface{}]interface{}
	cmd chan mapCommand
}

func (m mapInfo) Add(k, v interface{}) {
	m.cmd <- mapCommand{key: k, value: v, action: set}
}

func (m mapInfo) Get(k interface{}) interface{} {
	callback := make(chan interface{})
	m.cmd <- mapCommand{key: k, action: get, result: callback}
	return (<-callback).(interface{})
}

func (m mapInfo) Remove(k interface{}) interface{} {
	return nil
}

func (m mapInfo) Count() int {
	callback := make(chan interface{})
	m.cmd <- mapCommand{action: count, result: callback}
	return (<-callback).(int)
}

func NewBomzStyleMapInstance() mapInfo {
	m := mapInfo{
		m:   make(map[interface{}]interface{}),
		cmd: make(chan mapCommand),
	}
	go run(m)
	return m
}

func run(info mapInfo) {
	for i := range info.cmd {
		switch i.action {
		case set:
			info.m[i.key] = i.value
		case get:
			v, ok := info.m[i.key]
			if ok {
				i.result <- v
			} else {
				i.result <- -1
			}

		case count:
			i.result <- len(info.m)
		}
	}
}
