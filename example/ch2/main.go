package main

import (
	"fmt"
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
	//这个优化的思路是将对 fmt.Println 的调用从 goroutine 中移除，改为在主线程中输出。
	//这样做可以避免多个 goroutine 同时向标准输出写入数据，从而减少竞争和锁的使用，提高并发性能。
	fmt.Println()
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
	// 优化的思路是将互斥锁 mutexMessage 中的 WaitGroup 和 Mutex 分别传递给 add 和 sub 方法，而不是在 mutexMessage 中定义。这样可以避免在 mutexMessage 中定义 WaitGroup 和 Mutex 带来的额外开销。

	// 另外，我们还将 add 和 sub 方法改为接收 WaitGroup 和 Mutex 作为参数，而不是在方法内部创建。这样可以避免在每次调用 add 和 sub 方法时创建 WaitGroup 和 Mutex 带来的额外开销。
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
	LMutex()
	// LAtomic()
	// codesnippet.LReadWrite()
	// codesnippet.Lchannel()
	// codesnippet.Lchannel2()
	// codesnippet.Lchannel3()
	// codesnippet.Lselect()
	// codesnippet.Lcontext()

}
