/**
* @FileName   : myWebDemo
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 用Go搭建一个web服务器,并一步步完成返回静态页面渲染
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
	http.HandleFunc("/index/", sayhelloName) // 这里可以修改url路径和匹配的函数(由fmt.Fprintf()返回http.ResponseWriter)
	err := http.ListenAndServe(":9090", nil) // 修改监听端口，端口占用会抛出异常
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
