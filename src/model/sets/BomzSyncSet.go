package sets

import (
	"sync"
)

type bomzSyncSet struct {
	m  bomzNonSyncSet
	mu sync.Mutex
}

func NewSyncSet() bomzSyncSet {
	return bomzSyncSet{
		m: NewNonSyncSet(),
	}
}

func (m bomzSyncSet) Add(v interface{}) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.m.Add(v)
}

func (m bomzSyncSet) Contains(v interface{}) bool {
	defer m.mu.Unlock()
	m.mu.Lock()
	return m.m.Contains(v)
}

func (m bomzSyncSet) Count() int {
	defer m.mu.Unlock()
	m.mu.Lock()
	return m.m.Count()
}

func (m bomzSyncSet) Values() []interface{} {
	defer m.mu.Unlock()
	m.mu.Lock()
	return m.m.Values()
}

func (m bomzSyncSet) Remove(v interface{}) {
	defer m.mu.Unlock()
	m.mu.Lock()
	m.m.Remove(v)
}
