package main

import (
	"fmt"
)

// 与 C 不同, Go 没有指针运算。默认空指针为nil
func ptrLearn(){
	var p *int
	i := 22
	p = &i // 取地址符号
	fmt.Println(p) // *p 输出为22; p 输出为保存i的地址
	fmt.Println(&i)
}

func structDefine(){
	type Vertex struct {
		X int
		Y int
	}
	// 如果没有初始化值，默认为各自类型的空值 int 为 0
	v := Vertex{1,2}
	fmt.Println(v)
	p := &v
	p.X = 10 // 语言允许隐式间接引用 (*p).X 直接写成 p.X
	fmt.Println("修改后的结果为:",v)
	ptr := &Vertex{X:1}
	ptr.Y = 20
	fmt.Println(ptr, ptr.X, ptr.Y)

}

// 可切片来处理数组,数组的大小一旦确定不能更改
// 一个数组变量表示整个数组，它不是指向第一个元素的指针（不像 C 语言的数组）
func ArrayLearn(){
	var a[2] string // var a [2]string 也是对的 int *a 和 int* a 在 C 都是允许的
	a[0] = "hello"
	a[1] = "world" // var a = [...] string {"hello","world"} 会在类型检查期间编译器自动判断出数组大小
	// 也可以var a  = [2] string{"hello","world"}
	fmt.Println(a,a[0],a[1]) // a 表示整个数组而不是首元素指针
	// 创建一个字符串指针数组
	var ptrArr[2] *string
	ptrArr[0] = &a[0]
	fmt.Println(*ptrArr[0])
	// 尝试创建一个字符串类型的数组指针(指针指向字符串数组)
	// Go 语言不支持指简单针指向字符串数组的首地址a;长度.
	// var Arrptr *string = a 是错误的！cannot use 'a'(type [2]string)as type *string in assignment
	var p *[2]string // p 的类型是 *[2]string
	p = &a
	fmt.Printf("%T,%v",p,p)

}

// 切片下标从 0 开始，左闭右开，取头不取尾，类似python的切片用法
// 切片不存储数据，更改切片的元素会修改底层数组中对应的元素(值、长度、容量).
// 本质上一个切片是一个数组片段的描述。它包含了指向数组的指针，片段的长度， 和容量
func sliceLearn(){
	primes := [6] int{0,1,3,5,7,9}
	var set = primes[1:3]// length = right -left
	fmt.Println(set) // [1,3]
	set[0] = 100 // 将原数据1修改100
	fmt.Println(primes) // [0,100,3,5,7,9]
	// 切片文法类似于没有长度的数组文法
	 s := []struct {
		i int
		b bool
	}{
		{2,true},
		{3,true},
		{5,false}, // ‘,’ 不能省略！！！
	}
	// 上述是申明并创建一个结构体数组同时构建了引用他的切片
	fmt.Println(s)
	 // 切片长度和容量可以用len(s)、cap(s)表示
	 // 截取切片使其长度为 0 (其实就是左闭右开)
	 s0 := s[:0] // len=0,cap=3
	 fmt.Printf("len=%d,cap=%d \n",len(s0),cap(s0))
	 s1 := s[:2] // len=2,cap=3
	 fmt.Printf("len=%d,cap=%d \n",len(s1),cap(s1))
	 s2 := s[2:] // len=1,cap=1
	 fmt.Printf("len=%d,cap=%d \n",len(s2),cap(s2))
	 // make 动态创建数组并返回引用他的切片
	 b := make([]int,0,5) // return a slice len=0,cap=5
	 fmt.Printf("len=%d,cap=%d \n",len(b),cap(b))
	 // 切片可以嵌套
	 sliceDemoString := [][]string{
	 	[]string{"_","_","_"},
	 	[]string{"_","_","_"},
	 	[]string{"_","_","_"},
	 }
	 sliceDemoString[0][0] = "X"
	 sliceDemoString[0][2] = "X"
	 sliceDemoString[1][1] = "O"
	 sliceDemoString[2][0] = "X"
	 sliceDemoString[2][2] = "X"
	 sliceDemo(sliceDemoString)
	 // 切片可以append.创建切片的时候，cap 的值一定要保持清醒，避免共享原数组导致的 bug。
	 p := []byte{2, 3, 5}
	 z := append(p,7, 11, 13)
	 y := append(p,1,1,1,1,1,1)
	// p == []byte{2, 3, 5, 7, 11, 13}
	fmt.Println(p,len(p),cap(p)) // [2 3 5] 3 3
	fmt.Println(z,len(z),cap(z)) // [2 3 5 7 11 13] 6 8;此时的p仍然是[2,3,5]因为因为原来数组的容量已经达到了最大值，再想扩容， Go 默认会先开一片内存区域，把原来的值拷贝过来，然后再执行 append() 操作。这种情况丝毫不影响原数组。
	fmt.Println(y,len(y),cap(y)) // [2 3 5 1 1 1 1 1 1] 9 16
	/*GoLang中的切片扩容机制，与切片的数据类型、原本切片的容量、所需要的容量都有关系，比较复杂。对于常见数据类型，在元素数量较少时，大致可以认为扩容是按照翻倍进行的。但具体情况需要具体分析。*/

}

func sliceDemo(game [][]string)  {
	for i:=0;i<len(game) ;i++  {
		for j := 0; j < len(game[i]);j++{
			fmt.Printf(game[i][j]+" ")
		}
		fmt.Println()
	}
}


func main(){
	sliceLearn()

}