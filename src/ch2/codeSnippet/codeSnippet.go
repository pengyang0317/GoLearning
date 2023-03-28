package codesnippet

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type rwMessage struct {
	total int64
	wg    sync.WaitGroup
	mu    sync.RWMutex
}

func (rw *rwMessage) handlerWrite() {
	defer rw.wg.Done()
	for {
		rw.mu.Lock()
		rw.total = 18
		time.Sleep(5 * time.Second)
		rw.mu.Unlock()
		time.Sleep(3 * time.Second)
	}

}
func (rw *rwMessage) handlerRead() {
	defer rw.wg.Done()
	for {
		rw.mu.RLock()
		fmt.Printf("total: %d\n", rw.total)
		time.Sleep(1 * time.Second)
		rw.mu.RUnlock()
	}

}

// 通过读写锁 实现静态数据的并发安全  读写锁的效率比互斥锁高
func LReadWrite() {
	rw := rwMessage{}
	rw.wg.Add(3)
	go rw.handlerWrite()
	go rw.handlerRead()
	go rw.handlerRead()
	rw.wg.Wait()
	fmt.Printf("total: %d\n", rw.total)
}

// channel 通信 channel默认是阻塞的， 双向通道
func Lchannel() {
	var ch = make(chan int, 2)

	go func() {
		// for {
		// 	i := <-ch
		// 	fmt.Println(i)
		// }
		for i := range ch {
			fmt.Println(i)
		}
		fmt.Printf("all done")
	}()

	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	time.Sleep(1 * time.Second)

}

// 单向通道
func Lchannel2() {
	var ch = make(chan int, 2)
	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}(ch)

	go func(in <-chan int) {
		for i := range in {
			fmt.Println(i)
		}
		fmt.Printf("all done")
	}(ch)

	time.Sleep(1 * time.Second)

}

// 利用channel 交替打印
func Lchannel3() {
	var chInt, chStr = make(chan bool), make(chan bool)

	go func() {
		i := 1
		for {
			<-chInt
			fmt.Printf("%d%d", i, i+1)
			i += 2
			chStr <- true
		}
	}()
	go func() {
		i := 0
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		for {
			<-chStr
			if i >= len(str) {
				return
			}

			fmt.Printf(str[i : i+2])
			i += 2
			chInt <- true
		}
	}()
	chInt <- true
	time.Sleep(10 * time.Second)
}

// 利用select 语句实现多路复用
func Lselect() {

	var g1 = func(ch chan struct{}) {
		time.Sleep(time.Second)
		ch <- struct{}{}
	}
	var g2 = func(ch chan struct{}) {
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}

	var g1c, g2c = make(chan struct{}), make(chan struct{})

	go g1(g1c)
	go g2(g2c)

	var tc = time.NewTimer(5 * time.Second)

	for {
		select {
		case <-g1c:
			fmt.Printf("g1 done\n")
		case <-g2c:
			fmt.Printf("g2 done\n")
		case <-tc.C:
			fmt.Printf("default\n")
			return
		}
	}
}


//用context 控制多个goroutine的退出
func Lcontext() {
	var ctx, cancel = context.WithCancel(context.Background())
	var g1 = func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("g1 done\n")
				return
			default:
				fmt.Printf("g1 running\n")
				time.Sleep(time.Second)
			}
		}
	}

	var g2 = func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("g2 done\n")
				return
			default:
				fmt.Printf("g2 running\n")
				time.Sleep(time.Second)
			}
		}
	}

	go g1(ctx)
	go g2(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	fmt.Printf("监控完成\n")
}