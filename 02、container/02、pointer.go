package main

import "fmt"

// 交换两个变量的值
func swapByRef(a, b *int) {
	*a, *b = *b, *a
}

func swapByVal(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 3, 4
	swapByRef(&a, &b)
	fmt.Println(a, b)
	a, b = swapByVal(a, b)
	fmt.Println(a, b)
}
