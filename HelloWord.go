package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*func add(x int, y int) int {
//	return x+y
//
//}*/
// 函数和参数定义
func add(x, y int) int {
	return x+y
}
// 函数多返回值
func swap(a, b string) (string,string) {
	return b,a

}
// 函数返回值被命名
func split(sum int) (a,b int) {
	a = sum*2/4
	b = sum-a
	//return a,b  也可以
	return  // 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
}
// var声明一组同类型的变量列表,未初始化默认字符串为空“”，布尔值为false，数值类型为0；
var c,python,java string
// var声明一组不同类型的变量，但是要赋值
// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
var name,age,gender = "Jack",10,true//name ="jack",age=10,gender=true;

/*
同导入语句一样var变量声明也可以使用分组语句块
var (father = "Mr.Green"
	father_age  = 45
	sex         = true)
*/

//if 语句可以在条件表达式前执行一个简单的语句。该语句声明的变量作用域仅在 if 之内。
func mypow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Helloworld()  {
	fmt.Println("Hello World")
	//fmt.Println("This time is ",time.Now())
	fmt.Println(add(3,4))
	fmt.Println(swap("hello","world"))
	fmt.Println(split(15))
	fmt.Println(mynewton(2))
	fmt.Println(mypow(3,2,10),mypow(3,3,20))
}

// 小练习newton法求根
func getabs(x float64) float64 {

	if x < 0 {
		return -x
	}
	if x == 0 || x > 0 {
		return x
	}
	return x

}
func mynewton(x float64) float64 {

	z := float64(1)
	i := 0
	for i < 10 {
		if getabs(z*z-x) < 1e-6 {
			return z
		}
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main(){
	Helloworld()
	// fmt.Println("my favorite number is",rand.Intn(10))

	// i:= 0 也是可以的，但是在函数外面就必须使用var;  :=适用于函数内,不适用于函数外和常量类型const  for语句后面的三个构成部分外没有小括号
	for i:=0;i<3 ;i++  {
		fmt.Println(name,age,gender,i)
	}
	// C 的 while 在 Go 中叫做 for。
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	var ii = 0
	ff := 0.4
	fmt.Println(ii,ff)
	// switch 的 case 无需为常量，且取值不必为整数。自动提供了break语句，除非是fallthrough结束
	switch c := runtime.GOOS;c {
	case "Linux":
		fmt.Println("This OS is ",c)
	case "darwin":
		fmt.Println("This OS is ",c)
	default:
		fmt.Printf("This OS is %s.\n", c)
	}
	// 无值的case相当于if--then--else
	t:=time.Now().Hour()
	switch {
	case t<8:
		fmt.Println("It's too early")
	case t<10:
		fmt.Println("good morning")
	case t < 17:
		fmt.Println("good afternoon")
	default:
		fmt.Println("good night")
	}
	// defer 推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
	// 如果一个函数内有多个defer，先出现的后调用(类似栈)-->defer 栈

}


