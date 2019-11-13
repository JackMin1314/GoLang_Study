/**
* @FileName   : myWebDemo
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 用Go搭建一个web服务器,根据源码详解如何让web运行,http包运行大体流程和注意的问题
* @Time       : 2019/11/12 17:49
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析参数,然后才可以调用r.Form，默认是不会解析的
	fmt.Println(r.Form) // 这些信息时候输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"]) // 后去url的参数
	fmt.Println(r.Host)             // 获取请求主机
	fmt.Println(r.UserAgent())      // 获取user-agent
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "hello Jack!") // 这个写入到w的是输出到客户端的
}

// 定义自己的路由
func main() {
	// 注册了请求/index/的路由规则，为了当请求uri为"/index/"，路由就会转到函数sayhelloName
	http.HandleFunc("/index/", sayhelloName) // 这里可以修改url路径和匹配的函数(由fmt.Fprintf()返回http.ResponseWriter)
	err := http.ListenAndServe(":9090", nil) // 修改监听端口，端口占用会抛出异常.
	// ListenAndServer底层是先初始化一个Server结构体对象，调用结构体的ListenAndServe()方法！然后调用了net.Listen("tcp", addr)，也就是底层用TCP协议搭建了一个服务，然后监控我们设置的端口。
	// 然后返回Server()方法(属于结构体Server的方法)处理客户端的请求信息，有个for死循环{},里面先会调用net包里面Listener接口的Accept()方法能接受(服务未关闭)，
	// 则调Server的c, err := srv.newConn(rw)来创建新的conn(为http包的conn结构体指针，里面包括了请求数据rw还有后面的server操作)
	// 最后创建协程goroutine去跑 go c.serve() ---->高并发的体现 （用户的每一次请求都是在一个新的goroutine去服务，相互不影响）
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

/*详见https://github.com/JackMin1314/build-web-application-with-golang/blob/master/zh/03.3.md
浅谈http包执行大体流程
1. 创建Listen Socket。监听指定的端口等待客户端请求的到来，
2. Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与客户端通信
3. 处理客户端的请求. 首先从Client Socket读取HTTP请求的协议头, 如果是POST方法, 还可能要读取客户端提交的数据, 然后交给相应的handler处理请求, handler处理完毕准备好客户端需要的数据, 通过Client Socket写给客户端。
流程需要解决的问题
1).如何监听端口？在这之前还需要解决什么问题？
Socket(创建socket并初始化)->Bind(进程和端口绑定)->Listen(监听端口)->Listen Socket   如果收到请求且Accept()->Client Socket
2).如何接收客户端请求？
见main里面的注释关于ListenAndServe()剖析
3).如何分配handler？
ListenAndServe()的第二个参数为handle，我们给定nil表示默认使用DefaultServeMux，调用ServeHTTP方法，这个方法内部其实就是调用sayhelloName本身，通过写入response的信息反馈到客户端。见sayhelloName()
*/
