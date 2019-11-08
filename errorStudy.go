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

import (
	"errors"
	"fmt"
	"time"
)

// Go1.13 版本以及修改了 errors包，增加了几个函数(Is/As/Unwrap,再次包装和识别处理)，用于增强error的功能。且抛弃了GOPATH 和vendor特性，用module管理包。
// error错误跟异常不是完全相同的.Go 语言不支持传统的 try…catch…finally 这种处理。
// 异常：Go 中可以抛出一个 panic 的异常，然后在 defer 中通过 recover 捕获这个异常，然后处理.(通常是非业务逻辑,未预料的)
// 错误: Go 语言为错误处理定义了一个标准模式，即 error 接口,包含了一个 Error() 方法，用于返回错误消息.(代码需要考虑的问题例如分母为0的问题)

/* 大多数函数要想返回错误，通常采用返回两个参数;将错误类型作为第二个参数返回，然后根据错误类型进行判断，函数调用正常err为nil.
n, err := Foo(0)
if err != nil {
// 错误处理
} else {
// 使用返回值 n
}
*/

/*
type error interface{
     Error() string  // 适用于一个字符串就能说清的情况
}
*/

/* errors 包里面定义了内置变量未导出
type errorString struct{
      s string
}
*/

/*而对于Error()，返回的是string. 那么实现返回的error类型呢？见下面
func (e *errorString)Error() string{
      return e.s
}
*/

/* 返回error类型。 使用 New 函数创建出来的 error 类型实际上是 errors 包里未导出的 errorString，包含唯一字段s，实现了唯一方法Error() string
func New(text string) error {
      return &errorString{text}
}
*/

// 1.New返回格式化为给定文本的错误。即使文本是相同的，每个对New的调用都会返回一个不同的错误值
func testErrorNew() {
	err := errors.New("this is error test for New")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("err is nil")
	}
}

// 2.fmt 的 Errorf 接口可以简单的实现可描述错误的信息提示,1.13以后支持%w谓词应用于错误参数:fmt.Errorf("... %w ...", ..., err, ...)
func testFmtErrorf() {
	name, id := "lin", 15
	err := fmt.Errorf("user %q (id %d) not found", name, id) // 当遇到问题的时候可描述性的返回error
	if err != nil {
		fmt.Print(err)
	}
}

// 3.自定义结构体,实现自定义错误类型
type myErrType struct {
	When time.Time
	What string
}

// 因为fmt包里面定义了error接口存有Error()方法(见前面StringReaderImage的文章后面),这里只需要实现Error()方法即可，搭配Sprintf,返回什么样的数据
func (myerr myErrType) Error() string {
	return fmt.Sprintf("at %v: %v", myerr.When, myerr.What)
}
func getArea(width, length float64) (float64, error) {
	errorInfo := ""
	if width < 0 && length < 0 {
		errorInfo = fmt.Sprintf("Error:长度:%v, 宽度:%v, 均为负数", width, length)

	} else if length < 0 {
		errorInfo = fmt.Sprintf("Error:长度:%v, 出现了负数", length)
	} else if width < 0 {
		errorInfo = fmt.Sprintf("Error:宽度:%v, 出现了负数", width)
	}

	if errorInfo != "" {
		return 0, myErrType{time.Now(), errorInfo}
	} else {
		return width * length, nil
	}
} // 你可能觉得这样写很繁琐不够简洁,当长度或宽度有一个小于0就报错即可。何必这么复杂(优雅)?
// 但真实上并不是这样的，实际开发中当业务逻辑复杂，代码量增加的时候，仅仅是返回有错误是不够具体的，更因该给出描述性具体那块出错，附加信息
// 的确Go的错误处理相对其他语言可能会让代码量增加了，但是他提供了更大的错误处理自由度和细节，这是很有必要的.这种错误处理可以方便用户，也可以为后期开发人员维护快速定位。
func testMyErrType() {
	ans, err := getArea(-7.0, 10.0) // 尝试修改数值
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ans)
	}
}

// 4.Go 1.13 新增了errors一些特性,来一起玩套娃吧!
/*注意比较函数或者方法的参数和返回值
func As(err error, target interface{}) bool
func Is(err, target error) bool
func New(text string) error
func Unwrap(err error) error
*/
// fmt.Errorf()新加了参数%w用来返回一个被包装的error,多次调用实现多重包装
func someNewError() {
	// 套--> errors.New()
	err0 := errors.New("this is err0 new")
	err1 := fmt.Errorf("err1:(%w)", err0) // 每层都可以再添加其他信息
	err2 := fmt.Errorf("err2:(%w)", err1)
	fmt.Println(err2) // 输出了 err2:(err1:(this is err0 new))
	// 脱--> errors.Unwrap()
	fmt.Println(errors.Unwrap(err2))                               // err1:(this is err0 new)
	fmt.Println(errors.Unwrap(errors.Unwrap(err2)))                // this is err0 new
	fmt.Println(errors.Unwrap(errors.Unwrap(errors.Unwrap(err2)))) // <nil>,脱过头了 (可以判断返回结果是否nil，全部去掉,见下)

	// 循环全部输出error信息。
	var tmp = err2
	for tmp != nil {
		fmt.Println(tmp)
		tmp = errors.Unwrap(tmp)
	}
	// 假如现在我拿到了一个包装的err3，我怎么确定他来源于底层的err呢？
	// (每去掉一层就比较类型是否相等)->errors.As();提取指定类型的错误
	// (每去掉一层就比较值是否相等)->errors.Is();是否包含指定错误
	fmt.Println(errors.Is(err2, err1)) // true
	fmt.Println(errors.Is(err2, err0)) // true

	var myerrPtr *myErrType
	err3 := fmt.Errorf("test for errors.As %w", &myErrType{time.Now(), "myErrType"})
	fmt.Println(errors.As(err3, &myerrPtr)) // true;判断类型是否相同，并提取第一个符合目标类型的错误赋给后面的第二个参数返回。
	fmt.Println(myerrPtr)                   // 体会到为什么errors.As()第二个参数为什么是非空指针了吗？如果是空指针或者对象则无法保存符合类型的错误

}

func main() {
	// testErrorNew()
	// testFmtErrorf()
	// testMyErrType()
	someNewError()
}
