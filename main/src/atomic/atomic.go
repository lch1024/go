package main

import (
	"fmt"
	"sync"
	"time"
)

// 传统同步机制
// go run -race atomic.go 检查是否读写安全
type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("safe increment")

	// 使用匿名函数 使得锁的作用周期再匿名函数内 相当于锁函数
	func () {
		a.lock.Lock()
		defer a.lock.Unlock()

		a.value++
	}()
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()

	return a.value
}


func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()

	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

