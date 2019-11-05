/**
* @FileName   : FunctionVales
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 详细介绍函数，函数值以及闭包
* @Time       : 2019/11/4 18:33
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
)

// 函数也是值--> 函数的参数可以是函数，函数的返回值也可以是函数！！python里面也是允许这样的(高阶函数)，允许函数多返回值且函数参数可以是函数名，语法不同
// 对于其他的编译型语言例如C,C++,C#等是不允许函数多返回值(要想返回多个值，除非使用返回一个结构体或者数组指针或者函数指针等)
// 同时C,C++,C#函数的参数有两种传递方式：值传递,引用传递。Go是值传递的,函数值也是值传递。
func compute(fx func(float64, float64) float64) float64 {
	return fx(3, 4)
}

/*这里compute的参数是"申明一个函数变量" 将“func(float64, float64) float64”看成float64会更好理解就变成func compute(fx float64)float64{..}
fx 是一个值，只不过是函数值，你甚至可以这样去声明并赋值valueFun:= func(x,y,float64)float64{return math.Sqrt(x*x+y*y)}
因而fx func(float64, float64) float64作为compute的参数，后面的float64是函数fx的返回值类型.最最后的一个float64是compute的返回类型必须和fx返回一致
*/
// 为了更好的体会函数值是一个参数，修改上面的compute,添加两个float64参数x,y
func compute2(fx func(float64, float64) float64, x, y float64) float64 {
	return fx(x, y)
}

func addFunc(x, y float64) float64 {
	return x + y
}

func testFunc() {
	fmt.Println(compute(addFunc))
	fmt.Println(compute2(addFunc, 1, 3))
}

// python里面会提到柯里化装饰器：
/*
def show():
    print('zym') #定义一个函数，显示 zym

def logger(m):  #定义一个函数来接收一个参数
    def wrapper():  #这个函数是用来包装我们需要装饰的函数
    #在这里可以进行装饰（前）
        ret = m()   #原有函数的执行，一定不能改动
    #在这里可以装饰 （后）
        return ret
    return wrapper #返回一个函数，调用wrapper包装函数
n = logger(show) #这一部分可见返回函数浅析
print(n())   #zym
*/

// 函数的闭包。上面的compute、compute2特点是函数参数是函数，返回值是一个常见的数据类型.那如果当返回值是一个函数呢？
func compute3() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 这里compute3功能不仅仅是下面的简单的保存每次的记录(sum);还可以起到装饰的作用，装饰不等于侵入语句.具体请自行体会功能的强大。
func testFun2() {
	trans := compute3() // 此时的trans相当于 trans := func(x int)int{return sum},又因为sum存在受限于compute3,被返回的函数给包住了
	for i := 0; i < 4; i++ {
		fmt.Println(trans(i))
	}
}

// 斐波那契数列 0,1,1,2,3,5,8,13,21... (下面这个函数不是递归调用，可以将t,sum,fan看成"全局变量")
func fibonacci() func(int) int {
	sum := 0
	t := 0
	fans := 0
	return func(ans int) int {
		if ans == 2 || ans == 1 {
			sum = 1
			t = 1
			fans = 1
			return 1
		}
		t = sum
		sum = fans
		fans = t + sum
		return fans
	}
}

func testFun3() {
	fibo := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Printf("%d ", fibo(i))
	}

}

func main() {
	testFunc()
	testFun2()
	testFun3()
}
