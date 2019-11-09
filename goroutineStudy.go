/**
* @FileName   : goroutineStudy
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 这里简单介绍Go里面的并发理念使用和与其他语言比较的不同之处。
goroutine 和 channel 如何使用它们来实现不同的并发模式
* @Time       : 2019/11/9 10:28
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
*/

package main

import (
	"fmt"
	"time"
)

// 在学习goroutine前先了解下并发并行、进程线程、同步异步、多路复用、阻塞(se)非阻塞等概念以及彼此区别，协程的优势.可以参考我之前的文章：
// https://github.com/JackMin1314/Learning_Skill/blob/master/同步异步和阻塞非阻塞.md
// 这一块内容不仅在Go中非常重要，在许多的语言里面都很重要，深入理解对软件开发很重要特别是web端。更多学习参考网上资料
// 这块内容(不仅是Go)也是实际项目中均衡负载优化的不可忽略之处，很影响性能，通常需要结合其他问题分析比如数据库读写，内存磁盘读写，网络请求等

// 线程是操作系统的内核对象。协程，是在应用层模拟的线程，是基于线程的，并不是由操作系统调度的，有程序控制，而且应用程序也没有能力和权限执行 cpu 调度;
// goroutine 协程是由go运行时管理的轻量级线程（不像线程的切换需要额外的开销(内存、寄存器)）,Go的调度器减小了开销且更有效利用CPU缓存,通过使用数量合适的线程并在每一个线程上执行更多的工作来降低操作系统和硬件的负载
// 不要通过共享内存的方式进行通信，而是应该通过通信的方式共享内存。协程更适合于用来实现彼此熟悉的程序组件，如合作式多任务，生成器,迭代器，无限列表和管道。

// 协程的执行:协程执行的代码被扔进一个待执行队列中，由这 n 个线程从队列中拉出来执行。
// 协程的切换:golang对各种io函数封装，内部调用了操作系统的异步io函数(linux的epoll,select;windows的iocp、event)。当异步返回busy,blocking时，golang将现有执行序列压栈，让线程去从队列里拉另一个协程

// 工程上并发通信模型：共享内存(例如锁变量Python) 和 消息channel机制(Go)
func sumChannel(s []int, c chan int) { // 注意到函数是没有返回值的，因为被调用的函数返回时候，goroutine自动结束，返回的函数值被丢弃
	var sum int
	for _, index := range s {
		sum += index
	}
	c <- sum // 将数据塞到信道里
}

func testGoroutineChannel() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int) // 用信道的时候先创建，未指明大小，默认为unbuffer
	go sumChannel(s[len(s)/2:], c)
	go sumChannel(s[:len(s)/2], c)
	x, y := <-c, <-c                                      // 从信道里取数据
	fmt.Printf("x is %v, y is %v, chan is %v\n", x, y, c) // x is 17, y is -5, chan is 0xc000056060
	ch := make(chan int, 2)                               // 创建带有缓冲区大小为2的信道
	ch <- 1
	ch <- 2
	//ch <- -3 // 当信道缓冲区满了，发送方再向其塞数据会阻塞；fatal error:deadlock！ goroutine 1 [chan send]:
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	//fmt.Println(<-ch) // 当缓冲区为空的时候，接收方再向其取数据会阻塞：fatal error:deadlock！ goroutine 1 [chan receive]:
	// 单独的 channel 塞数据是类似队列的方式(现实的管道);而 goroutine 的切换执行序列(函数)是需要判断入栈的
	// 在需要使用缓冲信道时候，如何避免塞取不确定导致阻塞?
	// 只能在发送者地方 close(ch) 表明发送完毕,不在使用这个 channel，接受者通过接收表达式第二个参数判断信道是否关闭以及执行完.v,ok := <-ch;执行完后，ok为false
} // channel 实现了类似锁的功能，并保证了所有 goroutine 完成后 main() 才返回。

// C语言或UNIX中，select()函数用来监控一组描述符，该机制常被用于实现高并发的socket服务器程序.Go支持语言级别select关键字，用于处理异步IO问题。（python没有上升到语言，只是底层封装）
//  select控制结构时候一能在 Channel 上进行非阻塞的收发操作，二能是 select 在遇到多个 Channel 同时响应时能够随机挑选 case 执行。
func testSelect() {
	ch := make(chan int, 2) // 给channel的buffer长度为2
	ch <- 6
	// select 用default实现非阻塞的收发
	select {
	case i := <-ch:
		fmt.Println(i)
	default:
		fmt.Println("default") // 如果注释了ch <- 6 会直接执行default,不阻塞goroutine
	}
	// select 多匹配(多个分支准备好)时随机执行一个
	ch1 := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) { // time.Tick会返回一个chan Time(如果不shutdown的话将无法被garbage collector回收)
			ch1 <- 0
		}
	}() // 并发的方式调用匿名函数func
	for {
		select {
		case <-ch1:
			fmt.Println("case1")
		case <-ch1:
			fmt.Println("case2")
		}
	} // 这个是死循环,里面 select 随机执行 case
}
func main() {
	// testGoroutineChannel()
	testSelect()
}

/*协程是基于线程的，内部封装了系统的异步函数(不同系统不一样),根据异步函数的返回值判断，然后将当前的执行序列(要执行的函数)压栈，然后利用线程从队列里面拉其他的协程。见代码testGoroutineChannel()（而channel管道则是类似队列的结构）
(为什么放在队列里？按顺序执行,因为协程是基于线程的，对于单个线程而言，多核cpu或者说系统不负责调度，上下文切换由程序员控制，减少额外开销，又因为程序代码是串行执行的函数调用，因而在一个线程里协程串行执行。
*/
