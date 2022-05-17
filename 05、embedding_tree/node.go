package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// Print
// 给结构体定义方法
// 方法前面加上接收者
func (node Node) Print() {
	fmt.Print(node.Value, " \n")
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. ignored")
		return
	}
	node.Value = value
}

// CreateNode
//工厂函数来控制结构体的构造
func CreateNode(value int) *Node {
	// 返回的是一个局部变量的地址
	return &Node{
		Value: value,
	}
}
