package main

import "fmt"

func printArray(arr []int) {
	arr[0] = 1000
	for _, v := range arr {
		fmt.Println(v)
	}
}

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	fmt.Println("-------------------")
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// 数组取得slice，arr[:]
	printArray(arr[:])
	fmt.Println("-------------------")
	fmt.Printf("arr[2:6] = %v\n", arr[2:6])
	fmt.Printf("arr[:6] = %v\n", arr[:6])
	fmt.Printf("arr[2:] = %v\n", arr[2:])
	fmt.Printf("arr[:] = %v\n", arr[:])
	// slice 是对 arr 的一个 view
	// arr[2:6] 表示看 arr 的 2到6，包括2不包括6
	fmt.Println("---------update slice----------")
	s1 := arr[2:6]
	updateSlice(s1)
	fmt.Println(s1)
	s2 := arr[:]
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println("---------reslice----------")
	// 在slice的基础上再次slice
	s2 = s2[3:]
	fmt.Println(s2)
	s2 = s2[:2]
	fmt.Println(s2)
	fmt.Println("---------extending slice----------")
	arr[0], arr[2] = 0, 2
	slice1 := arr[2:6]
	slice2 := slice1[3:5]
	fmt.Println(slice1)
	fmt.Println(slice2)
}
