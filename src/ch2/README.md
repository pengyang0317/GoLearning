
# go 语言中的并发编程

[TOC]

完整代码地址： <https://github1s.com/pengyang0317/GoLearning/blob/main/src/ch2/main.go>

### 1.协程提高CPU利用率

多进程、多线程已经提高了系统的并发能力，但是在当今互联网高并发场景下，为每个任务都创建一个线程是不现实的，因为会消耗大量的内存 (进程虚拟内存会占用 4GB [32 位操作系统], 而线程也要大约 4MB)。

大量的进程 / 线程出现了新的问题

- 高内存占用
- 调度的高消耗 CPU

然后工程师们就发现，其实一个线程分为 “内核态 “线程和” 用户态 “线程。

一个 “用户态线程” 必须要绑定一个 “内核态线程”，但是 CPU 并不知道有 “用户态线程” 的存在，它只知道它运行的是一个 “内核态线程”(Linux 的 PCB 进程控制块)。

这样，我们再去细化去分类一下，内核线程依然叫 “线程 (thread)”，用户线程叫 “协程 (co-routine)”。

协程跟线程是有区别的，线程由 CPU 调度是抢占式的，协程由用户态调度是协作式的，一个协程让出 CPU 后，才执行下一个协程

### 2.协程 goroutine

Go 为了提供更容易使用的并发方法，使用了 goroutine 和 channel。goroutine 来自协程的概念，让一组可复用的函数运行在一组线程之上，即使有协程阻塞，该线程的其他协程也可以被 runtime 调度，转移到其他可运行的线程上。最关键的是，程序员看不到这些底层的细节，这就降低了编程的难度，提供了更容易的并发。
Go 中，协程被称为 goroutine，它非常轻量，一个 goroutine 只占几 KB，并且这几 KB 就足够 goroutine 运行完，这就能在有限的内存空间内支持大量 goroutine，支持了更多的并发。虽然一个 goroutine 的栈只占几 KB，但实际是可伸缩的，如果需要更多内容，runtime 会自动为 goroutine 分配。

Goroutine 特点：

- 占用内存更小（几 kb）
- 调度更灵活 (runtime 调度)

### 3.go的gmp调度原理

这个内容太多了！  找百度 和 B站吧！

### 4.go通过waitgroup等待协程的执行

在 Go 语言中，sync.WaitGroup 是一个同步原语，用于等待一组 goroutine 完成它们的工作。它的作用是帮助管理多个 goroutine，等待它们全部完成后再执行下一步操作。

- sync.WaitGroup 提供了三个方法：Add、Done 和 Wait。

- Add(n int): 让 WaitGroup 的计数器加上 n，表示需要等待 n 个 goroutine 完成工作。
- Done(): 让 WaitGroup 的计数器减去 1，表示一个 goroutine 完成了工作。
- Wait(): 阻塞当前 goroutine，直到 WaitGroup 的计数器归零。

```
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

```

### 5.通过mutex和atomic完成全局变量的原子操作

在 Go 语言中，可以使用 sync.Mutex 和 sync/atomic 包来实现全局变量的原子操作。

sync.Mutex 是一种互斥锁，用于保护共享资源不被并发访问。在访问共享资源前，我们使用 Mutex.Lock() 方法获取锁，访问完毕后使用 Mutex.Unlock() 方法释放锁。这样可以保证同一时刻只有一个 goroutine 可以访问共享资源，从而避免数据竞争和并发访问问题。

```
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
```

另一种实现全局变量的原子操作的方法是使用 sync/atomic 包。sync/atomic 包提供了一些原子操作函数，例如 atomic.AddInt32、atomic.LoadInt32 和 atomic.StoreInt32 等，可以实现对 32 位整数类型的原子操作。

```
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
//优化点跟上面一样

```

### 6.RWMutex读写锁

RWMutex是Go语言标准库sync包中提供的读写锁，用于保护共享资源在多个协程间的读写操作。RWMutex主要包含两种锁：读锁和写锁。

