package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("slice=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	var s []int // Zero value for slice is nil
	fmt.Println("--------- Append slice ----------")
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	fmt.Println("--------- Creating slice ----------")
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	s2 := make([]int, 16, 32)
	printSlice(s2)
	fmt.Println("--------- Copy slice ----------")
	// 将s1复制给s2
	copy(s2, s1)
	printSlice(s2)
	fmt.Println("--------- Delete slice ----------")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
	fmt.Println("--------- Popping from front ----------")
	front := s2[0]
	fmt.Println(front)
	s2 = s2[1:]
	printSlice(s2)
	fmt.Println("--------- Popping from back ----------")
	tail := s1[len(s1)-1]
	fmt.Println(tail)
	s1 = s1[:len(s1)-1]
	printSlice(s1)

}
