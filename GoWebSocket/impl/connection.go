/**
* @FileName   : impl
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 封装底层的websocket实现多线程安全，添加一些工程化设计
* @Time       : 2020/1/16 20:52
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
	// 当队列满了，writeLoop()会进入ERR，导致底层被关闭，为了使得readLoop()也会关闭
	// 采用的是改写Close(),关闭连接同时关闭closeChan。-->导致readLoop()进入ERR跳出select
	closeChan chan byte

	// 为了解决closeChan被多次关闭的问题，添加一个状态判断
	isClosed bool
	// 为了保证Close()方法的线程安全,在判断的时候需要加锁
	mutex sync.Mutex
}

// 初始化websocket大写开头字母可以被包外引用
func InitConnection(wsConn *websocket.Conn) (conn *Connection,err error){
	conn = &Connection{
		wsConn:wsConn,
		inChan:make(chan []byte, 1000),
		outChan:make(chan []byte, 1000),
		closeChan:make(chan byte, 1),
		isClosed:false,
	}
	// 初始化连接的时候启动一个读协程
	go conn.readLoop()
	// 启动一个写协程
	go conn.writeLoop()
	return
}
// API,支持并发调用,线程安全
func (conn *Connection)ReadMessage() (data []byte,err error) {
	// data = <- conn.inChan
	// 添加当底层长链接出错退出的时候，避免没有数据可取的时候，用户阻塞在inChan。因而需要求用户api报错退出
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan: // closeChan被关闭了
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection)WriteMessage(data []byte) (err error) {
	// conn.outChan <- date
	select {
	case conn.outChan <- data:
	case <- conn.closeChan:
		err = errors.New("connection is closed")

	}
	return
}

// 保证Close方法是线程安全的，且可以重入执行，底层closeChan被执行一次
func (conn *Connection)Close()  {
	// websocket 的close()方法是线程安全的,可重入的Close
	conn.wsConn.Close()

	// 这行代码只执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		// 关闭channel
		close(conn.closeChan)
		// 保证只关闭一次，需要修改下状态
		conn.isClosed = true
	}
	conn.mutex.Unlock()

}

// 内部实现inChan有数据往里面放和outChan往外拿数据;(注意函数开头字母大小写)
// 不停的从websocket长连接上读取数据
func (conn *Connection)readLoop() {
	var (
		data []byte
		err error
	)
	for {
		if _,data,err = conn.wsConn.ReadMessage();err != nil{
			goto ERR
		}
		// 阻塞在这里,等待inChan有空闲的位置；即便是writeLoop报错，导致conn.Close()也不会影响到这个
		// 解决阻塞是用select，添加closeChan是否关闭判断是否阻塞
		select {
		case conn.inChan <- data:
		case <- conn.closeChan:// 当closeChan别关闭的时候,跳出真个函数
			goto ERR

		}
	}
	ERR:
		conn.Close()
}

func (conn *Connection)writeLoop() {
	var (
		data []byte
		err error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <- conn.closeChan:
		    goto ERR
		}
		data = <-conn.outChan
		if conn.wsConn.WriteMessage(websocket.TextMessage,data);err != nil {
			goto ERR
		}
	}
	ERR:
		conn.Close()
	
}