读锁：多个协程可以同时获取读锁进行读操作，但是不能同时获取写锁。

写锁：只能有一个协程获取写锁进行写操作，此时不能同时获取读锁或写锁。

RWMutex的实现是基于互斥锁Mutex和条件变量Cond的，读锁和写锁都是通过锁定互斥锁来实现。当有多个协程同时请求读锁时，只要没有协程请求写锁，则这些协程可以同时获取读锁，否则它们会被阻塞，直到写锁被释放。当有一个协程请求写锁时，如果当前没有协程持有读锁或写锁，则这个协程可以获取写锁；否则它们会被阻塞，直到读锁或写锁被释放。

使用RWMutex可以提高并发程序的性能，因为读操作可以并发执行，写操作可以串行执行。但是需要注意的是，当读操作非常频繁时，可能会导致写操作一直被阻塞，从而影响程序的性能。因此，在实际使用中需要根据具体场景选择合适的锁机制。

```
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
```

### 7.通过channel进行goroutine之间的通信

通过channel进行goroutine之间的通信是一种常见的并发编程方式，可以用于实现协程之间的数据传输和同步。

通过channel进行通信的基本操作包括：

创建channel：可以使用make函数创建一个新的channel，指定缓存大小，例如：ch := make(chan int, 10)。

发送数据到channel：使用<-符号将数据发送到channel中，例如：ch <- 20。

从channel接收数据：使用<-符号从channel中读取数据，例如：num := <-ch。

关闭channel：使用close函数关闭channel，表示不再向channel中发送数据。

通过channel进行通信的常见应用场景包括：

线程之间的数据传输：将数据发送到channel中，让其他线程从channel中读取。

线程之间的同步：使用无缓冲的channel实现同步，例如：一个线程向channel中发送数据，另一个线程从channel中读取数据。

控制goroutine的执行顺序：使用channel控制不同goroutine的执行顺序，例如：一个goroutine只有在另一个goroutine发送了信号后才能执行。

需要注意的是，当channel被关闭后，仍然可以从中读取数据，但是读取的值为该channel类型的零值。如果尝试向已关闭的channel中发送数据，则会导致panic异常。因此，在使用channel进行通信时，需要注意控制channel的关闭时机，以避免出现异常。

```
// channel 通信 channel默认是阻塞的， 双向通道
func Lchannel() {
 var ch = make(chan int, 2)

 go func() {
  // for {
  //  i := <-ch
  //  fmt.Println(i)
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
```

### 8.有缓冲channel和无缓冲channel的应用场景

有缓冲channel和无缓冲channel的应用场景不同。

无缓冲channel：在创建时没有指定缓冲区大小，也称为同步channel。无缓冲channel的特点是只有在发送和接收都准备好时才能进行数据交换。如果没有接收者，发送者就会一直阻塞；如果没有发送者，接收者也会一直阻塞。无缓冲channel适用于控制goroutine的执行顺序和实现goroutine之间的同步。

常见的应用场景包括：

控制goroutine的执行顺序：通过无缓冲channel实现信号量机制，一个goroutine只有在另一个goroutine发送信号后才能执行。

实现goroutine之间的同步：在多个goroutine之间共享数据时，可以使用无缓冲channel实现同步，确保数据的正确性。

避免竞态条件：当多个goroutine同时访问共享资源时，使用无缓冲channel可以避免竞态条件的发生。

有缓冲channel：在创建时指定缓冲区大小，也称为异步channel。有缓冲channel的特点是在缓冲区未满时可以进行数据交换，如果缓冲区已满，则发送者会阻塞。如果缓冲区为空，则接收者会阻塞。有缓冲channel适用于实现goroutine之间的数据传输。

常见的应用场景包括：

实现生产者-消费者模式：使用有缓冲channel作为生产者和消费者之间的缓冲区，生产者将数据发送到缓冲区，消费者从缓冲区中读取数据。

控制goroutine的并发数量：使用有缓冲channel实现信号量机制，限制同时执行的goroutine数量。

