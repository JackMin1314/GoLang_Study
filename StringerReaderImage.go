/**
* @FileName   : StringerReaderImage
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 介绍fmt中包的Stringer使用和IO包Reader接口以及image接口
* @Time       : 2019/11/6 12:13
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
	"io"
	"strings"
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

// io 包指定了 io.Reader 接口，它表示从数据流的末尾读取字符。很多其他的标准库包含该接口，文件、网络连接、压缩和加密
// func (T) Read(b []byte) (n int, err error) 数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
func testReader() {
	s := strings.NewReader("This is Reader!") // 返回的是*Reader;"func NewReader(s string) *Reader { return &Reader{s, 0, -1} }"
	/*type Reader struct {
		s        string
		i        int64 // current reading index
		prevRune int   // index of previous rune; or < 0
	}*/
	fmt.Printf("%t,%v\n", s, s) // 输出了 &{This is Reader! 0 -1}
	b := make([]byte, 8)        // make主要是创建Slice,Map,Channal。这里返回的b不是指针，而是具体类型定义的切片.等价于b:=[8]byte
	// var b [8]byte是错误的，这个创建的是byte数组
	for {
		n, err := s.Read(b) // 遇到数据流末尾时候返回一个io.EOF   End Of File
		if err != io.EOF {
			fmt.Println(n, err, b[:n])
			fmt.Printf("%q,%s\n", b[:n], b[:n]) // %q 字符串带双引号，且安全打印b[:n]包括自动转义
		} else {
			break
		}
	}
}

// Image 包也定义了接口image
/*
type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
color.Model  color.Color
*/
func main() {
	// testStringer()
	// testError()
	testReader()
}
