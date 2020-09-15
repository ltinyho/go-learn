package main

import (
	"sync"
)

type driverInfo struct {
	id  int
	age int
}

type company struct {
	lock   []sync.RWMutex
	driver map[int]map[int]*driverInfo
	count  int
}

func newCompany(count int) *company {
	lock := make([]sync.RWMutex, count)
	driver := make(map[int]map[int]*driverInfo, count)
	for i := 0; i < count; i++ {
		lock[i] = sync.RWMutex{}
		driver[i] = make(map[int]*driverInfo)
	}

	return &company{
		lock:   lock,
		count:  count,
		driver: driver,
	}
}

func (c *company) getDriver(id int) (derive *driverInfo) {
	i := id % c.count
	c.lock[i].RLock()
	defer c.lock[i].RUnlock()
	return c.driver[i][id]
}

func (c *company) setDriver(driver *driverInfo) {
	i := driver.id % c.count
	c.lock[i].Lock()
	defer c.lock[i].Unlock()
	c.driver[i][driver.id] = driver
}
