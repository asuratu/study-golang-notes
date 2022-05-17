## 面向对象
go语言仅支持封装，不支持继承和多态

## 结构的创建
```
type treeNode struct {
	value       int
	left, right *treeNode
}

var root treeNode
root = treeNode{value: 3}
root.left = &treeNode{}
root.right = &treeNode{5, nil, nil}
root.right.left = new(treeNode)
fmt.Println(root)

// 切片中的元素可以省略结构体名称
nodes := []treeNode{
    {value: 0},
    {},
    {6, nil, &root},
}

```
1. 不论是地址还是结构本身，都一律用 . 来访问成员
2. go语言没用构造函数，可以用自定义工厂函数来控制结构体的构造，返回局部变量的地址
```
func createNode(value int) *treeNode {
    return &treeNode{
        value: value,
    }
}
```
3. 不用知道变量分配在栈上还是在堆上，go根据代码运行环境自动分配
4. 给结构体定义方法，显式定义和命名方法接收者
```
func (node treeNode) print() {
    fmt.Println(node.value)
}
```
5. nil指针也可以调用方法

## 值接收者 vs 指针接收者（重要）
1. 要改变内容必须是指针接收者
2. 结构过大也考虑使用指针接收者
3. 一致性：如果有指针接收者，最好都是指针接收者
4. 值接收者是go语言特有的
5. 值/指针接收者均可以接收值/指针 

## 封装
1. 名字一般使用CamelCase
2. 首字母大写：public
3. 首字母小写：private

## 包
1. 每个目录一个包
2. main包包含了可执行入口
3. 为结构定义的方法必须放在同一个包内，但是可以是不同的文件

## 如何扩充系统类型或者别人的类型
1. 定义别名
2. 使用组合
3. 使用内嵌 Embedding
