package main

import (
	"fmt"
	"time"
)

// goroutine (协程Coroutine)
// 轻量级 “线程”
// "非抢占式"多任务处理,由协程主动交出控制权
// 编译器/解释器/虚拟机层面的多任务 （go 编译器级别多任务）
// 多个协程可能在一个或多个线程上运行

// 1.4.2 Coroutines
// -子程序是协程的一个特例(所有函数调用都相当于一个子程序)

// 普通函数
// 线程内： main() ---> doWork

// 协程						     双向通道
// 线程内(一个或者多个) main() <--------->  doWork

// goroutine可能的切换点
// 1. I/O, select
// 2. channel
// 3. 等待锁
// 4. 函数调用 (有时)
// 5. runtime.Gosched()
// 以上只是参考,不能保证切换,不能保证其他地方不切换

// 其他语言中的协程
// C++ : Boost.Coroutine
// Java : 不支持协程
// python 使用yield关键字实现协程 Python 3.5加入了async def对协程原生支持

// -------------------------------------------------
//|  ___________     ___________      __________    |
//| |   线程   |	|	 线程	|	 |	 线程   |   |
//| | goroutine|	| goroutine	|	 | goroutine|	|
//| | goroutine|	|			|	 | ...		|	|
//|  ----------      -----------      ----------    |
//|  _____________________________________________	|
//| |				  调度器					  | |
//|  ---------------------------------------------	|
// -------------------------------------------------

// go run -race goroutine
// -race 数据访问冲突检测
// 检测之后 可以检测出a被协程写 又被 fmt.Println读
func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		// 函数前加go 声明函数为协程
		go func(ii int){
			for {
				//a[ii]++ // 只有a[i]++ 交不出控制权 所以会死循环
				//runtime.Gosched() // 交出控制权
				fmt.Printf("Hello from goroutine %d\n", ii) // io操作可以交出控制权
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}