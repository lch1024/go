package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// go语言闭包的应用
// 1.更为自然,不需要修饰如何访问自由变量
// 2.没有Lambda表达式,但是有匿名函数

// 非正统函数式编程
// 闭包概念
// 函数体有“局部变量” 比如V  有“自由变量” 比如sum（ 可以是结构 机构下可能还有结构 自由变量会把所有的都连接到
// 这是返回的 就是一个闭包）
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

// 正统函数式编程
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// 1、菲波那切数列
func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

// 2、为函数实现接口
// type类型就可以实现接口
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// 有可能 byte太大  一次read不全 这里下一次才能read完 但是 next值已经改变了 就会有问题
	// 这里就需要对 Reader 做缓存
	return strings.NewReader(s).Read(p)
}

func printfContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	// 传入 intGen 类型的io 实现Read()
	// scanner.Scan() 会调用io.Read()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// 3、使用函数遍历二叉树
type Nnode struct {
	Value       int
	Left, Right *Nnode
}

func (node Nnode) Print() {
	fmt.Printf("%d ", node.Value)
}

func (node *Nnode) Traverse() {
	node.TraverseFunc(func(n *Nnode) {
		n.Print()
	})
	fmt.Println()
}

func (node *Nnode) TraverseFunc(f func(*Nnode)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a(i))
	}

	b := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, b = b(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}

	f := fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	printfContents(f)

	var root Nnode

	root = Nnode{Value: 3}
	root.Left = &Nnode{}
	root.Right = &Nnode{5, nil, nil}
	root.Right.Left = new(Nnode)
	root.Left.Right = &Nnode{Value:2}

	root.Traverse()

	nCount := 0
	root.TraverseFunc(func(n *Nnode){
		nCount++
	})
	fmt.Println("Node count: ", nCount)

}
