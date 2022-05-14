package main

import "fmt"

// 包内部变量
var (
	aa = "aa"
	bb = 123
	cc = false
)

// 定义一个变量
func variableZeroValue() {
	// var 定义一个变量之后，都会有初始值
	var a int    // 0
	var s string // ""
	// Println 直接打印变量
	fmt.Println(a, s)
	// Printf 格式化打印变量
	// %d 打印整数、%s 打印字符串、%q 打印字符串以及引号
	fmt.Printf("%d %q\n", a, s)
}

// 定义一个变量，并且赋初始值
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

// 自动推断变量类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "abc"
	fmt.Println(a, b, c, s)
}

// 简化定义变量
func variableShorter() {
	a, b, c, s := 3, 4, true, "abc"
	b = 5
	fmt.Println(a, b, c, s)
}

func horizon() {
	fmt.Println("---------------------")
}

func main() {
	fmt.Println("test")
	horizon()
	variableZeroValue()
	horizon()
	variableInitialValue()
	horizon()
	variableTypeDeduction()
	horizon()
	variableShorter()
	horizon()
	fmt.Println(aa, bb, cc)
}
