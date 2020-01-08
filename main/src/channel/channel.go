package main

import (
	"fmt"
	"time"
)

// 						channel
// -------------------------------------------------
//|  __________    			    	 __________     |
//| | goroutine| 《---------------》| goroutine|	|
//|  ----------					     ----------		|
//|  __________						 __________	    |
//| | goroutine| 《---------------》| goroutine|	|
//|  ----------     			     ----------     |
//|  _____________________________________________	|
//| |				  调度器					  | |
//|  ---------------------------------------------	|
// -------------------------------------------------

// channel
// bufferd channel
// range
// 理论基础：Communication Sequential Process(CSP) 这是一篇论文
// 不要通过共享内存来通信;通过通信来共享内存

func worker(id int, c chan int) {
//	for {
//		n, ok := <-c
//		if !ok {
//			break
//		}
//		fmt.Printf("Worker %d recieved %d\n", id, n)
//	}

	for n := range c {
		fmt.Printf("Worker %d recieved %d\n", id, n)
	}
}

//
func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	//var c chan int // 定义channel 里面是int 此时 c ==

	var channels [10]chan int
//	for i := 0; i < 10; i++ {
//		channels[i] = make(chan int)
//		go worker(i, channels[i])
//	}

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	// maked第二个参数为缓冲区 提升效率
	// 不然发送会一直等待接收
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

// 通知消息发完了
func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) // 执行完不断发0 string(空串)

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}
