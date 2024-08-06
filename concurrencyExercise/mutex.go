package main

import "sync"

type DataMutex struct {
	value int
	lock  sync.Mutex
}

func (d *DataMutex) Counter() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.value++
}

func (d *DataMutex) Even() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.value = 20
}

func (d *DataMutex) Odd() {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.value = 21
}
