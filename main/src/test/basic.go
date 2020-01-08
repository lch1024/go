package test

import "math"

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

// 练习
// 寻找最长不含有重复字符的字串（leetcode） 第三题 可以适配中文
//

// 防止总 gc回收 导致效率低
var SlicePos = make([]int, 0xffff)
func lengthOfLongestSubstring(s string) int {
	nStart := 0
	nMaxLength := 0
	for i := range SlicePos {
		SlicePos[i] = 0
	}

	for i, v := range []rune(s) {
		if  pos := SlicePos[v]; pos > nStart {
				nStart = pos
		}

		len := i - nStart + 1
		if len > nMaxLength {
			nMaxLength = len
		}
		SlicePos[v] = i + 1
	}

	return nMaxLength
}

func main() {

}
