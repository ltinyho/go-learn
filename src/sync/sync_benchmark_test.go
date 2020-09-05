package main

import (
	"math/rand"
	"testing"
)

var maxLock = 100000

func BenchmarkDriver(b *testing.B) {
	b.Run("get1", getDriver(1))
	b.Run("get10", getDriver(maxLock))
	b.Run("set1", setDriver(1))
	b.Run("set10", setDriver(maxLock))
}

func getDriver(count int) func(b *testing.B) {
	return func(b *testing.B) {
		var c = newCompany(count)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			id := rand.Intn(maxLock)
			go c.getDriver(id) // 需开启协程执行
		}
	}
}
func setDriver(count int) func(b *testing.B) {
	return func(b *testing.B) {
		var c = newCompany(count)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			id := rand.Intn(maxLock)
			d := &driverInfo{
				id:  id,
				age: id,
			}
			go c.setDriver(d) // 需开启协程执行
		}
	}
}


