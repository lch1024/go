package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

// 练习
// 寻找最长不含有重复字符的字串（leetcode） 第三题 可以适配中文
//

func lengthOfLongestSubstring(s string) int {
	nStart := 0
	nMaxLength := 0
	mapPos := make(map[rune]int)

	for i, v := range []rune(s) {
		if pos, ok := mapPos[v]; ok {
			if pos >= nStart {
				nStart = pos + 1
			}
		}

		nMaxLength = int(math.Max(float64(i-nStart+1), float64(nMaxLength)))
		mapPos[v] = i
	}

	return nMaxLength
}

func main() {

	s := "Yes我爱敲代码!"
	fmt.Println(len(s))

	// byte utf-8编码 (可变长 中文3字节)
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}

	fmt.Println()

	// rune utf-8 转成unicode编码值 rune就是 int32
	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	// 另外开数组
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("我爱敲代码!"))
	fmt.Println(lengthOfLongestSubstring("一二三二一"))
	fmt.Println(lengthOfLongestSubstring("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))

	// 使用range遍历pos，rune对
	// 使用utf8.RuneCountInString获得字符数量
	// 使用len获得字节长度
	// 使用[]byte获得字节

	// 其他操作
	// Fields, Split, Join
	// Contains, Index
	// ToLower, ToUpper
	// Trim, TrimRight, TrimLeft
}
