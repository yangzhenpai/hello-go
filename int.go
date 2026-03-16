package main

import "fmt"

func main() {
	// 定义整数int类型的变量
	var age int = 25
	// 大整数(num)用int64避免溢出
	var num int64 = 100000000
	fmt.Println(age) //输出整数25
	fmt.Println(num) // 输出：大整数：10000000000
	// 整数运算
	fmt.Println(10 + 5) // 加法输出结果15
	fmt.Println(10 - 5) // 减法输出结果5
}
