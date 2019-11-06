/**
* @FileName   : errorStudy
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: Go 语言错误以及异常处理error接口详解
* @Time       : 2019/11/6 17:46
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

// Go1.13 版本以及修改了 errors包，增加了几个函数(Is/As/Unwrap,再次包装和识别处理)，用于增强error的功能。且抛弃了GOPATH 和vendor特性，用module管理包。
// error错误跟异常不是完全相同的.Go 语言不支持传统的 try…catch…finally 这种处理。
// 异常：Go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后处理
// 错误:Go 语言为错误处理定义了一个标准模式，即 error 接口,包含了一个 Error() 方法，用于返回错误消息.
/*
type error interface{
     Error() string  // 适用于一个字符串就能说清的情况
}
*/
/* 大多数函数要想返回错误，通常采用返回两个参数;将错误类型作为第二个参数返回，然后根据错误类型进行判断，函数调用正常err为nil.
n, err := Foo(0)
if err != nil {
// 错误处理
} else {
// 使用返回值 n
}
*/

/* errors 包里面定义了内置变量未导出
type errorString struct{
      s string
}
*/

/*而对于Error()，返回的是string. 那么返回的error类型呢？见下面
func (e *errorString)Error() string{
      return e.s
}
*/

/* 返回error类型。 使用 New 函数创建出来的 error 类型实际上是 errors 包里未导出的 errorString，包含唯一字段s，实现了唯一方法Error() string
func New(text string) error {
      return &errorString{text}
}
*/

func main() {

}
