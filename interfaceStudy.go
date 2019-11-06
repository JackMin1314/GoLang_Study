/**
* @FileName   : interfaceStudy
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 接口的学习和注意事项的详细说明
* @Time       : 2019/11/5 18:02
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

// 接口也是值("值"是用来保存东西的,东西不止一个，也有可能多个),接口的命名一般约定 er,　r, able 结尾(视情况而定)
// 接口类型 是由一组[方法]签名定义的集合。(0个或多个方法的签名,不包含实现)
// 接口类型的变量 可以保存任何实现了这些方法的值。
// 实现接口不需要显式的声明，只需实现相应方法即可，故没有implements关键字

type I interface { // 接口定义
	M() time.Time // 注意需要指出方法的返回值类型,方法的接收者参数不用写.(为什么不写？你能感觉到什么？可能是因为他不确定你的参数具体是哪个->为了“重载”,详见后文)
} // 如果写了参数我还要用接口干嘛？那不就等同于方法了吗？不写参数怎么实现上面说的实现方法的“重载”?见开篇第一句。

type T struct {
	mytime time.Time
}

// 此方法表示 类型T实现了接口 I ，并成功调用了方法M().同样可以用其他类型去实现接口I 进而调用M，这时候可以重构新的M以达到“函数重载”的目的
func (t T) M() time.Time {
	t.mytime = time.Now()
	return t.mytime
}

func testInterface() {
	var i I = T{time.Now()} // 这里声明变量要用 "var" 不能用":="。因为 “:=” 只能适用于局部变量
	fmt.Println(i.M())      // 实现了接口就可以调相应的方法
}

// 自定义一个float64类型去实现接口I调用M
type MyFL64 float64

func (f MyFL64) M() time.Time {
	fmt.Println(time.Hour)
	return time.Time{}
}

func testInterface2() {
	var i I = MyFL64(3.0)
	i.M()
} // 接口值保存具体底层的类型((f,float64),(t,T))。接口值调用方法时会执行其底层类型(float64 or T)的同名方法M().那如果接口里面的值是nil呢？

// 对于方法的接收者是一个指针的时候，需要内部判断指针是否为空,避免空指针异常.用方法去解决空指针问题，很优雅~
// nil接口值不保存值也不保存具体类型(他不保存M(),因而不认识)， i=nil -> i.M()->产生运行时错误空指针
/* 注意:"nil接口值"不等于"空接口"
nil接口值：接口类型(有M()，不空)定义了一个变量 i ,但是i是空值，没赋值。
空接口：接口为空，接口值不空 -> var i interface{} , i = 22,i = "hello"->空接口可以接受任何类型的值->他就能处理未知类型的值(能接受任何类型的值就是能处理未知类型的值.这个思维要有)
*/

// 访问接口底层的具体值 => 类型断言,注意回顾之前map映射提到的内容读取映射
func testInterface3() {
	var i interface{} = "hello"
	elemStr, check := i.(string) // 采用此方式不会引起恐慌panic
	fmt.Println(elemStr, check)
	i = 15
	elemInt, check := i.(float64) // 0,false ; int -> 15 true
	fmt.Println(elemInt, check)   // elem将保存读取到的值，有为具体值，没有则为类型零值
} // 对于类型较多的情况，可以考虑使用switch case,但这时候的case 后面就是具体的类型了(string ,int,float64)，不是值

func main() {
	testInterface()
	testInterface2()
	testInterface3()
}
