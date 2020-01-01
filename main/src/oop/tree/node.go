package tree

import (
	"fmt"
)

// 结构声明 里面只有参数
type Node struct {
	Value       int
	Left, Right *Node
}

// 结构函数声明
// 接受者(node treeNode)
func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}

// 函数传引用 可以改变其值
func (node *Node) SetValue(value int) {
	if nil == node {
		fmt.Println("setting value to nil node. Ingored!")
		return
	}
	node.Value = value
}

// 工厂函数
// go局部变量也可以返回
// 再堆还是栈上 根据返回 分配 如果不是取地址 那可能是栈 如果是取地址 那么是堆
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

// 遍历
func (node *Node) Travere() {
	if nil == node {
		return
	}

	node.Left.Travere()
	node.Print()
	node.Right.Travere()
}
