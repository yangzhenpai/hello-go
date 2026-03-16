package main

import "fmt"

func main() {
	// 定义float类型变量
	var price float64 = 99.9
	var pi float32 = 3.1415926
	fmt.Println("价格输出：", price) // 输出命价格：99.9
	fmt.Println("圆周率", pi)      // 圆周率输出：3.1415926（float32精度限制）
	// 浮点运算
	fmt.Println(2.5 * 4)  // 乘法输出：10
	fmt.Println(10.0 / 3) // 除法输出：3.3333333333333335
}
