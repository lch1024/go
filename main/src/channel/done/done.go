package main

import (
	"fmt"
	"sync"
)
func doWorker(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d recieved %c\n", id, n)
		done <- true
	}
}

type worker struct {
	in 	 chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in: make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, w.in, w.done)
	return w
}

func DoWorker(id int, w Worker) {
	for n := range w.in {
		fmt.Printf("Worker %d recieved %c\n", id, n)
		w.done()
	}
}

type Worker struct {
	in 	 chan int
	done func()
}

func CreateWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go DoWorker(id, w)
	return w
}



func chanDemo() {
	//var c chan int // 定义channel 里面是int 此时 c ==

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//<-workers[i].done
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		//<-workers[i].done
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
	}
}

func chanDemo1() {
	//var c chan int // 定义channel 里面是int 此时 c ==

	var wg sync.WaitGroup
	var workers [10]Worker
	for i := 0; i < 10; i++ {
		workers[i] = CreateWorker(i, &wg)
	}

	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		//<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		//<-workers[i].done
	}

	wg.Wait()
}

func main() {
	fmt.Println("done 1")
	chanDemo()
	fmt.Println("\ndone 2")
	chanDemo1()
}
