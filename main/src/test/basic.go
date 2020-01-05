package test

import "math"

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

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

}
