package inc

import "fmt"

func enums(){
	const(
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(b, kb, mb, gb, tb, pb)
}

func getsum(a int, b int) int {
	return a + b
}
