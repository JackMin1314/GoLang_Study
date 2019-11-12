/**
* @FileName   : syncMutex
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: sync包提供了基本的并发编程的同步原语锁(高级的是信道与goroutine的使用使用);以及如何进行对共享内存访问控制
* @Time       : 2019/11/12 11:27
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// 在python、java里面开启多线程的时候对于全局变量操作(准确说是共享内存的修改)需要提前加锁，然后释放锁两个操作，来避免多线程导致共享变量被修改的不确定性情况
// 锁(Mutex)其实是一种并发编程中的同步原语（Synchronization Primitives）,目的就是为了解决多个线程或者 Goroutine 在访问同一片内存时不会出现混乱的问题.
// Go 标准库中提供了 sync.Mutex 互斥锁类型保证每次只有一个 Go 程能够访问一个共享的变量

/*
//Mutex 是一个互斥锁结构体在sync包里面 8个字节
//Mutex 可作为其它结构的一部分来创建；Mutex 的零值即为已解锁的互斥体。
type Mutex struct {
    		state int32 // 当前互斥锁状态 (四个字节 = 当前等待互斥锁被释放go程总数目+最低三位状态位mutexLocked、mutexWoken 和 mutexStarving)
    		sema  uint32 // 控制锁状态的信号量
    	}
//默认初始化的state都是0，当互斥锁被锁定时mutexLocked被设置成1,当互斥锁被正常模式下唤醒则mutexWoken置1,当互斥锁进入状态时候mutexStarving为1


//A Locker represents an object that can be locked and unlocked.

    	//Locker 表示可被锁定并解锁的对象。
type Locker interface {
    		Lock() // 用于获取锁sync.Mutex
    		Unlock() // 用于释放锁
    	}
*/
// 在Lock()和Unlock()之间的代码段称为资源的临界区(critical section)，在这一区间内的代码是严格被Lock()保护的，是线程安全的;注意Mutex还是RWMutex都不会和goroutine进行关联
type SafeCount struct {
	v   map[string]int // PS: map 本身并不是并发安全的
	mux sync.Mutex
}

func (c *SafeCount) Inc(key string) {
	c.mux.Lock()
	c.v[key]++ // Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.mux.Unlock()
}
func (c *SafeCount) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key] // 用 defer 语句来保证互斥锁一定会被解锁
}
func main() {
	c := SafeCount{v: make(map[string]int)} // 注意这种写法初始化创建对象
	for i := 0; i < 10; i++ {
		c.Inc("Test")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("Test"))
}
