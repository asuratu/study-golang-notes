package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex // 这里的Mutex就是互斥锁
}

// 数据的写入要加锁
func (a *atomicInt) increment() {
	fmt.Println("safe increment")
	// 要保证一个代码块的原子性，就用一个匿名函数包装一下
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
}

// 数据的读取也要加锁
func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	// 线程安全的
	//atomic.AddInt32()
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}

// 查看数据冲突
// go run -race atomic.go
/*==================
WARNING: DATA RACE
Read at 0x00c00013c018 by main goroutine:
main.main()
/Users/asura/Code/go/taobao/01、basic/atomic/atomic.go:27 +0xee

Previous write at 0x00c00013c018 by goroutine 7:
main.(*atomicInt).increment()
/Users/asura/Code/go/taobao/01、basic/atomic/atomic.go:11 +0x45
main.main.func1()
/Users/asura/Code/go/taobao/01、basic/atomic/atomic.go:24 +0x2e

Goroutine 7 (finished) created at:
main.main()
/Users/asura/Code/go/taobao/01、basic/atomic/atomic.go:23 +0xd8
==================*/

// Read at 27 Previous write at 11
// 意思27行正在读的时候，有另外一个人在写，所以冲突
