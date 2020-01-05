package main

import "fmt"

// recover 替代panic优雅拿未知错误
func tryRecover() {

	// 加() 执行匿名函数
	defer func() {
		err := recover()
		if err, ok := err.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			return
		}
	}()

	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
