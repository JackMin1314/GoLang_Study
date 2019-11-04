package main

import (
	"fmt"
	"unsafe"
)

// 与 C 不同, Go 没有指针运算。默认空指针为nil
func ptrLearn() {
	var p *int
	i := 22
	p = &i         // 取地址符号
	fmt.Println(p) // *p 输出为22; p 输出为保存i的地址
	fmt.Println(&i)
}

// 结构体学习
func structDefine() {
	type Vertex struct {
		X int
		Y int
	}
	// 如果没有初始化值，默认为各自类型的空值 int 为 0
	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v
	p.X = 10 // 语言允许隐式间接引用 (*p).X 直接写成 p.X
	fmt.Println("修改后的结果为:", v)
	ptr := &Vertex{X: 1}
	ptr.Y = 20
	fmt.Println(ptr, ptr.X, ptr.Y)

}

// 可切片来处理数组,数组的大小一旦确定不能更改
// 一个数组变量表示整个数组，它不是指向第一个元素的指针（不像 C 语言的数组）
func ArrayLearn() {
	var a [2]string // var a [2]string 也是对的 int *a 和 int* a 在 C 都是允许的
	a[0] = "hello"
	a[1] = "world" // var a = [...] string {"hello","world"} 会在类型检查期间编译器自动判断出数组大小
	// 也可以var a  = [2] string{"hello","world"}
	fmt.Println(a, a[0], a[1]) // a 表示整个数组而不是首元素指针
	// 创建一个字符串指针数组
	var ptrArr [2]*string
	ptrArr[0] = &a[0]
	fmt.Println(*ptrArr[0])
	// 尝试创建一个字符串类型的数组指针(指针指向字符串数组)
	// Go 语言不支持指简单针指向字符串数组的首地址a;长度.
	// var Arrptr *string = a 是错误的！cannot use 'a'(type [2]string)as type *string in assignment
	var p *[2]string // p 的类型是 *[2]string
	p = &a
	fmt.Printf("%T,%v", p, p)

}

// 切片下标从 0 开始，左闭右开，取头不取尾，类似python的切片用法
// 切片不存储数据，更改切片的元素会修改底层数组中对应的元素(值、长度、容量).
// 本质上一个切片是一个数组片段的描述(结构体)。它包含了指向数组的指针,片段的长度和容量
func sliceLearn() {
	primes := [6]int{0, 1, 3, 5, 7, 9}
	var set = primes[1:3] // length = right -left
	fmt.Println(set)      // [1,3]
	set[0] = 100          // 将原数据1修改100
	fmt.Println(primes)   // [0,100,3,5,7,9]
	// 切片文法类似于没有长度的数组文法
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{5, false}, // ‘,’ 不能省略！！！
	}
	// 上述是申明并创建一个结构体数组同时构建了引用他的切片
	fmt.Println(s)
	// 切片长度和容量可以用len(s)、cap(s)表示
	// 截取切片使其长度为 0 (其实就是左闭右开)
	s0 := s[:0] // len=0,cap=3
	fmt.Printf("len=%d,cap=%d \n", len(s0), cap(s0))
	s1 := s[:2] // len=2,cap=3
	fmt.Printf("len=%d,cap=%d \n", len(s1), cap(s1))
	s2 := s[2:] // len=1,cap=1
	fmt.Printf("len=%d,cap=%d \n", len(s2), cap(s2))
	// make 动态创建数组并返回引用他的切片
	b := make([]int, 0, 5) // return a slice len=0,cap=5
	fmt.Printf("len=%d,cap=%d \n", len(b), cap(b))
	// 切片可以嵌套
	sliceDemoString := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	sliceDemoString[0][0] = "X"
	sliceDemoString[0][2] = "X"
	sliceDemoString[1][1] = "O"
	sliceDemoString[2][0] = "X"
	sliceDemoString[2][2] = "X"
	sliceDemo(sliceDemoString)
	// 切片可以append.创建切片的时候，cap 的值一定要保持清醒，避免共享原数组导致的 bug。
	p := []byte{2, 3, 5}
	z := append(p, 7, 11, 13)
	y := append(p, 1, 1, 1, 1, 1, 1)
	// p == []byte{2, 3, 5, 7, 11, 13}
	fmt.Println(p, len(p), cap(p)) // [2 3 5] 3 3; 注意这里的p不包含append的添加元素，也就是说使用append会返回新的切片。需要在原来的基础上赋给p。p=append(p,elems)
	fmt.Println(z, len(z), cap(z)) // [2 3 5 7 11 13] 6 8;此时的p仍然是[2,3,5]因为因为原来数组的容量已经达到了最大值，再想扩容， Go 默认会先开一片内存区域，把原来的值拷贝过来，然后再执行 append() 操作。这种情况丝毫不影响原数组。
	fmt.Println(y, len(y), cap(y)) // [2 3 5 1 1 1 1 1 1] 9 16
	/*GoLang中的切片扩容机制，与切片的数据类型、原本切片的容量、所需要的容量都有关系，比较复杂。对于常见数据类型，在元素数量较少时，大致可以认为扩容是按照翻倍进行的。但具体情况需要具体分析。*/

}

// 打印二维字符串数组
func sliceDemo(game [][]string) {
	for i := 0; i < len(game); i++ {
		for j := 0; j < len(game[i]); j++ {
			fmt.Printf(game[i][j] + " ")
		}
		fmt.Println()
	}
}

