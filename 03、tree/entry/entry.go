package main

import (
	"fmt"
	"tree"
)

// 定义别名
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	// 不能通过字面量获取地址
	left := &myTreeNode{myNode.node.Left}
	left.postOrder()
	right := &myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

// 给结构体定义方法
// 方法前面加上接收者

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left = new(tree.Node)
	fmt.Println(root)
	root.Right.Left.SetValue(4)
	fmt.Println(root)

	fmt.Println("--------------------")

	nodes := []tree.Node{
		{Value: 0},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	fmt.Println("---------- 遍历tree ----------")
	root.Traverse()
	fmt.Println("---------- 遍历myTree ----------")
	myRoot := &myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println("--------------------")
	var pRoot *tree.Node
	pRoot.SetValue(200)
	// 在nil指针中获取属性值会报错
	pRoot = &root
	pRoot.SetValue(300)
	pRoot.Print()
}
