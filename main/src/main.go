package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"reflect"
	"runtime"
	"strconv"
)

// if
func pfile(){
	const filename = "src/123.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil{
		fmt.Println(string(contents))
	} else {
		fmt.Println("connot print file contents:", err)
	}
}

// switch
func eval(a, b int, op string) int {
	result := 0
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	case "%":
		_, result = div(a, b)
	default:
		panic(fmt.Sprintf("Wrong op " + op))
	}

	return result
}

// for
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}

	return result
}

// for（while）

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt. Println(scanner.Text())
	}
}

// func
func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("\nCalling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}

	return s
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// pointer


func main(){

	// hello world
	fmt.Println("Hello World!")

	// iftest
	pfile()

	// switch test
	fmt.Println(eval(13,5, "%"))

	// for test
	fmt.Println(convertToBin(5), convertToBin(13), convertToBin(1234534253))

	// for(which)
	printFile("src/123.txt")

	// func
	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))
}
