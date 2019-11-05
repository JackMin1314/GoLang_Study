/**
* @FileName   : method and interface
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 方法和函数,注意事项和以指针为接受者的适用情况
* @Time       : 2019/11/5 14:11
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
)

/*写在前面: 什么是面向对象语言oop？什么是面向过程？面向对象特点是什么？封装、多态和继承
C语言、FORTRAN都是面向过程语言,C++、java,python,C#等是面向对象的.java,C#都是单继承;C++是多继承;python都可以
Go语言是oop吗？Go不是完全的面向对象的。Go中没有类和对象的概念。仅支持封装(包,结构体范围),不支持继承和多态->继承和多态是在接口中实现的。
因而个人观点：Go语言从定义上严格来说不算，但是人家可以实现面向对象的功能特性,只是语法方式上有多不同.
*/

// 为结构体类型定义方法,结构体是值类型哦
type std struct {
	name string
	age  int
}

// 这是为结构体类型定义一个方法->特殊的带接受者receiver参数的函数
// 方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。还有就是接收者的类型定义和方法声明必须在同一包内;不能为内建类型声明方法

func (stdp std) showStruct() { // 输出打印结构体
	fmt.Println("真实的结构体内容为: ", stdp)
}

// 尝试修改结构体值并打印
func (stdp std) showStruct1() std {
	// 尝试修改传递过来的st内容，注意函数方法的参数reciver
	stdp.name = "ZhangYueMin"
	stdp.age = 20
	fmt.Println("当前局部的结构体为: ", stdp) // 当前局部结果
	return stdp
}

func (stdp *std) changeStruct() {
	stdp.name = "ChenWang"
	stdp.age = 21
}

func testStructMethod() {
	st1 := std{"CW", 23}
	st2 := std{"ZYM", 21}

	st2.showStruct()  // 原来情况：{ZYM 21}
	st2.showStruct1() // 局部变量打印返回的结果：{ZhangYueMin 20}
	st2.showStruct()  // 现在情况：{ZYM 21}。说明结构体变量的方式指明方法，是修改不了结构体的内容
	//
	st1.showStruct()   // 原来情况：{CW 23}
	st1.changeStruct() // 结构体指针修改内容；不返回结构体。
	st1.showStruct()   // 现在情况：{ChenWang 21}。说明结构体指明方法参数是(接受者为指针)指针的时候,才可以修改内容
}

// 上面的方法只是个带接收者参数的函数。同样可以定义一个简单的正常的函数,参数是结构体变量或者是结构体指针如下
func normalStruct(st std) std {
	st.name = "Chenwang"
	st.age = 10
	return st
}
func testStructMethod2() {
	st := std{"CW", 10}
	temp := normalStruct(st) // 这个只是个普通的函数调用，不过参数可以是结构体变量或者结构体指针(调用要用这个&)
	fmt.Println(temp)
}

// 通过上面的展示。请仔细体会下面的两句话以及样例！
// 以值为接收者的方法被调用testStructMethod()。以值作为参数的函数被调用testStructMethod2()。这两句话不一样！
/*
以值为接收者的方法被调用时，接收者既能为值又能为指针。个人建议不要混用！有讲究的，后面有说到.
var st std
fmt.Println(st.showStruct()) // OK
p := &st
fmt.Println(p.showStruct()) // OK 显然此时编译器会将p.showStruct()-->(*p).showStruct()
*/
/*
接受一个值作为参数的函数必须接受一个指定类型的值(正常的函数参数是变量，你传参就变量，传指针就报错)
var st std
fmt.Println(normalStruct(st)) // OK
fmt.Println(normalStruct(&st)) // 编译错误
*/

// 重要！何时选择指针作为接受者?
/*
1. 为了方法能够修改其接收者指向的值(例如，我要修改结构体里面的内容)
2. 为了可以避免在每次调用方法时复制该值.(时间空间的开销，对于大型的结构体传地址比直接拷贝一份操作更高效)。这里很类似C、C++里面的特点
*/

func main() {
	// testStructMethod()
	// testStructMethod2()
}