### 9. forrange对channel进行遍历

在Go语言中，使用for range语句可以对channel进行遍历。for range语句会不断从channel中读取值，直到channel被关闭。

语法格式如下：

for val := range ch {
    // 处理val
}
其中，ch是一个channel，val是从channel中读取的值。如果channel未关闭，则for循环会一直阻塞等待新的数据；如果channel已关闭并且缓冲区中没有数据，则循环会结束。

for range语句的优点是可以自动判断channel是否关闭，避免了手动判断channel是否关闭的问题。同时，如果channel中有多个值，for range语句会依次读取并处理这些值，不需要手动遍历channel。

需要注意的是，在使用for range遍历channel时，如果channel中的值未被完全消费，则程序会阻塞。因此，在编写程序时需要确保channel中的值能够被完全消费或者通过关闭channel来终止遍历。同时，如果channel未被关闭，则for range语句会一直阻塞等待新的数据，可能会导致程序陷入死循环或者出现其他问题。

### 10.单向channel

单向channel是指只能进行发送或者接收操作的channel，分为发送channel和接收channel。发送channel只能用于发送数据，不能用于接收数据；接收channel只能用于接收数据，不能用于发送数据。单向channel的定义方式如下：

var sendCh chan<- int // 发送channel，只能发送int类型数据
var recvCh <-chan int // 接收channel，只能接收int类型数据

```
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
```

### 11.通过channel实现交叉打印

通过channel实现交叉打印12AB34CD......

```
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

```

### 12.监控goroutine的执行

在Go语言中，可以使用select语句来监控goroutine的执行情况，同时可以处理多个channel的读写操作。

select语句的基本语法如下：

select {
case <-ch1:
    // 处理ch1读取的数据
case data := <-ch2:
    // 处理ch2读取的数据
case ch3 <- val:
    // 向ch3写入数据
default:
    // 如果所有channel都没有数据，则执行default语句块
}
在select语句中，可以同时监控多个channel的读写操作，当其中一个channel准备好时，就会执行对应的语句块。如果多个channel都准备好了，则会随机选择一个执行。

```
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

 //在这段代码中，两个goroutine g1 和 g2 分别在睡眠 1 秒和 2 秒后向其对应的通道中发送一个空的结构体 struct{}。主函数中使用了一个 select 语句来等待两个goroutine的完成情况。如果两个goroutine都完成了，则打印 "g1 done" 和 "g2 done"；如果超时了，则打印 "default"。

 //这段代码的主要问题在于使用了无缓冲通道，因此在 g1 和 g2 运行之前，主函数会一直阻塞等待通道的读取。这会导致程序的执行效率较低，因为 g1 和 g2 的执行时间很短，但是主函数需要等待很长时间才能得到结果。

 //为了优化这段代码，可以使用带缓冲通道来代替无缓冲通道。这样，在 g1 和 g2 运行之前，主函数可以先向带缓冲通道中写入一个空的结构体 struct{}，从而避免了阻塞等待的问题
 var g1c, g2c = make(chan struct{}, 1), make(chan struct{}, 1)

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

```

### 13.通过context解决goroutine的信息传递

在Go语言中，可以使用context包来解决goroutine之间的信息传递问题。context包提供了一种机制，用于在goroutine之间传递上下文信息，并提供了在超时、取消等情况下协调goroutine的方法。

context包中提供了context.Context类型，表示一个上下文信息。在一个goroutine中创建一个Context对象，并将其传递给其他goroutine，其他goroutine可以使用该对象来获取上下文信息，并进行相应的处理。

```
// 用context 控制多个goroutine的退出
func Lcontext() {
 //在程序中，首先使用context.Background()函数创建一个根上下文对象，并使用context.WithCancel()函数创建一个可取消的上下文对象。然后，创建两个goroutine
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
 //在主函数中，等待一段时间后，调用cancel()函数取消上下文对象。程序退出。
 cancel()
 fmt.Printf("监控完成\n")
}
```
