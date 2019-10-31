package main

import (
	"fmt"
	"math"
)

func NewTon(x float64) float64 {
	var z float64 = 1.0
	temp := 1e-4 // 精度0.0001不同电脑计算能力不一样
	for ABS(z*z-x) > temp {
		z -= (z*z - x) / 2 * z
	}
	return z
}
func main() {

	x := 2.0
	fmt.Println(NewTon(x))
	fmt.Println(math.Sqrt(x)) // 与sqrt函数的结果相比较
}
// 自定义实现正负判断
func ABS(t float64) float64 {
	if t < 0 {

		t = -t
	}
	return t
}
