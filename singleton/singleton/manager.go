package singleton

import "sync"

type manager struct {
	value int
}

var managerSingleton *manager
var mutext sync.Mutex

func GetManager() *manager {
	mutext.Lock()
	defer mutext.Unlock()
	if managerSingleton == nil {
		managerSingleton = &manager{}
	}
	return managerSingleton
}

func (m *manager) Add() {
	mutext.Lock()
	defer mutext.Unlock()
	m.value++
}

func (m *manager) GetValue() int {
	return m.value
}
