package main

import (
	"fmt"
	"io/ioutil"
)

func branch() {
	// 这里是文件的绝对路径
	const filename = "/Users/asura/Code/go/taobao/01、basic/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(contents))
	}
}

func grade(score int) {
	g := ""
	// switch 后面可以没有表达式
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score %d", score))
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100:
		g = "A"
	}
	fmt.Println(g)
}

func main() {
	branch()
	fmt.Println("##################")
	grade(56)
	grade(200)
	grade(70)
	grade(89)
	grade(99)
}
