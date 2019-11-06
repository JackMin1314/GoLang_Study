/**
* @FileName   : StringerReaderImage
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 介绍fmt中包的Stringer使用和IO包Reader
* @Time       : 2019/11/6 12:13
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
)

// fmt 包里面有Stringer是最普遍的接口之一
/*
type Stringer interface {
    String() string
}
*/

type Vertex struct {
	x int
	y int
}

// 可以实现字符串format打印
func (v Vertex) String() string {
	return fmt.Sprintf("v.x = %v; v.y = %v", v.x, v.y)
}

func testStringer() {
	hosts := map[string]Vertex{
		"Point1": {1, 2},
		"Point2": {2, 3},
	}
	for index, listHosts := range hosts {
		fmt.Printf("%v: %v\n", index, listHosts) // 这里输出的其实是Point1: v.x = 1; v.y = 2;Point2..
	}
}

// error类型是一个接口类型。Go 程序使用 error 值来显示表示错误状态. error 类型是一个内置的全局接口.函数成功调用返回error的值为nil,失败为非nil(可以扩展)
/*
type error interface {
	Error() string // error变量可以通过任何可以描述自己的string类型的值来拓展示自己
}
*/
// 额外话:A 的ASCII码值65 ;a 的ASCII码值为97; 空格的ASCII码的值为32; 65+32=97
// 实现 error 的打印
func (v *Vertex) Error() string {
	return fmt.Sprintf("at %v, %v", v.x, v.y)
}
func runError() error {
	return &Vertex{10, 11}
}
func testError() {
	if err := runError(); err != nil {
		fmt.Println(err)
	}
} // 更多错误处理详见errorStudy.go

func main() {
	testStringer()
	testError()
}
