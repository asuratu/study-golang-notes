package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			panic("when a/b, b can`t 0")
		}
		result = a / b
	default:
		panic("unknown op")
	}
	return result
}

func div(a, b int) (q, r int) {
	//q, r = a/b, a%b
	//return
	// 建议的写法
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	// 通过反射拿到方法名
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funtion %s with args (%d, %d )\n", opName, a, b)
	return op(a, b)
}

func add(a, b int) int {
	return a + b
}

// 可变参数列表
func sum(numbers ...int) int {
	fmt.Printf("numbers 的类型：%T\n", numbers)
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("-------------------")
	fmt.Println(eval(10, 20, "*"))
	fmt.Println("-------------------")
	fmt.Println(div(13, 3))
	fmt.Println("-------------------")
	fmt.Println(apply(add, 3, 4))
	// 这里也可以直接传一个闭包
	fmt.Println(apply(func(i int, i2 int) int {
		return int(math.Pow(float64(i), float64(i2)))
	}, 3, 4))
	fmt.Println("-------------------")
	fmt.Println(sum(1, 2, 3, 4, 5))
}
