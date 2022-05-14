package main

import (
	"fmt"
	"math"
)

// 强制类型转换，没有隐式类型转换
func triangle() {
	a, b := 3, 4
	var c int
	// TODO 这里存在浮点数转换时丢失精度的问题
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量的使用
func consts() {
	const (
		// 常量不定义类型的话，就相当于一个文本，数值可以为各种类型使用
		// 如果定义了类型，则只能被当做该类型使用
		filename     = "abc.text"
		a, b     int = 1, 2
	)
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c)
}

//枚举的使用
func enums() {
	const (
		// iota表示自增值
		cpp = iota
		_
		java
		goland
		python
		js
		php
	)
	fmt.Println(cpp, java, goland, python, js, php)

	// b kb mb tb pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		tb
		pb
	)
	fmt.Println(b, kb, mb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	consts()
	enums()
	triangle()
}
