package main

import (
	"fmt"
	"math"
)

// 练习
// 寻找最长不含有重复字符的字串（leetcode） 第三题
// 思路
// 1.记录当前字母最后一次出现的位置 如果比定位pos 小那么 继续 如果比定位的pos大 那么定位的pos 要当前位置 + 1
//
func lengthOfLongestSubstring(s string) int {
	nStart := 0
	nMaxLength := 0
	mapPos := make(map[byte]int)

	for i, v := range []byte(s) {
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
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m3 == empty map

	if m2 == nil {
		fmt.Println("m2 == nil")
	} else {
		fmt.Println("m2 != nil")
	}

	var m3 map[string]int // m3 == nil

	if m3 == nil {
		fmt.Println("m3 == nil")
	} else {
		fmt.Println("m3 != nil")
	}

	fmt.Println(m, m2, m3)

	// 遍历 无序 hash-map
	fmt.Println("Traversing map!!!")
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 获取value
	fmt.Println("Getting values!!!")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("key dose not exist")
	}

	// 删除key
	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

	fmt.Println(len(m))
	for k, v := range m {
		fmt.Println(k, v)
	}

	// 使用range遍历Key
	// 遍历是无序的
	// 使用len()获取长度
	// map 使用哈希表 必须比较相等
	// 除了slice map function 的内建类型都可以作为Key
	// Struct类型不包含上述字段，也可以作为Key
	// map[K] V 声明
}
