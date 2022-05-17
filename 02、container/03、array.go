package main

import "fmt"

func loop(arr [5]int) {
	for i, v := range arr {
		fmt.Println(arr[i], v)
	}
}

func main() {
	// var 定义可以不赋初始值
	var arr1 [5]int
	// 用 := 定义要给初始值
	arr2 := [3]int{1, 3, 5}
	// [...]表示自动计算数组的长度
	arr3 := [...]int{1, 3, 5, 4, 6}
	// 定义二维数组
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println("-------------------")
	for i, v := range arr3 {
		fmt.Println(arr3[i], v)
	}

	fmt.Println("-------------------")
	loop(arr3)
}
