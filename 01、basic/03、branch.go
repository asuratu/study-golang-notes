package main

import (
	"fmt"
	"io/ioutil"
)

// 示例1 if else 的使用
func branch() {
	// 这里是文件的绝对路径
	const filename = "/Users/asura/Code/go/taobao/01、basic/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(contents))
	}
}

// 示例2 switch 的使用
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

// 示例3 switch 的另一种用法
func grade(score int) string {
	var result string
	switch {
	case score < 60:
		result = "D"
	case score < 70:
		result = "B"
	case score < 90:
		result = "C"
	case score <= 100:
		result = "A"
	default:
		panic(fmt.Sprintf("error score: %d\n", score))
	}
	return result
}

func main() {
	branch()
	fmt.Println("-----------------------")
	fmt.Println(eval(3, 2, "/"))
	fmt.Println(grade(99))
	fmt.Println(grade(101))
}
