package main

import (
	"fmt"
	"queue"
)

func main() {
	// 使用组合

	// 定义一个 Queue
	// var q queue.Queue
	// q := make(queue.Queue, 0)
	q := queue.Queue{1}
	q.Push(3)
	q.Push(4)
	fmt.Println(q)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
