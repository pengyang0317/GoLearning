package main

import (
	"fmt"
	codesnippet "lgo/src/ch2/codeSnippet"
	"sync"
	"sync/atomic"
)

// 子goroutine结束的时候通知主goroutine 通过WaitGroup
func LWaitGroup() {
	defer fmt.Println("Done")
	const numbers = 10
	var wg sync.WaitGroup
	wg.Add(numbers)
	for i := 0; i < numbers; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%d ", i)

		}(i)
	}
	wg.Wait()
}

type mutexMessage struct {
	wg    sync.WaitGroup
	mu    sync.Mutex
	total int64
}

func (m *mutexMessage) add() {
	defer m.wg.Done()
	for i := 0; i < 10000; i++ {
		m.mu.Lock()
		m.total += 1
		m.mu.Unlock()

	}
}
func (m *mutexMessage) sub() {
	defer m.wg.Done()
	for i := 0; i < 10000; i++ {
		m.mu.Lock()
		m.total -= 1
		m.mu.Unlock()
	}
}

// 互斥锁
func LMutex() {
	var m mutexMessage
	m.wg.Add(2)
	go m.add()
	go m.sub()
	m.wg.Wait()
	fmt.Printf("total: %d\n", m.total)
}

type atomicMessage struct {
	total int64
	wg    sync.WaitGroup
}

func (a *atomicMessage) atomicAdd() {
	defer a.wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&a.total, 1)
	}
}
func (a *atomicMessage) atomicSub() {
	defer a.wg.Done()
	for i := 0; i < 10000; i++ {
		atomic.AddInt64(&a.total, -1)
	}
}

// 通过原子操作 实现静态数据的并发安全
func LAtomic() {
	var at atomicMessage
	at.wg.Add(2)
	go at.atomicAdd()
	go at.atomicSub()
	at.wg.Wait()
	fmt.Printf("total: %d\n", at.total)
}

func main() {
	// LWaitGroup()
	// LMutex()
	// LAtomic()
	// codesnippet.LReadWrite()
	// codesnippet.Lchannel()
	// codesnippet.Lchannel2()
	// codesnippet.Lchannel3()
	// codesnippet.Lselect()
	codesnippet.Lcontext()

}
