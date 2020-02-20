/**
* @FileName   : Server
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description:
* @Time       : 2019/12/23 10:48
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"GoWebSocket/impl"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upgrade = websocket.Upgrader{
		// 允许跨域
		CheckOrigin : func(r *http.Request)bool {
			return true
		},
	}
)
// 回调函数
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err error
		// msgType int
		data []byte
		conn *impl.Connection
	)
	// w.Write([]byte("hello"))
	if wsConn, err = upgrade.Upgrade(w,r,nil); err != nil{
		return // 出错默认中断调用
	}
	// // ReadMessage()、WriteMessage()是线程不安全的，因而需要进行对其封装。放在channel里面，Go的channel是线程安全的;
	// // 得到长链接时候
	// for{
	// 	// wensocket支持的消息类型:text;binary
	// 	if _, data,err = conn.ReadMessage();err != nil {
	// 		goto ERR
	// 	}
	// 	if err = conn.WriteMessage(websocket.TextMessage,data); err != nil {
	// 		goto ERR
	// 	}
	//
	// }
	// ERR:
	// 	// 关闭连接
	// 	conn.Close()
	if conn,err = impl.InitConnection(wsConn);err !=nil{
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for  {

			if err = conn.WriteMessage([]byte("heartbeat!!!"));err != nil {
				return
			}
			time.Sleep(1*time.Second)
		}
	}()

	for {
		if data,err = conn.ReadMessage();err !=nil {
			goto ERR
		}
		if err = conn.WriteMessage(data);err !=nil {
			goto ERR
		}
	}
	ERR:
		// TODO 关闭连接的操作
		conn.Close()

}

func main() {
	// http://localhost:7777/ws
	http.HandleFunc("/ws",wsHandler)
	http.ListenAndServe("0.0.0.0:7777",nil)
}