// for 循环的 range 形式可遍历切片或映射。每次循环返回两个值，一个为当前元素下标，第二个为该下标对应元素的副本！( Value 其实是切片里面的值拷贝,地址不变)
func rangeLearn() {
	var intSlice = [4]int{1, 0, 2, 0}
	for index, value := range intSlice {
		fmt.Printf("The index is %d,value is %v\n", index, value)
	}
	// 可以忽略索引或者值，响应用_代替
	strSlice := [...]string{"hello", "world", "2019", "1020"}
	for _, value := range strSlice {
		fmt.Println(value)
	}
	// 如果只需要index可以只用索引, 注意 range 下标从 0 开始
	// 利用移位求 2 的 10 次方
	for indexInt := range make([]int, 11) {
		fmt.Println(1 << uint(indexInt)) // uint占8个字节 每次输出...00000001,...00000010,...00000100,...00001000,...00010000,...00100000,...01000000,...10000000,......
	} //              1            2          4             8        16           32           64         128   .....
	fmt.Println(unsafe.Sizeof(uint(1)))  // 8 = 4 + 4
	fmt.Println(unsafe.Sizeof(strSlice)) // 64 = 4*16
	// 详细的sizeof见下面testSizeOf()

}

// 关于 sizeof 的理解，需要导入unsafe包,用于Go编译器，编译阶段使用,绕过go语言的类型系统(更高效)直接读写内存.
// Golang 的sizeof 跟 C 的 sizeof 不一样。
// 一般官方不建议使用unsafe包，会引发有莫名其妙的bug(试图引入一个uintptr类型的临时变量,一个非指针的临时变量tmp,只是一个普通的数字,导致垃圾收集器无法正确识别这个是一个指向变量x的指针)
// Sizeof 返回被值 v 所占用的字节大小,但是不包括v所指向的内存大小
// if v is a slice, it returns the size of the slice descriptor, not the size of the memory referenced by the slice.
// 而非该切片引用的内存大小。
func testSizeOf() {
	var arrInt = [4]int{0, 1, 3, 2}                      // 数组
	var myStr = "String"                                 // 字符串
	var arrStr = [3]string{"China", "France", "England"} // 字符串数组
	var intSlice = arrInt[:3]                            // 切片
	stuStruct := struct {
		name string
		age  int
	}{
		name: "Chen Wang",
		age:  10,
	} // 结构体
	// 数组大小，考虑 个数*类型
	fmt.Println(unsafe.Sizeof(arrInt)) // 32 = 4 * 8
	/**
	本机int占8个字节，4*8=32
	*/

	// 字符串大小
	fmt.Println(unsafe.Sizeof(myStr)) // 16 = 1*16
	// 第一个域是指向该字符串的指针,第二个域是字符串的长度
	/* 字符串底层数据结构
	type StringHeader struct {
	    Data uintptr     // 8 byte
	    Len  int         // 8 byte
	}
	*/

	// 字符串数组大小
	fmt.Println(unsafe.Sizeof(arrStr)) // 48 = 3*16

	// 切片大小,注意是切片,不考虑指针指向那块区域大小
	fmt.Println(unsafe.Sizeof(intSlice)) // 24 = 1*24
	/* 切片底层结构为
	type SliceHeader struct {
	    Data uintptr // 8 byte
	    Len  int     // 8 byte
	    Cap  int     // 8 byte
	}
	*/

	// 结构体大小
	fmt.Println(unsafe.Sizeof(stuStruct)) // 24 = 16 + 8

}

// 映射map,可以看成是具有键的结构体,key-value.类似python的字典,跟C++中的map类似;但是底层上map实现的还是C++里面的unordered_map(哈希表)
func mapLearn() {
	type std struct {
		name string
		age  int
	}
	var m map[string]std     // 仅申明,没有分配空间
	m = make(map[string]std) // 基本类型分配空间
	fmt.Println(unsafe.Sizeof(m))
	m["CW"] = std{
		"ChenWang", 23,
	}
	m["ZYM"] = std{
		name: "ZhangYueMin",
		age:  21,
	}
	fmt.Println(m)       // map[Chen Wang:{ChenWang 23} ZYM:{ZhangYueMin 21}]
	fmt.Println(m["CW"]) // {ChenWang 23}
	// 对于上面还可以简单表示为
	var mm = map[string]std{
		"CW":  {"ChenWang", 23},
		"ZYM": {"zym", 21},
	}
	fmt.Println(mm)
}

// 修改map
func mapChange() {
	m := make(map[string]int) // 声明并创建一个map
	// 插入数据
	m["Chinese"] = 80
	m["English"] = 70
	m["Math"] = 90
	// 修改数据
	m["English"] = 80
	// 删除数据
	delete(m, "Math")
	// 检测键是否存在
	k, guess := m["Math"]
	fmt.Println(m)
	fmt.Printf("The value is %v, is Present? %v\n", k, guess) // 0,false;检测“Math”是否存在，不存在返回false,然后k为元素类型的零值
	index, check := m["English"]
	fmt.Println(index, check) // 80,true;检测key "English"存在，然后返回对应的value
	// 遍历map;python里面用for li in m:
	for keys, values := range m {
		fmt.Println(keys, values)
	}
}

func main() {
	mapChange()
}
