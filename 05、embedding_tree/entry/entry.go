package main

import (
	"fmt"
	"tree"
)

// 定义内嵌 Embedding
type myTreeNode struct {
	// node *tree.Node
	// 这里省略node就是内嵌，其实就是一个语法糖
	*tree.Node
}

func (myNode *myTreeNode) Traverse() {
	fmt.Println("this method is shadowed.")
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}
	// 不能通过字面量获取地址
	left := &myTreeNode{myNode.Left}
	left.postOrder()
	right := &myTreeNode{myNode.Right}
	right.postOrder()
	myNode.Print()
}

func main() {
	root := myTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = new(tree.Node)
	root.Right.Left.SetValue(4)
	fmt.Println("---------- my traverse ----------")
	root.Traverse()
	fmt.Println("---------- tree traverse ----------")
	root.Node.Traverse()
	fmt.Println("---------- 遍历myTree ----------")
	root.postOrder()
	fmt.Println("-------------------------------")
}
