package main

import (
	"fmt"
	"oop/queue"
	"oop/tree"
)

// 定义别名
type MyTreeNode struct {
	node *tree.Node
}

// 扩展后续遍历算法
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	Left := MyTreeNode{myNode.node.Left}
	Right := MyTreeNode{myNode.node.Right}
	Left.postOrder()
	Right.postOrder()
	myNode.node.Print()
}

func main() {

	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, nil},
	}
	fmt.Println(nodes)

	root.Print()
	fmt.Println()

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(100)
	pRoot.Print()
	root.Print()

	var ppRoot *tree.Node
	fmt.Println(ppRoot)
	ppRoot.SetValue(200)
	ppRoot = &root
	ppRoot.SetValue(300)
	ppRoot.Print()
	fmt.Println()

	// 定义别名
	root.Travere()
	fmt.Println()
	mynode := MyTreeNode{&root}
	mynode.postOrder()
	fmt.Println()

	// 组合
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}

// go语言只支持封装，不支持继承和多态
// go语言没有class 只有 struct
// 显示定义和命名方法接收者
// 只有使用指针才能改变结构内容
// nil指针也可以调用方法!

// 值接收者 vs 指针接收者
// 1.要改变内容必须使用指针接收者
// 2.结构过大也考虑使用指针接收者
// 3.一致性：如果有指针接收者，最好都是指针接收者
// 值接收者是go语言特有的 (c++ this指针是指针接收者)

// 封装
// 名字一般使用CamelCase
// 首字母大写：public  针对包
// 首字母小写：private 针对包

// 包
// 每个目录一个包
// main包包含可执行入口
// 为结构定义的方法必须放在同一个包内
// 可以是不同的文件

// 如何扩充
// 可以扩充系统类型或者别人的类型
// 方法： 1.定义别名 2.使用组合
//
