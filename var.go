package main

import "fmt"

// 全局变量：只能用 var 声明(用法1：var 变量名 类型 = 值).
var name string = "小明"

func main() {
	/*
	   知识点：Go变量声明
	   1. var 声明（完整写法，可在函数数内外使用）
	   var 变量名 类型 = 值
	*/

	fmt.Println("我的名字叫：", name)
	// var 变量名 类型 （先声明后赋值）
	var age int
	age = 18
	fmt.Printf("今年%d岁。\n", age)
	// 2. 短变量 := （自动推导类型）（简洁写法，只能在函数内部使用）
	height := 175.5
	fmt.Println("我的身高是：", height)
}
