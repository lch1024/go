package main

import (
	"bufio"
	"fmt"
	"os"
)

// defer调用
// 1.确保调用在函数结束时发生
// 2.参数在defer语句时计算
// 3.defer列表为后进先出

// 何时使用defer调用
// 1.Open/Close
// 2.Lock/Unlock
// 3.PrintHeader/PrintFooter

func tryDefer() {

	// defer 函数结束后执行
	// defer 相当于个栈先进后出
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)

	// 参数在defer语句时计算
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			println("printed too many")
			return
		}
	}
}

type intGen func() int

func Fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

// 自定义错误结构
type myError struct {
	error string
	nRet int
}

// 实现错误接口
func (err *myError) Error() string {
	return err.error
}

func writeFile(filename string) {
	file, err := os.Create(filename)

	// 出错处理
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		fmt.Println("Error:", err.Error())
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}

}

func main() {
	tryDefer()
	writeFile("./src/errhanding/fib.txt")
}
