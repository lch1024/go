package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Printf("%v len=%d, cap=%d\n", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// 半开半闭区间
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:]
	s2 := arr[:]
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	// 扩展 可以向后扩展 不可以向前扩展
	s1 = arr[2:6]
	s2 = s1[3:5]

	fmt.Println("arr = ", arr)
	fmt.Println("s1 = ", s1)
	fmt.Println("s2 = ", s2)
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	// 添加
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3, s4, s5 =", s3, s4, s5)

	// 因为s2 caps 为3 所以s3 为 5 6 10 这时arr 最后一个个数也变成10
	// 而说s4 caps 已经不够用了 所以会重新分配更大的底层数组
	// s5同上
	fmt.Println("arr = ", arr)

	// 创建
	fmt.Println("Create slice")
	var s []int // Zero value for silice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 = []int{2, 4, 6, 8}
	printSlice(s1)

	s2 = make([]int, 16)
	printSlice(s2)

	s3 = make([]int, 10, 32)
	printSlice(s3)

	// 拷贝
	fmt.Println("Copy slice")
	copy(s2, s1)
	printSlice(s2)

	// 删除
	fmt.Println("Deleting elements from slice")
	// s2[:3] + s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	front := s2[0]
	s2 = s2[1:]
	fmt.Printf("front = %d, s2 = %v\n", front, s2)

	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Printf("tail = %d, s2 = %v\n", tail, s2)
}
