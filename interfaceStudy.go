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

// 接口的命名一般约定 er,　r, able 结尾(视情况而定)
// 接口类型 是由一组[方法]签名定义的集合。(0个或多个方法的签名,不包含实现)
// 接口类型的变量 可以保存任何实现了这些方法的值。
// 实现接口不需要显式的声明，只需实现相应方法即可，故没有implements关键字

type I interface { // 接口定义
	M() time.Time // 注意需要指出方法的返回值类型
}

type T struct {
	mytime time.Time
}

// 此方法表示类型T实现了接口 I ，并成功调用了方法M()
func (t T) M() time.Time {
	t.mytime = time.Now()
	return t.mytime
}

func main() {
	var i I = T{time.Now()} // 这里声明变量要用 "var" 不能用":="。因为 “:=” 只能适用于局部变量
	fmt.Println(i.M())      // 实现了接口就可以调相应的方法

}
